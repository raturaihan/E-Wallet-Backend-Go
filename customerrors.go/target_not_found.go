package customerrors

type TargetNotFoundError struct {
}

func (e TargetNotFoundError) Error() string {
	return "Target wallet not found"
}
