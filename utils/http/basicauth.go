package http

/********************************************************************************
* Temancode Basic Auth Package                       		  	                *
*                                                                               *
* Version: 1.0.0                                                                *
* Date:    2023-01-05                                                           *
* Author:  Waluyo Ade Prasetio                                                  *
* Github:  https://github.com/abdullahPrasetio                                  *
********************************************************************************/
import "encoding/base64"

func BasicAuthString(username, password string) string {
	auth := username + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}
