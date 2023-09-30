package response

type ResGetsUser struct {
	Users []GetsUserInfo `json:"users"`
	Count int32          `json:"count"`
}

type GetsUserInfo struct {
	ID      string `json:"id"`
	Age     int32  `json:"age"`
	Country string `json:"country"`
	Name    string `json:"name"`
	Gender  string `json:"gender"`
}
