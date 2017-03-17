package chat

// Attachment is a slack chat message attachment.
// See: https://api.slack.com/docs/message-attachments
type Attachment struct {
	// Fallback is the plain-text summary of the attachment.
	Fallback string `json:"fallback"`

	// Color is a color indicating the classification of the message.
	Color string `json:"color"`

	// PreText is optional text that appears above the message attachment block.
	PreText string `json:"pretext,omitempty"`

	// AuthorName is the small text used to display the author's name.
	AuthorName string `json:"author_name,omitempty"`

	// AuthorLink is a valid URL that will hyperlink the author_name text mentioned above.
	AuthorLink string `json:"author_link,omitempty"`

	// AuthorIcon is a valid URL that displays a small 16x16px image to the left of the author_name text.
	AuthorIcon string `json:"author_icon,omitempty"`

	// Title is displayed as larger, bold text near the top of a message attachment.
	Title string `json:"title,omitempty"`

	// TitleLink is the optional url of the hyperlink to be used for the title.
	TitleLink string `json:"title_link,omitempty"`

	// Text is the main text of the attachment.
	Text string `json:"text,omitempty"`

	// Fields contains optional fields to be displayed in the in a table inside the attachment.
	Fields []*Field `json:"fields,omitempty"`

	// MarkdownIn enables Markdown support. Valid values are ["pretext", "text", "fields"].
	// Setting "fields" will enable markup formatting for the value of each field.
	MarkdownIn []string `json:"mrkdwn_in,omitempty"`

	// ImageURL is the URL to an image file that will be displayed inside the attachment.
	ImageURL string `json:"image_url"`

	// ThumbURL is the URL to an image file that will be displayed as a thumbnail on the right side of a attachment.
	ThumbURL string `json:"ThumbURL"`

	// Footer is optional text to help contextualize and identify an attachment (300 chars max).
	Footer string `json:"footer"`

	// FooterIcon is the URL to a small icon beside your footer text.
	FooterIcon string `json:"footer_icon"`

	// TimeStamp if set is the epoch time that will display as part of the attachment's footer.
	TimeStamp int `json:"ts,omitempty"`
}

// NewField creates a new field, adds it to the attachment and then returns it.
func (a *Attachment) NewField(title, value string) *Field {
	f := NewField(title, value)
	a.AddField(f)

	return f
}

// AddField adds f to the attachments fields.
func (a *Attachment) AddField(f *Field) {
	a.Fields = append(a.Fields, f)
}
