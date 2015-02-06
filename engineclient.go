package predict

import (
	"net/http"
)

type EngineClient struct {
	client
}

func NewEngineClient(path, appkey string) *EngineClient {
	return &EngineClient{
		client{
			http.Client{},
			path,
			appkey,
		},
	}
}

func (cl *EngineClient) GetStatus() {

}

func (cl *EngineClient) GetStatusString() {

}

func (cl *EngineClient) SendQuery() {

}
