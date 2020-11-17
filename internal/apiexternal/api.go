package apiexternal

import (
	"io/ioutil"
	"net"
	"net/http"
	"paint/internal/apiexternal/cache"
	"paint/internal/apiexternal/cache/memory"
	"time"
)

const (
	cachedTime  = 5 * time.Minute
	clearedTime = 30 * time.Minute
	baseURL     = ""
)

type Api interface {
	externalApiTest()
}

type api struct {
	client  *http.Client
	storage cache.Storage
}

func (c *api) req(params map[string]string) ([]byte, error) {

	req, err := c.client.Get(baseURL)

	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *api) externalApiTest() {

}

func NewAlphaVantage() Api {
	defaultTransport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:        20,
		MaxIdleConnsPerHost: 20,
		TLSHandshakeTimeout: 15 * time.Second,
	}

	client := &http.Client{
		Transport: defaultTransport,
		Timeout:   15 * time.Second,
	}

	return &api{
		client:  client,
		storage: memory.InitCash(cachedTime, clearedTime),
	}
}
