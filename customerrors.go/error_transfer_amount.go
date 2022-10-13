package customerrors

type ErrorTransferAmount struct {
}

func (e ErrorTransferAmount) Error() string {
	return "Minimum Transfer 1000 and Maximum 50.000.000"
}
