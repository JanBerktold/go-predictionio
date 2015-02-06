package predict

import (
	"net/http"
)

type EventClient struct {
	client
}

func NewEventClient(path string) *EventClient {
	return &EventClient{
		client{
			http.Client{},
			path,
		},
	}
}

func (cl *EventClient) CreateEvent() {

}

func (cl *EventClient) DeleteItem() {

}

func (cl *EventClient) DeleteUser() {

}

func (cl *EventClient) RecordUserActionOnItem() {

}

func (cl *EventClient) SetItem() {

}

func (cl *EventClient) UnsetItem() {

}

func (cl *EventClient) UnsetUser() {

}
