package types

type NotFoundError struct{}

func (error *NotFoundError) Error() string {
	return "not found"
}
