package dblayer

import (
	"errors"
	"hands-on/RestfulServer/backend/src/models"
)

type DBLayer interface {
	GetAllProducts() ([]models.Product, error)
	GetPromos() ([]models.Product, error)
	GetCustomerByName(string, string) (models.Customer, error)
	GetCustomerById(int) (models.Customer, error)
	GetProduct(int) (models.Product, error)
	AddUser(models.Customer) (models.Customer, error)
	SignInUser(username, password string) (models.Customer, error)
	SignOutUserById(int) error
	GetCustomerOrdersByID(int) ([]models.Order, error)
	AddOrder(models.Order) error
	GetCreditCardID(int) (string, error)
	SaveCreditCardForCustomer(int, string) error
}

var ErrINVALIDPASSWORD = errors.New("invalid password")
