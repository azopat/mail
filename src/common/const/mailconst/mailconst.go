package mailconst

type (
	QueueMailDoc struct {
		FromAddress string `valid:"-"`
		FromName    string `valid:"-"`
		To          string `valid:"-,required"`
		Subject     string `valid:"-,required"`
		Template    string `valid:"-,required"`
		InvoiceLink string `valid:"-,required"`
		Params      interface{}
	}
)

const (
	NSQ_TOPIC = "mails"
)
