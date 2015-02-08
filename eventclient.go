package predict

import (
	"fmt"
	"net/http"
	"time"
)

type EventClient struct {
	client
	eventPath string
}

func NewEventClient(path, appkey string) *EventClient {
	return &EventClient{
		client{
			http.Client{},
			path,
			appkey,
		},
		fmt.Sprintf("%s/events.json?accessKey=%s", path, appkey),
	}
}

func (cl *EventClient) createEvent(eventType, entityType, entityId string, optional OptionalMap) {
	body, err := cl.execute("POST", cl.eventPath, Event{
		eventType,
		entityType,
		entityId,
		optional,
		time.Now(),
	})
	fmt.Println(body, err)
}

func (cl *EventClient) DeleteItem(iid string) {
	cl.createEvent("$delete", "item", iid, nil)
}

func (cl *EventClient) DeleteUser(uid string) {
	cl.createEvent("$delete", "user", uid, nil)
}

func (cl *EventClient) RecordUserActionOnItem(action, uid, iid string, optional ...OptionalMap) {
	var param OptionalMap
	if len(optional) > 0 {
		param = optional[0]
	} else {
		param = make(OptionalMap)
	}

	param["targetEntityType"] = "item"
	param["targetEntityType"] = iid

	cl.createEvent(action, "item", iid, param)
}

func (cl *EventClient) SetItem(iid string, optional ...OptionalMap) {
	var param OptionalMap
	if len(optional) > 0 {
		param = optional[0]
	}
	cl.createEvent("$set", "item", iid, param)
}

func (cl *EventClient) UnsetItem(iid string, optional OptionalMap) error {
	if len(optional) == 0 {
		return ErrEmptyParam
	}
	cl.createEvent("$unset", "item", iid, optional)
	return nil
}

func (cl *EventClient) SetUser(uid string, optional ...OptionalMap) {
	var param OptionalMap
	if len(optional) > 0 {
		param = optional[0]
	}
	fmt.Println(uid, param)
	cl.createEvent("$set", "user", uid, param)
}

func (cl *EventClient) UnsetUser(uid string, optional OptionalMap) error {
	if len(optional) == 0 {
		return ErrEmptyParam
	}
	cl.createEvent("$unset", "user", uid, optional)
	return nil
}
