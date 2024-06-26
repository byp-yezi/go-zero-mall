// Code generated by goctl. DO NOT EDIT.
package types

type CreateRequest struct {
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Stock  int64  `json:"stock"`
	Amount int64  `json:"amount"`
	Status int64  `json:"status"`
}

type CreateResponse struct {
	Id int64 `json:"id"`
}

type DetailRequest struct {
	Id int64 `json:"id"`
}

type DetailResponse struct {
	Product Product `json:"product"`
}

type ListRequest struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
}

type ListResponse struct {
	Products []Product `json:"productlist"`
}

type Product struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Stock      int64  `json:"stock"`
	Amount     int64  `json:"amount"`
	Status     int64  `json:"status"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

type RemoveRequest struct {
	Id int64 `json:"id"`
}

type RemoveResponse struct {
}

type UpdateRequest struct {
	Id     int64  `json:"id"`
	Name   string `json:"name,optional"`
	Desc   string `json:"desc,optional"`
	Stock  int64  `json:"stock"`
	Amount int64  `json:"amount,optional"`
	Status int64  `json:"status,optional"`
}

type UpdateResponse struct {
}
