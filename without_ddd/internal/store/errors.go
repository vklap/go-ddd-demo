package store

type ReleaseError struct {
	message             string
	unavailableProducts []*Product
}

func NewReleaseError(message string, unavailableProducts []*Product) *ReleaseError {
	return &ReleaseError{
		message:             message,
		unavailableProducts: unavailableProducts,
	}
}

func (e *ReleaseError) Error() string {
	return e.message
}

func (e *ReleaseError) UnavailableProducts() []*Product {
	return e.unavailableProducts
}
