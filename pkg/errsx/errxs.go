package errsx

type Errsx struct {
	Code     int
	Location string
	Message  string
	Err      error
}

func (c Errsx) Error() string {
	return c.Err.Error()
}

func NewCustomError(code int, location string, message string, err error) Errsx {
	return Errsx{
		Code:     code,
		Location: location,
		Message:  message,
		Err:      err,
	}
}
