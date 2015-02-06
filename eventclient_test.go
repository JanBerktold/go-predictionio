package predict

import (
	"testing"
)

const (
	test_url = "localhost"
)

func TestBasicClient(t *testing.T) {
	client := NewEventClient(test_url)

	if client == nil {
		t.FailNow()
	}

}
