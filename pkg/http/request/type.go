package request

/*
Host → Specifies the host and port of the server.
User-Agent → Identifies the client making the request.
Accept → Specifies the media types the client can process.
Accept-Encoding → Specifies the content encodings the client supports (e.g., gzip, deflate).
Accept-Language → Specifies the preferred language of the response.
Referer → Specifies the URL of the page that made the request.
Connection → Controls whether the connection is kept open (keep-alive) or closed.
*/
type HTTPGeneralHeaders struct {
	Host           string
	UserAgent      string
	Accept         string
	AcceptEncoding string
	RemoteAddr     string
	RequestURI     string
	Referer        string
	Connection     string
}

/*
Authorization → Contains authentication credentials (e.g., Bearer <token> or Basic <credentials>).
Proxy-Authorization → Authentication for a proxy server.
Cookie → Contains cookies sent by the client.
*/
type HTTPSecurityHeaders struct {
	IsAuth             bool
	Authorization      string
	ProxyAuthorization string
	Cookie             string
}

/*
Content-Type → Specifies the media type of the request body (e.g., application/json, multipart/form-data).
Content-Length → Specifies the length of the request body in bytes.
*/
type HTTPContentRelatedHeaders struct {
	ContentType   string
	ContentLength int64
}

// UserAccount holds data of basic user account possible data
/*
	The provided data as follows:
	1. ID (string). Assuming that the account id saved in the database is string. If no data provided, default value of "" used.
	2. IDAutoIncr (uint64). Assuming that the database only provide auto incr id, use this instead. If no data provided, default value of 0 used.
	3. Email (string). Assuming there is email data provided for the account. If no data provided, default value of "" used.
	3. Username (string). Assuming there is username data provided for the account. If no data provided, this will be filled with email data as fallback.
*/
type UserAccount struct {
	ID         string
	IDAutoIncr uint64
	Email      string
	Username   string
}
