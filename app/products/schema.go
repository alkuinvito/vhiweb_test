package products

type CreateProductRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=32"`
	Description string `json:"description" validate:"required,min=10,max=200"`
	Price       int    `json:"price" validate:"numeric,min=0,max=1000000"`
}

type GetProductResponse struct {
	ID          string
	Name        string
	Description string
	Price       int
	VendorID    string
}

type UpdateProductRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=32"`
	Description string `json:"description" validate:"required,min=10,max=200"`
	Price       int    `json:"price" validate:"numeric,min=0,max=1000000"`
}
