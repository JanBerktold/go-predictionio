package predict

import (
	"testing"
)

const (
	test_url = "http://localhost"
	test_key = "testkey"
)

func TestBasicClient(t *testing.T) {
	client := NewEventClient(test_url, test_key)
	t.Log(client.eventPath)

	if client == nil {
		t.FailNow()
	}

	client.SetUser("abcad")
}
