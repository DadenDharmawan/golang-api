package entity

type Customer struct {
	Id       int
	Name     string
	Email    string
	Gender   string
	Address  string
	Password string
}

type Product struct {
	Id       int
	Name     string
	Category string
	Desc     string
	Price    int
	Qty      int
	Image    string
}
