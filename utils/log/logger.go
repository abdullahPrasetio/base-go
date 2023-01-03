package log

import (
	"fmt"
	"github.com/abdullahPrasetio/base-go/constants"
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"os"
	"time"

	"github.com/abdullahPrasetio/base-go/configs"
	param "github.com/abdullahPrasetio/base-go/models/log"
	log "github.com/sirupsen/logrus"
)

var (
	serviceName string
	folder      = folderDev
	logText     *log.Logger
	logJSON     *log.Logger
	Logger      *log.Logger
)

func init() {
	setText()
	setJSON()

	config := configs.Configs
	if config.DEBUG {
		logText.SetLevel(log.DebugLevel)
		logJSON.SetLevel(log.DebugLevel)
	} else {
		logText.SetLevel(log.InfoLevel)
		logJSON.SetLevel(log.InfoLevel)
	}

	if config.APP_ENV != "development" {
		folder = folderOcp
	}

	serviceName = constants.ServerName
}

func LoadLogger() {
	Logger = log.New()

	switch os.Getenv("LOG_LEVEL") {
	case "trace":
		Logger.SetLevel(log.TraceLevel)
	case "debug":
		Logger.SetLevel(log.DebugLevel)
	case "info":
		Logger.SetLevel(log.InfoLevel)
	case "warn":
		Logger.SetLevel(log.WarnLevel)
	case "error":
		Logger.SetLevel(log.ErrorLevel)
	case "fatal":
		Logger.SetLevel(log.FatalLevel)
	case "panic":
		Logger.SetLevel(log.PanicLevel)
	}

	Logger.SetFormatter(&log.JSONFormatter{})
	Logger.Info("logger Service transaction successfully configured")
}

const (
	httpRequest  = "REQUEST"
	httpResponse = "RESPONSE"
	timeformat   = "2006-01-02T15:04:05-0700"
	nameformat   = "log-2006-01-02.log"
	folderDev    = "log/"
	folderOcp    = "logs/"
)

func LogRequest(username, trx_type string, request interface{}, header http.Header) {
	timestamp := setLogFile()
	logJSON.WithFields(log.Fields{
		"service":        serviceName,
		"http_type":      httpRequest,
		"request_header": header,
		"request_body":   request,
		"trx_type":       trx_type,
		"username":       username,
		"timestamp":      timestamp,
	}).Info(httpRequest)

}

func setLogFile() string {
	currentTime := time.Now()
	timestamp := currentTime.Format(timeformat)
	filename := folder + currentTime.Format(nameformat)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	} else {
		logText.SetOutput(file)
		logJSON.SetOutput(file)
	}
	return timestamp
}

func setJSON() {
	logJSON = log.New()
	formatter := new(log.JSONFormatter)
	formatter.DisableTimestamp = true
	logJSON.SetFormatter(formatter)
}

func LogResponse(param param.LogResponseParam) {
	timestamp := setLogFile()
	logJSON.WithFields(log.Fields{
		"username":        param.Username,
		"third_party":     param.ThirdParty,
		"service":         serviceName,
		"http_type":       httpResponse,
		"response_header": param.ResponseHeader,
		"response_body":   minifyResponse(param.ResponseBody),
		"response_code":   param.ResponseCode,
		"timestamp":       timestamp,
	}).Info(httpResponse)
}

func setText() {
	logText = log.New()
	formatter := new(log.TextFormatter)
	formatter.DisableTimestamp = true
	formatter.DisableQuote = true
	logText.SetFormatter(formatter)
}

func minifyResponse(r interface{}) map[string]interface{} {
	js, _ := jsoniter.Marshal(r)
	var m map[string]interface{}
	_ = jsoniter.Unmarshal(js, &m)
	for k, v := range m {
		s := fmt.Sprintf("%v", v)
		if len(s) > 10000 {
			delete(m, k)
		}
	}
	return m
}
