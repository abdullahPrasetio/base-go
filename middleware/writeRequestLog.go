package middleware

import (
	"bytes"
	"encoding/json"
	json2 "github.com/abdullahPrasetio/base-go/utils/json"
	"github.com/abdullahPrasetio/base-go/utils/log"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
)

func WriteRequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf, _ := ioutil.ReadAll(c.Request.Body)
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf)) //We have to create a new Buffer, because rdr1 will be read.

		x := map[string]any{}
		err := json.Unmarshal(readBodyByte(rdr1), &x)
		if err != nil {
			//c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			//return
		}
		c.Request.Body = rdr2
		log.LogRequest(c.Request.Method, c.FullPath(), json2.MinifyJson(x), c.Request.Header)
	}
}
func readBodyByte(reader io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	s := buf.Bytes()
	return s
}

func readBodyString(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	s := buf.String()
	return s
}
