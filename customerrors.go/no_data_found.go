package customerrors

type NoDataFoundError struct {
}

func (e NoDataFoundError) Error() string {
	return "No data found"
}
