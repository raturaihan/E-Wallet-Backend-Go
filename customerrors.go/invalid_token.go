package customerrors

type InvalidTokenError struct {
}

func (e InvalidTokenError) Error() string {
	return "Invalid Request Token"
}
