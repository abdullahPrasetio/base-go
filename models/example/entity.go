package example

type Employee struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Address   string `json:"address"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
