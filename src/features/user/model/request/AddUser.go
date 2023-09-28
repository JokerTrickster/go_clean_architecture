package request

type ReqAddUser struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Age     int32  `json:"age"`
	Country string `json:"country"`
	Gender  string `json:"gender"`
}
