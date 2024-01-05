package labs

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/dms"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DmsMysqlPostgresOptions struct {
	AlocatedStorage int32
	KeyPairName     string
	AmiId           string
	PostgresUser    string
	PostgresPass    string
	MariaDBUser     string
	MariaDBPassword string
}

type DmsMysqlPostgresOutput struct {
	Mariadb             *rds.Instance
	Postgres            *ec2.Instance
	ReplicationInstance *dms.ReplicationInstance
	ReplicationTask     *dms.ReplicationTask
}

// Creates the infrastructure to do the migration/CDC lab.
// It requires further postgres configuration and some manual work
// to make it work 100%
func NewDmsMysqlPostgres(ctx *pulumi.Context, options *DmsMysqlPostgresOptions) (*DmsMysqlPostgresOutput, error) {
	group, err := ec2.NewSecurityGroup(ctx, "postgres-db-sg", &ec2.SecurityGroupArgs{
		Ingress: ec2.SecurityGroupIngressArray{
			&ec2.SecurityGroupIngressArgs{
				Protocol:   pulumi.String("tcp"),
				FromPort:   pulumi.Int(22),
				ToPort:     pulumi.Int(22),
				CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
			},
			&ec2.SecurityGroupIngressArgs{
				Protocol:   pulumi.String("tcp"),
				FromPort:   pulumi.Int(5432),
				ToPort:     pulumi.Int(5432),
				CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
			},
		},
		Egress: ec2.SecurityGroupEgressArray{
			&ec2.SecurityGroupEgressArgs{
				Description: pulumi.String("Allow all outbound traffic"),
				FromPort:    pulumi.Int(0),
				ToPort:      pulumi.Int(0),
				Protocol:    pulumi.String("-1"),
				CidrBlocks:  pulumi.StringArray{pulumi.String("0.0.0.0/0")},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	postgres, err := ec2.NewInstance(ctx, "postgres-db", &ec2.InstanceArgs{
		InstanceType:             pulumi.String("t4g.medium"),
		SecurityGroups:           pulumi.StringArray{group.Name},
		Ami:                      pulumi.String(options.AmiId),
		KeyName:                  pulumi.String(options.KeyPairName),
		AssociatePublicIpAddress: pulumi.Bool(true),
		RootBlockDevice: &ec2.InstanceRootBlockDeviceArgs{
			VolumeSize: pulumi.Int(options.AlocatedStorage),
		},
		Tags: pulumi.StringMap{
			"Name": pulumi.String("cdc-lab-postgres-db"),
			"Type": pulumi.String("lab"),
		},
	})
	if err != nil {
		return nil, err
	}

	mariadb, err := rds.NewInstance(ctx, "mariadb-rds", &rds.InstanceArgs{
		AllocatedStorage:  pulumi.Int(options.AlocatedStorage),
		DbName:            pulumi.String("mariadblab"),
		Engine:            pulumi.String("mariadb"),
		EngineVersion:     pulumi.String("10.11.6"),
		InstanceClass:     pulumi.String("db.t3.micro"),
		Username:          pulumi.String(options.MariaDBUser),
		Password:          pulumi.String(options.MariaDBPassword),
		SkipFinalSnapshot: pulumi.Bool(true),
		Tags: pulumi.StringMap{
			"Name": pulumi.String("cdc-lab-mariadb-rds"),
			"Type": pulumi.String("lab"),
		},
	})
	if err != nil {
		return nil, err
	}

	postgresEndpoint, err := dms.NewEndpoint(ctx, "postgres-endpoint", &dms.EndpointArgs{
		ServerName:   postgres.PublicIp,
		DatabaseName: pulumi.String("postgresdb"),
		Username:     pulumi.String(options.PostgresUser),
		Password:     pulumi.String(options.PostgresPass),
		EngineName:   pulumi.String("postgres"),
		Port:         pulumi.Int(5432),
		EndpointId:   pulumi.String("postgres-endpoint"),
		EndpointType: pulumi.String("source"),
	}, pulumi.DependsOn([]pulumi.Resource{postgres}))
	if err != nil {
		return nil, err
	}

	mariadbEndpoint, err := dms.NewEndpoint(ctx, "mariadb-endpoint", &dms.EndpointArgs{
		ServerName:   mariadb.Endpoint,
		DatabaseName: mariadb.DbName,
		Username:     mariadb.Username,
		EngineName:   mariadb.Engine,
		Port:         mariadb.Port,
		Password:     pulumi.String(options.MariaDBPassword),
		EndpointId:   pulumi.String("mariadb-endpoint"),
		EndpointType: pulumi.String("target"),
	}, pulumi.DependsOn([]pulumi.Resource{mariadb}))
	if err != nil {
		return nil, err
	}

	replicationInstance, err := dms.NewReplicationInstance(ctx, "dms-replication-lab", &dms.ReplicationInstanceArgs{
		AllocatedStorage:         pulumi.Int(options.AlocatedStorage * 2),
		MultiAz:                  pulumi.Bool(false), // dev pourposes
		ReplicationInstanceId:    pulumi.String("cdc-lab-replication-instance"),
		ReplicationInstanceClass: pulumi.String("dms.t2.medium"),
		Tags: pulumi.StringMap{
			"Name": pulumi.String("cdc-lab-instance"),
		},
	})
	if err != nil {
		return nil, err
	}

	replicationTask, err := dms.NewReplicationTask(ctx, "dms-replication-task", &dms.ReplicationTaskArgs{
		ReplicationInstanceArn: replicationInstance.ReplicationInstanceArn,
		SourceEndpointArn:      postgresEndpoint.EndpointArn,
		TargetEndpointArn:      mariadbEndpoint.EndpointArn,
		ReplicationTaskId:      pulumi.String("cdc-replication-task"),
		MigrationType:          pulumi.String("full-load-and-cdc"),
		TableMappings: pulumi.String(`{
			"rules": [
				{
					"rule-type": "selection",
					"rule-id": "1",
					"rule-name": "1",
					"object-locator": {
						"schema-name": "%",
						"table-name": "%"
					},
					"rule-action": "include"
				}
			]
		}`),
		Tags: pulumi.StringMap{
			"Name": pulumi.String("cdc-lab-task"),
		},
	})
	if err != nil {
		return nil, err
	}

	return &DmsMysqlPostgresOutput{
		Mariadb:             mariadb,
		Postgres:            postgres,
		ReplicationInstance: replicationInstance,
		ReplicationTask:     replicationTask,
	}, nil
}
