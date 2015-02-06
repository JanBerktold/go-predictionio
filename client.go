package predict

import (
	"net/http"
)

type client struct {
	http.Client
	path string
}
