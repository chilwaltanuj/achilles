package model

// HttpResponse : success flag with payload is essential for
// easy Client-Side Handling, no dependency on http codes
// easy error handling
type HttpResponse struct {
	HttpResponseData
	MetaData *RequestMetaData
}

type HttpResponseData struct {
	Success bool
	Status  int
	Message string
	*ResponseData
}

type ResponseData struct {
	Data          any         `json:"Data,omitempty"`
	DataArray     []any       `json:"DataArrary,omitempty"`
	HasPagination bool        `json:"HasPagination,omitempty"`
	Pagination    *Pagination `json:"Pagination,omitempty"`
}
type Pagination struct {
	TotalItems int `json:"totalItems,omitempty"`
	Offset     int `json:"offset,omitempty"`
	Limit      int `json:"limit,omitempty"`
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
	ApplicationID       int64
	TracingID           string
	Application         string
	Query               map[string][]string
}
