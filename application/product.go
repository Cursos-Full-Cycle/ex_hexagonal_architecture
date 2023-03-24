package application

import (
	"errors"	
)


type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetId() string
	GetName() string
	GetStatus() string
	GetPricec() float64
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"	
)

type Product struct {
	Id     string 
	Name   string 
	Status string 
	Price  float64
}

func (p *Product) IsValid() (bool, error) {
	return true, nil
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}
	return errors.New("the price must be greater than zero to enable the product")
}

func (p *Product) Disable() error {
	if p.Status == DISABLED {
		return errors.New("product already disabled")
	}
	
	p.Status = DISABLED
	return nil
}

func (p *Product) GetId() string {
	return p.Id
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}