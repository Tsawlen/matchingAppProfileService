package dataStructures

type Profile struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	Id          int    `json:"id"`
	City        string `json:"city"`
	Email       string `json:"email"`
	First_name  string `json:"first_name"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	Street      string `json:"street"`
	HouseNumber string `json:"houseNumber"`
	Username    string `json:"username"`
}
