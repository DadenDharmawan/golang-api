package web

import "mime/multipart"

type ProductInsertRequest struct {
	Name     string                  `json:"name" validate:"required,max=200,min=1"`
	Category string                  `json:"category" validate:"required,max=200,min=1"`
	Desc     string                  `json:"desc"`
	Price    int                     `json:"price" validate:"required"`
	Qty      int                     `json:"qty" validate:"required"`
	Image    []*multipart.FileHeader `json:"image"`
}

type ProductUpdateRequest struct {
	Id       int                     `json:"id" validate:"required,max=200,min=1"`
	Name     string                  `json:"name" validate:"required,max=200,min=1"`
	Category string                  `json:"category"`
	Desc     string                  `json:"desc" validate:"required"`
	Price    int                     `json:"price" validate:"required"`
	Qty      int                     `json:"qty" validate:"required"`
}

type ProductResponse struct {
	Id       int
	Name     string
	Category string
	Desc     string
	Price    int
	Qty      int
	Image    string
}
