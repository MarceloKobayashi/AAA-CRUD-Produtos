package main

import (
	"encoding/json"
	
	"strings"
	"errors"
	"fmt"
	"strconv"
	"math"
)

type ProductFactory struct{}

func NewProductFactory() *ProductFactory {
	return &ProductFactory{}
}

func (f *ProductFactory) CreateFromJSON(data []byte) (*Product, error) {
	var product Product
	
	if err := json.Unmarshal(data, &product); err != nil {
		return nil, err
	}
	
	return f.Validate(&product)
}


func (f *ProductFactory) Validate(product *Product) (*Product, error) {
	
	if err := f.validateName(product); err != nil {
		return product, err
	}
	
	if err := f.validateDescription(product); err != nil {
		return product, err
	}

	if err := f.validatePrice(product); err != nil {
		return product, err
	}
	
	if err := f.validateAmount(product); err != nil {
		return product, err
	}
	
	if err := f.validateWeight(product); err != nil {
		return product, err
	}
	
	f.ensureID(product)
	f.ensureDisponibility(product)
	
	return product, nil
}


func (f *ProductFactory) ensureID(product *Product) {
	if product.ID == "" {
		product.ID = GenerateStringID()
		return
	}
	
	if !ValidateID(product.ID) {
		product.ID = GenerateStringID()
		return
	}

	return
}

func (f *ProductFactory) ensureDisponibility(product *Product) {
	if product.Amount > 0 {
		product.Disponibility = true
		return
	} 
	
	product.Disponibility = false
	return
}

func (f *ProductFactory) validateName(product *Product) error{
	if product.Name == "" {
		return errors.New("invalid name: cannot be empty")
	}
	
	return nil
}

func (f *ProductFactory) validateDescription(product *Product) error{
	if len(product.Description) > 512 {
		product.Description = product.Description[:512]
	}

	return nil
}

func (f *ProductFactory) validatePrice(product *Product) error{
	precoFormatado := fmt.Sprintf("%.2f", product.Price)
	
	productPrice, err := strconv.ParseFloat(precoFormatado, 64)
	if err != nil {
		return errors.New("invalid price format: must be numbers")
	}
	
	product.Price = productPrice
	return nil
}

func (f *ProductFactory) validateAmount(product *Product) error {
	if product.Amount < 0 {
		return errors.New("Amount must be higher or equal than zero '0'.")
	}
	
	//segunda parte
	if product.Amount < math.MinInt32 || product.Amount > math.MaxInt32 {
		return errors.New("Amount out of bounds for int32")
	}
	
	productAmount := int32(product.Amount)
	if int(productAmount) != product.Amount {
		return errors.New("Invalid amount format: must be an integer")
	}
	
	product.Amount = int(productAmount)
	
	return nil
}

func (f *ProductFactory) validateWeight(product *Product) error{
	if !strings.HasSuffix(product.Weight, "kg") && !strings.HasSuffix(product.Weight, "g") {
		weight, err := strconv.Atoi(product.Weight)
		
		if err != nil {
			return errors.New("invalid weight: must be numbers")
		}
		
		var unit string
		if weight > 20 {
			unit = "g"
		} else {
			unit = "kg"
		}
		
		product.Weight += unit
	}
	
	pesoSemUnidade := strings.TrimRight(product.Weight, "kggGK")
	
	weight, err := strconv.Atoi(pesoSemUnidade)
	if err != nil {
		return errors.New("invalid weight: must be numbers")
	}
	
	if strings.HasSuffix(product.Weight, "kg") {
		product.Weight = strconv.Itoa(weight) + "kg"
	} else if strings.HasSuffix(product.Weight, "g") {
		product.Weight = strconv.Itoa(weight) + "g"
	}
	
	return nil
}

