package customerrors

type WrongPasswordError struct {
}

func (e WrongPasswordError) Error() string {
	return "Wrong password or username"
}
