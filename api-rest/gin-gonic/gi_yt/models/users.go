package models

type Address struct {
	State   string `json:"state"`
	City    string `json:"city" `
	Pincode int    `json:"pincode"`
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age" `
}
