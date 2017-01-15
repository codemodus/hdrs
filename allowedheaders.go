package hedrs

import "sync"

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

// AllowedHeaders ...
type AllowedHeaders struct {
	sync.Mutex
	hdrs map[string]struct{}
	str  string
}

// NewAllowedHeaders ...
func NewAllowedHeaders(valid ...string) *AllowedHeaders {
	h := &AllowedHeaders{
		hdrs: map[string]struct{}{},
	}

	h.Add(valid...)

	return h
}

// Add ...
func (h *AllowedHeaders) Add(valid ...string) {
	if valid == nil {
		return
	}

	h.Lock()
	defer h.Unlock()

	for _, v := range valid {
		if v == "" {
			continue
		}

		h.hdrs[v] = struct{}{}
	}

	h.str = keysToString(h.hdrs)
}

// Rmv ...
func (h *AllowedHeaders) Rmv(invalid ...string) {
	if invalid == nil {
		return
	}

	h.Lock()
	defer h.Unlock()

	for _, v := range invalid {
		if v == "" {
			continue
		}

		delete(h.hdrs, v)
	}

	h.str = keysToString(h.hdrs)
}

// String ...
func (h *AllowedHeaders) String() string {
	h.Lock()
	defer h.Unlock()

	return h.str
}
