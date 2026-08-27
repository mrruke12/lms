package apperr

type Type string

const (
	ApplicationError    Type = "application"
	DomainError         Type = "domain"
	InfrastructureError Type = "infrastructure"
)

type AppError struct {
	Type    Type
	Code    string
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Message + ": " + e.Err.Error()
	}

	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}
