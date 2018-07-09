package hedrs

var (
	// DefaultHeaders ...
	DefaultHeaders = []string{
		Accept,
		AcceptEncoding,
		AcceptVersion,
		Authorization,
		ContentLength,
		ContentMD5,
		ContentType,
		Date,
		Origin,
		APIVersion,
		CSRFToken,
		RequestedWith,
	}
)

// Values ...
type Values struct {
	s string
}

// NewValues ...
func NewValues(vs ...string) *Values {
	return &Values{
		s: valuesToString(vs),
	}
}

// String ...
func (vs *Values) String() string {
	return vs.s
}
