package monitor

import (
	"net/http"
	"time"
)

func WebAdvRequest(url string, expCode int, timeout time.Duration, retries int) bool {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100

	//httpClient = &http.Client{
	//	Timeout:   countTimeout,
	//	Transport: t,
	//}

	client := &http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get("http://example.com")
	// ...

	req, err := http.NewRequest("GET", "http://example.com", nil)
	// ...
	//req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	println(resp)

	return false

}
