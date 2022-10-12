package customerrors

type AuthHeaderUnavailable struct {
}

func (e AuthHeaderUnavailable) Error() string {
	return "Authorization header is unavailable"
}
