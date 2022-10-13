package customerrors

type TransferFailed struct {
}

func (e TransferFailed) Error() string {
	return "Can not transfer to your own wallet id"
}
