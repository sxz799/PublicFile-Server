package model

type Result struct {
	Status  string  `json:"status"`
	Success bool    `json:"success"`
	Message string  `json:"message"`
	FileObj FileObj `json:"fileObj"`
}
