package ocs

type ocs struct {
	OCSMeta ocsMeta `json:"meta"`
}

type ocsMeta struct {
	Status     string `json:"status"`
	StatusCode string `json:"statuscode"`
	Message    string `json:"message"`
}
