package predict

import (
	"net/http"
)

type EngineClient struct {
	client
}

func NewEngineClient(path string) *EngineClient {
	return &EngineClient{
		client{
			http.Client{},
			path,
		},
	}
}

func (cl *EngineClient) GetStatus() {

}

func (cl *EngineClient) GetStatusString() {

}

func (cl *EngineClient) SendQuery() {

}
