package predict

import (
	"net/http"
)

type client struct {
	cl   http.Client
	path string
}

func (c *client) execute() {

}
