package predict

import (
	"testing"
)

const (
	test_url = "localhost"
	test_key = "testkey"
)

func TestBasicClient(t *testing.T) {
	client := NewEventClient(test_url, test_key)

	if client == nil {
		t.FailNow()
	}

}
