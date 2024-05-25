package apperror

type ApplicationError struct {
	Status  int
	Message string
}

func (e *ApplicationError) Error() string {
	return e.Message
}

func New(status int, err error) error {
	return &ApplicationError{
		Status:  status,
		Message: err.Error(),
	}
}
