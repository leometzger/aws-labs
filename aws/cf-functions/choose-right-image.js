/**
 * Cloudfron Function Handler
 *
 * @param {*} event
 * @returns
 */
function handler(event) {
  var request = event.request;
  var headers = request.headers;

  // Ignores when user-agent header was not sent
  if (typeof headers["user-agent"] === "undefined") {
    return request;
  }

  // Ignores when it is not a image resource
  if (!isImageResource(request.uri)) {
    return request;
  }

  if (isMobileAgent(headers["user-agent"].value)) {
    var extension = getExtension(request.uri);
    var replacer = new RegExp(extension + "$");

    request.uri = request.uri.replace(replacer, "_mobile" + extension);
  }

  return request;
}

/**
 * @param {string} uri
 * @returns
 */
function getExtension(uri) {
  var regex = /(?P<filename>[a-zA-Z1-9]+)(?P<extension>\.(jpeg|jpg))$/;
  var result = regex.exec(uri);

  return result.groups && result.groups.extension;
}

/**
 * Verify if a path is an image or not
 *
 * @param {string} uri
 * @returns {boolean}
 */
function isImageResource(uri) {
  return /\.(jpg|jpeg)$/.test(uri);
}

/**
 *
 * @param {*} userAgent
 * @returns {boolean}
 */
function isMobileAgent(userAgent) {
  return (
    /Android/i.test(userAgent) ||
    /webOS/i.test(userAgent) ||
    /iPhone/i.test(userAgent) ||
    /iPad/i.test(userAgent) ||
    /iPod/i.test(userAgent) ||
    /BlackBerry/i.test(userAgent) ||
    /Windows Phone/i.test(userAgent)
  );
}
