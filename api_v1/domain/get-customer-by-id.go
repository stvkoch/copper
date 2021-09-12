package domain

import (
	"github.com/stvkoch/copper/repository"
)

func GetCustomerByID(id int) (interface{}, error) {
	where := map[string]interface{}{"ID": id}

	customerRepo := repository.Customer{}
	customer, err := customerRepo.Get(where)

	return customer, err
}
