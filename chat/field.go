package chat

var (
	// ShortFieldLen is the length of a field value which is to be deemed short.
	ShortFieldLen = 20
)

// Field will be displayed in a table inside the message attachment.
type Field struct {
	// Title is shown as a bold heading above the value text.
	// It cannot contain markup and will be escaped for you.
	Title string `json:"title"`

	// Value is the text value of the field.
	// It may contain standard message markup and must be escaped as normal.
	// May be multi-line.
	Value string `json:"value"`

	// Short is an optional flag indicating whether the value is short enough to be displayed side-by-side with other values.
	Short bool `json:"short"`
}

// NewField returns a fully initialised field with Short set to true if the length of value is less than ShortFieldLen.
func NewField(title, value string) *Field {
	return &Field{Title: title, Value: value, Short: len(value) < ShortFieldLen}
}
