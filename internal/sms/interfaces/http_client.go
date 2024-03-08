package interfaces

type HTTPClient interface {

	//Get performs a GET request
	Get(url string) <-chan *HTTPResponse
	//Post performs a POST request
	Post(url string, body []byte) <-chan *HTTPResponse
}

type HTTPResponse struct {
	StatusCode int
	Body       []byte
	Error      error
}
