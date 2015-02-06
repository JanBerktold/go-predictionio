package predict

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type client struct {
	cl     http.Client
	path   string
	appkey string
}

func (c *client) execute(method, url string, obj interface{}) (string, error) {
	by, enErr := json.Marshal(obj)

	if enErr != nil {
		return "", enErr
	}

	fmt.Println(string(by))
	req, _ := http.NewRequest(method, url, bytes.NewReader(by))
	req.Header.Add("Content-Type", AppJsonContentType)
	resp, hErr := c.cl.Do(req)

	if hErr != nil {
		return "", hErr
	}

	by, err := ioutil.ReadAll(resp.Body)

	if err == nil {
		return string(by), nil
	} else {
		return "", err
	}
}

func (c *client) GetStatus() {

}
