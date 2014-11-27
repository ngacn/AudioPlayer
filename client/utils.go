package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func (cli *WarptenCli) HTTPClient() *http.Client {
	return &http.Client{Transport: cli.transport}
}

func (cli *WarptenCli) encodeData(data interface{}) (*bytes.Buffer, error) {
	params := bytes.NewBuffer(nil)
	if data != nil {
		buf, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		if _, err := params.Write(buf); err != nil {
			return nil, err
		}
	}
	return params, nil
}

// 发起请求
func (cli *WarptenCli) call(method, path string, data interface{}) (io.ReadCloser, int, error) {
	params, err := cli.encodeData(data)
	if err != nil {
		return nil, -1, err
	}
	req, err := http.NewRequest(method, path, params)
	if err != nil {
		return nil, -1, err
	}
	req.Header.Set("User-Agent", "Warpten-Client")
	req.URL.Host = cli.addr
	// 总是使用http协议
	req.URL.Scheme = "http"
	if data != nil {
		req.Header.Set("Content-Type", "application/json")
	} else if method == "POST" {
		req.Header.Set("Content-Type", "plain/text")
	}
	resp, err := cli.HTTPClient().Do(req)
	if err != nil {
		if strings.Contains(err.Error(), "connection refused") {
			return nil, -1, errors.New("Cannot connect to the WarptenSrv")
		}
		return nil, -1, fmt.Errorf("An error occurred trying to connect: %v", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, -1, err
		}
		if len(body) == 0 {
			return nil, resp.StatusCode, fmt.Errorf("Error: request returned %s for API route and version %s, check if the server supports the requested API version", http.StatusText(resp.StatusCode), req.URL)
		}
		return nil, resp.StatusCode, fmt.Errorf("Error response from daemon: %s", bytes.TrimSpace(body))
	}
	return resp.Body, resp.StatusCode, nil
}

func readBody(stream io.ReadCloser, statusCode int, err error) ([]byte, int, error) {
	if stream != nil {
		defer stream.Close()
	}
	if err != nil {
		return nil, statusCode, err
	}
	body, err := ioutil.ReadAll(stream)
	if err != nil {
		return nil, -1, err
	}
	return body, statusCode, nil
}
