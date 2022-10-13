package customerrors

type ErrorTopUp struct {
}

func (e ErrorTopUp) Error() string {
	return "Minimum Top Up 50.000 and Maximum 10.000.000"
}
