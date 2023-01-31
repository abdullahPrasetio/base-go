package http

/********************************************************************************
* Temancode Example Request Http Package                                        *
*                                                                               *
* Version: 1.0.0                                                                *
* Date:    2023-01-05                                                           *
* Author:  Waluyo Ade Prasetio                                                  *
* Github:  https://github.com/abdullahPrasetio                                  *
********************************************************************************/

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/abdullahPrasetio/base-go/configs"
	"github.com/abdullahPrasetio/base-go/utils/log"
)

type Headers struct {
	Key   string
	Value string
}

func Client_Req(ctx context.Context, h []Headers, uri string, m string, d []byte) ([]byte, http.Header, error) {
	config := configs.Configs
	client := &http.Client{Timeout: time.Second * time.Duration(config.TIMEOUT_REQ)}
	logger := log.Logger
	logger.Info(string(d))

	refnum, _ := ctx.Value("refnum").(string)
	req, err := http.NewRequest(m, uri, bytes.NewReader(d))
	if err != nil {
		return nil, nil, err
	}

	if h == nil {
		req.Header.Set("Content-Type", "application/json")
	} else {
		req.Header.Set("Content-Type", "application/json")
		//header tambahan
		for i := range h {
			req.Header[h[i].Key] = []string{h[i].Value}
		}
	}

	reqHeader, _ := json.Marshal(req.Header)
	log.LogDebug(refnum, "URL = "+uri)
	log.LogDebug(refnum, "request header = "+string(reqHeader))
	log.LogDebug(refnum, "request Body = "+readBodyString(bytes.NewReader(d)))

	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(err)
	if err != nil {
		return nil, nil, err
	}

	//check response status
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode >= http.StatusInternalServerError {
			return nil, nil, errors.New(fmt.Sprintf("%d error: %s", resp.StatusCode, resp.Status))
		}
		return nil, nil, fmt.Errorf("%d error: %s", resp.StatusCode, resp.Status)
	}

	respHeader, _ := json.Marshal(resp.Header)

	log.LogDebug(refnum, "response header = "+string(respHeader))
	log.LogDebug(refnum, "response body = "+string(body))

	logger.Info(string(body))

	return body, resp.Header, nil
}

func readBodyString(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	s := buf.String()
	return s
}
