package entity

import (
	"errors"
	"gokomodo/pkg/common"
	"time"

	"github.com/hashicorp/go-multierror"
)

type Product struct {
	Id          common.ID
	Name        string
	Description string
	Price       int
	Seller      *User
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ProductDTO struct {
	Id          *common.ID
	Name        string
	Description string
	Price       int
	SellerId    common.ID
}

func NewProduct(productDTO *ProductDTO) (*Product, *multierror.Error) {
	var multierr *multierror.Error

	if productDTO.Id == nil {
		id := common.NewID()
		productDTO.Id = &id
	}

	product := &Product{
		Id:          *productDTO.Id,
		Name:        productDTO.Name,
		Description: productDTO.Description,
		Price:       productDTO.Price,
		Seller: &User{
			Id: productDTO.SellerId,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if errValidate := product.Validate(); errValidate != nil {
		multierr = multierror.Append(multierr, errValidate)
	}

	if multierr != nil {
		return nil, multierr
	}

	return product, nil
}

func (product *Product) Validate() *multierror.Error {
	var multierr *multierror.Error

	if product.Name == "" {
		multierr = multierror.Append(multierr, errors.New("product name cannot be empty"))
	}

	if product.Description == "" {
		multierr = multierror.Append(multierr, errors.New("product description cannot be empty"))
	}

	if product.Price == 0 {
		multierr = multierror.Append(multierr, errors.New("product price cannot be empty"))
	}

	return multierr
}
