package domain

import "time"

type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	LogoURL     string    `json:"logo_url"`
	AccountID   string    `json:"account_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProductUseCase interface {
	FetchProducts() ([]Product, error)
	GetProduct(productID string) (Product, error)
	CreateProduct(name, description, accountID string) error
	UploadLogo() error
	DeleteProduct(productID, accountID string) error
}