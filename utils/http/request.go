package http

import (
	"bytes"
	"github.com/abdullahPrasetio/base-go/configs"
	"github.com/abdullahPrasetio/base-go/utils/log"
	"io/ioutil"
	"net/http"
	"time"
)

type Headers struct {
	Key   string
	Value string
}

func Client_Req(h []Headers, uri string, m string, d []byte) ([]byte, error) {
	//btr_timeout, _ := strconv.Atoi(os.Getenv("BTR_TO"))
	config := configs.Configs
	client := &http.Client{Timeout: time.Second * time.Duration(config.TIMEOUT_REQ)}
	logger := log.Logger
	//l.Info("REQ", zap.Any("body", string(d)))
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

	//reqDump, err := httputil.DumpRequestOut(req, true)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//bodies, err := ioutil.ReadAll(req.Body)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Printf("%s", bodies)
	//
	//fmt.Printf("REQUEST:\n%s", string(reqDump))
	//fmt.Println()

	//log.Println(req.Header)
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
