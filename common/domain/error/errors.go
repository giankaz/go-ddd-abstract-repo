package common

type IError struct {
	Field    string `json:"field,omitempty"`
	Message  string `json:"message,omitempty"`
	Cause    string `json:"cause,omitempty"`
	Solution string `json:"solution,omitempty"`
}

type ErrorStruct struct {
	Errors []IError `json:"errors,omitempty"`
}
