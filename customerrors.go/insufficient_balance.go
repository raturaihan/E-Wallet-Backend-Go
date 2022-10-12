package customerrors

type InsufficientBalanceError struct {
}

func (e InsufficientBalanceError) Error() string {
	return "Insufficient Balance"
}
