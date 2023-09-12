'use strict';

function getExtension(uri) {
	if (uri.endsWith('webp')) {
		return 'webp'
	} else if (uri.endsWith('jpeg')) {
		return 'jpeg'
	} else {
		return 'png'
	}
}

function isImageResource(uri) {
	return /\.(jpg|jpeg|webp)$/.test(uri);
}

function isMobileAgent(headers) {
	return headers['cloudfront-is-mobile-viewer']
		&& headers['cloudfront-is-mobile-viewer'][0].value === 'true'
		&& headers['cloudfront-is-tablet-viewer']
		&& headers['cloudfront-is-tablet-viewer'][0].value === 'false';
}

exports.handler = (event, _, callback) => {
	const request = event.Records[0].cf.request;
	const headers = request.headers;


	if (!isImageResource(request.uri)) {
		callback(null, request);
		return;
	}

	if (isMobileAgent(headers)) {
		var extension = getExtension(request.uri);
		var replacer = new RegExp('.' + extension + "$");

		request.uri = request.uri.replace(replacer, "_mobile." + extension);
	}

	callback(null, request);
};

