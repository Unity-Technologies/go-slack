package slack

// Response is a generic response from Slack which implements the SendResponse interface.
type Response struct {
	OK      bool   `json:"ok"`
	Error   string `json:"error,omitempty"`
	Warning string `json:"warning,omitempty"`
}

// Ok implements the SendResponse.Ok method.
func (r Response) Ok() bool {
	return r.OK
}

// Err implements the SendResponse.Err method.
func (r Response) Err() string {
	return r.Error
}

// Warn implements the SendResponse.Warn method.
func (r Response) Warn() string {
	return r.Warning
}
