package read

type Response struct {
	Id        string
	Email     string
	Hotelname string
	Password  string
	Verified  bool
}

func NewResponse(id, email, hotelname, password string, verified bool) *Response {
	return &Response{
		Id:        id,
		Email:     email,
		Hotelname: hotelname,
		Password:  password,
		Verified:  verified,
	}
}
