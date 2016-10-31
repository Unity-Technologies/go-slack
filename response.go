package slack

// Response is a generic response from slack which implements SendResponse.
type Response struct {
	OK      bool   `json:"ok"`
	Error   string `json:"error,omitempty"`
	Warning string `json:"warning,omitempty"`
}

// Ok implements SendResponse.Ok.
func (r Response) Ok() bool {
	return r.OK
}

// Err implements SendResponse.Err.
func (r Response) Err() string {
	return r.Error
}

// Warn implements SendResponse.Warn.
func (r Response) Warn() string {
	return r.Warning
}
