package http

/********************************************************************************
* Temancode Example Response Package                                            *
*                                                                               *
* Version: 1.0.0                                                                *
* Date:    2023-01-05                                                           *
* Author:  Waluyo Ade Prasetio                                                  *
* Github:  https://github.com/abdullahPrasetio                                  *
********************************************************************************/

type Response struct {
	ResponseCode   string      `json:"responseCode"`
	ResponseDesc   string      `json:"responseDesc"`
	ResponseData   interface{} `json:"responseData"`
	ResponseErrors interface{} `json:"responseErrors"`
}

func APIResponseSuccess(responseDesc string, responseData interface{}) Response {
	jsonResponse := Response{
		ResponseCode: "00",
		ResponseDesc: responseDesc,
		ResponseData: responseData,
	}

	return jsonResponse
}

func APIResponseError(responseDesc string, responseCode string, errors interface{}) Response {
	jsonResponse := Response{
		ResponseCode:   responseCode,
		ResponseDesc:   responseDesc,
		ResponseErrors: errors,
	}

	return jsonResponse
}
