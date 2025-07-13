package vendors

type RegisterVendorRequest struct {
	Name string `json:"name" validate:"required,min=3,max=32"`
}

type UpdateVendorRequest struct {
	Name string `json:"name" validate:"required,min=3,max=32"`
}

type GetVendorResponse struct {
	ID     string
	Name   string
	UserID string
}
