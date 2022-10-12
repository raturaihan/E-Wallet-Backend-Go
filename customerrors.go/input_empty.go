package customerrors

type InputEmptyError struct {
}

func (e InputEmptyError) Error() string {
	return "Name, email, and password can't be empty"
}
