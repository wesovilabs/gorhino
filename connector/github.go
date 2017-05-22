package connector

import (
	"bytes"
	"net/http"
	"io/ioutil"
)

type GitHub struct {
	repository	string
	branch 		string
	resource 	string
}

const gitHubUrl = "https://raw.githubusercontent.com/"

//Configuration get Github configuration
func (gitHub GitHub) Configuration() string {
	var urlBuffer bytes.Buffer
	urlBuffer.WriteString(gitHubUrl)
	urlBuffer.WriteString(gitHub.repository)
	urlBuffer.WriteString("/")
	urlBuffer.WriteString(gitHub.branch)
	urlBuffer.WriteString("/")
	urlBuffer.WriteString(gitHub.resource)
	url := urlBuffer.String()
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	return string(body)
}

