package web

type CustomerInsertRequest struct {
	Name     string `json:"name" validate:"required,max=200,min=1"`
	Email    string `json:"email" validate:"required,max=200,min=1,contains=@gmail.com"`
	Gender   string `json:"gender" validate:"required"`
	Address  string `json:"address" validate:"required"`
	Password string `json:"password" validate:"required,alphanum,min=8"`
}

type CustomerUpdateRequest struct {
	Id       int    `validate:"required"`
	Name     string `validate:"required,max=200,min=1"`
	Email    string `validate:"required,max=200,min=1,contains=@gmail.com"`
	Gender   string `validate:"required"`
	Address  string `validate:"required"`
	Password string `validate:"required,alphanum,min=8"`
}

type CustomerResponse struct {
	Id       int
	Name     string
	Email    string
	Gender   string
	Address  string
	Password string
}
