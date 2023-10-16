package model

// HttpResponse : success flag with payload is essential for
// easy Client-Side Handling, no dependency on http codes
// easy error handling
type HttpResponse struct {
	HttpResponseData
	MetaData RequestMetaData
}

type Response struct {
	Success bool
	Status  int
	Message string
	ResponseData
}

type HttpResponseData struct {
	Success bool
	Status  int
	Message string
	ResponseData
}

type ResponseData struct {
	Data          interface{}
	DataArray     []interface{}
	HasPagination bool
	Pagination    Pagination
}
type Pagination struct {
	TotalItems int `json:"totalItems"`
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
}

type RequestMetaData struct {
	ID         string
	IP         string
	URL        string
	HttpMethod string
	UserAgent  string

	StatusCode int

	StartEpoch          int64
	LatencyInNanoSecond int64
	ServiceID           int64
	TracingID           string
	ServiceName         string
	Query               map[string][]string
}
