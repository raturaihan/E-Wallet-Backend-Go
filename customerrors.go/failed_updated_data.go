package customerrors

type NoDataUpdatedError struct {
}

func (e NoDataUpdatedError) Error() string {
	return "Failed Updated Data"
}
