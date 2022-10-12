package customerrors

type DataAlreadyExistError struct {
}

func (e DataAlreadyExistError) Error() string {
	return "Email already registered"
}
