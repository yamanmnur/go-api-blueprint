package data

type CatsData struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Race   string `json:"race"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}
