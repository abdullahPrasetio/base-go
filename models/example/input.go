package example

type RequestGetEmploye struct {
	Name string `json:"name"`
}

type RequestEmployee struct {
	Name    string `json:"name" binding:"required"`
	Age     int    `json:"age" binding:"required"`
	Address string `json:"address" binding:"required"`
	//Date    string `json:"date" binding:"required,date=yyyy-mm-dd"`
}

//type RequestUpdateEmployee struct {
//	Name    string `json:"name" binding:"required"`
//	Age     int    `json:"age" binding:"required"`
//	Address string `json:"address" binding:"required"`
//}

type GetCampaignDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
