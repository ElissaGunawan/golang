package entities

type Product struct {
	Id       uint
	Name     string
	Category Category
	Brand    string
	Price    uint
	Status   string
}
