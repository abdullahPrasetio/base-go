package example

/********************************************************************************
* Temancode Example Input Struct Package                                        *
*                                                                               *
* Version: 1.0.0                                                                *
* Date:    2023-01-05                                                           *
* Author:  Waluyo Ade Prasetio                                                  *
* Github:  https://github.com/abdullahPrasetio                                  *
********************************************************************************/

type RequestGetEmploye struct {
	Name string `json:"name"`
}

type RequestEmployee struct {
	Name    string `json:"name" binding:"required"`
	Age     int    `json:"age" binding:"required"`
	Address string `json:"address" binding:"required"`
}

type GetCampaignDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
