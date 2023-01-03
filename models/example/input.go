package example

type RequestGetEmploye struct {
	Name string `json:"name"`
}

type RequestCreateEmployee struct {
	Name    string `json:"name" binding:"required"`
	Age     int    `json:"age" binding:"required"`
	Address string `json:"address" binding:"required"`
}
