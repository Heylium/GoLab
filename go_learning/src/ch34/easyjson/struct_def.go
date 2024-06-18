package easyjson

/**
# for Go >= 1.17
go get github.com/mailru/easyjson && go install github.com/mailru/easyjson/...@latest
#then run
easyjson -all <file>.go
*/

type BasicInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type JobInfo struct {
	Skills []string `json:"skills"`
}

type Employee struct {
	BasicInfo BasicInfo `json:"basic_info"`
	JobInfo   JobInfo   `json:"job_info"`
}
