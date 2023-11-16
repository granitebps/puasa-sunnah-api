package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func SetupResty() *resty.Request {
	return resty.New().
		EnableTrace().
		SetDisableWarn(true).
		SetJSONMarshaler(json.Marshal).
		SetJSONUnmarshaler(json.Unmarshal).
		SetTransport(newrelic.NewRoundTripper(http.DefaultTransport)).
		R()
}

// This will return customize log message from client response
func LogClient(res *resty.Response) string {
	url := res.Request.URL
	status := res.StatusCode()
	method := res.Request.Method
	resTime := res.Time()
	var requestBody string
	if method == http.MethodPost || method == http.MethodPut {
		if res.Request.RawRequest.Body != nil {
			bodyReq, _ := io.ReadAll(res.Request.RawRequest.Body)
			requestBody = string(bodyReq)
		}
	}

	logMsg := "Requesting %s to %s with params %s. Returned (%d) with response body %s for %s"
	return fmt.Sprintf(logMsg, method, url, requestBody, status, string(res.Body()), resTime)
}
