package main

type Product struct {
	ID string		`json:"id" 		bson:"id"`
	Name string		`json:"name" 		bson:"name"`
	Price float64		`json:"price" 		bson:"price"`
	Description string	`json:"description" 	bson:"description"`
	Amount int		`json:"amount" 		bson:"amount"`
	Disponibility bool	`json:"disponibility" 	bson:"disponibility"`
	Weight string		`json:"weight" 		bson:"weight"`
	//LactoseFree bool	`json:"lactosefree" bson:"lactosefree"`
	//ValidityDate time.Time	`json:"validitydate" bson:"validitydate"`
}

func CreateProduct (id, name, description, weight string, amount int, price float64, disponibility bool) Product {
	return Product {
		ID: id,
		Name: name,
		Price: price,
		Description: description,
		Amount: amount,
		Disponibility: disponibility,
		Weight: weight,
	}
}
