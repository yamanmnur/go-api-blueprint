package responses

type MetaData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

// GenericResponse struct
type GenericResponse struct {
	MetaData MetaData    `json:"meta_data"`
	Data     interface{} `json:"data"`
}

// BasicResponse struct
type BasicResponse struct {
	MetaData MetaData `json:"meta_data"`
}
