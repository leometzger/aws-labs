FROM public.ecr.aws/lambda/provided:al2 as build

RUN yum install -y golang
RUN go env -w GOPROXY=direct

ADD go.mod go.sum ./
RUN go mod download

ADD ../lambda .
RUN go build -o /main

# copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2
COPY --from=build /main /main

ENTRYPOINT [ "/main" ]           
