package controllers

// Data contains the Type of the request and the context
type Data struct {
	Type    string      `json:"type,omitempty"`
	Context interface{} `json:"context,omitempty"`
}

// Datas is a slice of HATEOAS datas
type Datas []Data

// Wrapper is the HATEOAS wrapper
type Wrapper struct {
	Datas  *Datas  `json:"datas,omitempty"`
	Errors *Errors `json:"errors,omitempty"`
}
