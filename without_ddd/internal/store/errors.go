package store

type ReleaseError struct {
	message             string
	UnavailableProducts []*Product
}

func NewReleaseError(message string, unavailableProducts []*Product) *ReleaseError {
	return &ReleaseError{
		message:             message,
		UnavailableProducts: unavailableProducts,
	}
}

func (e *ReleaseError) Error() string {
	return e.message
}
