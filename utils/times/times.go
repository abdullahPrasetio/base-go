package times

/********************************************************************************
* Temancode Times Package                       		  	                    *
*                                                                               *
* Version: 1.0.0                                                                *
* Date:    2023-01-05                                                           *
* Author:  Waluyo Ade Prasetio                                                  *
* Github:  https://github.com/abdullahPrasetio                                  *
********************************************************************************/

import (
	"time"

	"github.com/abdullahPrasetio/base-go/configs"
)

type timeNow struct {
	ValueTime   time.Time
	ValueString string
}

func GetTimeNow(format string) timeNow {
	config := configs.Configs
	defaultformat := "2006-01-02 15:04:05"
	if len(format) > 0 {
		defaultformat = format
	}
	location, _ := time.LoadLocation(config.TIME_LOCATION)
	currentTime := time.Now().In(location)
	currentTimeFormatted := currentTime.Format(defaultformat)
	return timeNow{currentTime, currentTimeFormatted}
}
