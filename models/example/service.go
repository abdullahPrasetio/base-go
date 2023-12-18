package example

import (
	"context"
)

/********************************************************************************
* Temancode Example Service Package                                             *
*                                                                               *
* Version: 1.0.0                                                                *
* Date:    2023-01-05                                                           *
* Author:  Waluyo Ade Prasetio                                                  *
* Github:  https://github.com/abdullahPrasetio                                  *
********************************************************************************/

type service struct {
	repository Repository
}
type Service interface {
	Create(input RequestEmployee) (Employee, error)
	GetAll() ([]Employee, error)
	GetFromApi(ctx context.Context) ([]Employee, error)
}

func NewService(repo Repository) Service {
	return &service{repository: repo}
}

func (s *service) Create(input RequestEmployee) (Employee, error) {
	var err error
	var result Employee
	result, err = s.repository.Create(input)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *service) GetAll() ([]Employee, error) {
	var err error
	var results []Employee
	results, err = s.repository.GetAll()
	if err != nil {
		panic(err)
		return results, err
	}
	return results, nil
}

type ResponseBodyData struct {
	ResponseCode  string     `json:"responseCode"`
	ResponseDesc  string     `json:"responseDesc"`
	ResponseData  []Employee `json:"responseData"`
	ResponseError string     `json:"responseError"`
}

func (s *service) GetFromApi(ctx context.Context) ([]Employee, error) {
	//var err error
	//var results []Employee

	var responseBody ResponseBodyData

	//headers := []http.Headers{
	//	{
	//		Key:   "Content-Type",
	//		Value: "Application/json",
	//	},
	//}
	//reqBody := []byte{}
	//
	//client := &http2.Client{Timeout: 3 * time.Second}
	//res, respHeaders, err := http.Client_Req(ctx, client, headers, "https://mocki.io/v1/40d9df6f-5599-444d-99f0-815529ccae18", "GET", reqBody)
	//fmt.Println(respHeaders)
	//if err != nil {
	//	return results, err
	//}
	//
	//err = json.Unmarshal(res, &responseBody)

	return responseBody.ResponseData, nil
}
