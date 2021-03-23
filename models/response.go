package models

type Response struct {
	CveError   int         `json:"Cve_Error"`
	CveMensaje interface{} `json:"Cve_Mensaje"`
}
