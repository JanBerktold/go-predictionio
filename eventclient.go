package predict

import (
	"net/http"
)

type EventClient struct {
	client
}

func NewEventClient(path, appkey string) *EventClient {
	return &EventClient{
		client{
			http.Client{},
			path,
			appkey,
		},
	}
}

func (cl *EventClient) CreateEvent() {

}

func (cl *EventClient) DeleteItem() {

}

func (cl *EventClient) DeleteUser(uid string) {

}

func (cl *EventClient) RecordUserActionOnItem() {

}

func (cl *EventClient) SetItem() {

}

func (cl *EventClient) SetUser(uid string) {

}

func (cl *EventClient) UnsetItem() {

}

func (cl *EventClient) UnsetUser() {

}
