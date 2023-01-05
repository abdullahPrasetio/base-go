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

func Client_Req(h []Headers, uri string, m string, d []byte) ([]byte, error) {
	config := configs.Configs
	client := &http.Client{Timeout: time.Second * time.Duration(config.TIMEOUT_REQ)}
	logger := log.Logger
	logger.Info(string(d))
	req, err := http.NewRequest(m, uri, bytes.NewReader(d))
	if err != nil {
		return nil, err
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

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	logger.Info(string(body))
	return body, nil
}
