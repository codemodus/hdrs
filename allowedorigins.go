package hedrs

import "sync"

// Allower ...
type Allower interface {
	IsAllowed(string) bool
}

var (
	// DefaultOrigins ...
	DefaultOrigins = []string{
		"https://127.0.0.1",
		"http://127.0.0.1",
		"https://localhost",
		"http://localhost",
	}
)

// AllowedOrigins ...
type AllowedOrigins struct {
	sync.Mutex
	hosts map[string]struct{}
}

// NewAllowedOrigins ...
func NewAllowedOrigins(valid ...string) *AllowedOrigins {
	o := &AllowedOrigins{
		hosts: map[string]struct{}{},
	}

	o.Add(valid...)

	return o
}

// Add ...
func (o *AllowedOrigins) Add(valid ...string) {
	if valid == nil {
		return
	}

	o.Lock()
	defer o.Unlock()

	for _, v := range valid {
		host := prunePort(v)

		if host == "" {
			continue
		}

		o.hosts[host] = struct{}{}
	}
}

// Rmv ...
func (o *AllowedOrigins) Rmv(invalid ...string) {
	if invalid == nil {
		return
	}

	o.Lock()
	defer o.Unlock()

	for _, v := range invalid {
		host := prunePort(v)

		if host == "" {
			continue
		}

		delete(o.hosts, host)
	}
}

// IsAllowed ...
func (o *AllowedOrigins) IsAllowed(origin string) bool {
	origin = prunePort(origin)

	if origin == "" {
		return false
	}

	o.Lock()
	defer o.Unlock()

	_, ok := o.hosts[origin]
	if ok {
		return true
	}

	_, ok = o.hosts["*"]

	return ok
}
