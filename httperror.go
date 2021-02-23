package tableau

type HttpError struct {
	Code int   `json:"code"`
	Err  error `json:"error"`
}

func (h *HttpError) Error() string {
	if h.Err != nil {
		return h.Err.Error()
	}

	return ""
}
