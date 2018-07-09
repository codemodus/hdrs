package hedrs

import "sync"

var (
	// DefaultOrigins ...
	DefaultOrigins = []string{
		"https://127.0.0.1",
		"http://127.0.0.1",
		"https://localhost",
		"http://localhost",
	}
)

// Allowed ...
type Allowed struct {
	hosts map[string]struct{}
}

// NewAllowed ...
func NewAllowed(valid ...string) *Allowed {
	m := make(map[string]struct{})
	applyAndAddToMap(m, prunePort, valid...)

	return &Allowed{
		hosts: m,
	}
}

// IsAllowed ...
func (o *Allowed) IsAllowed(origin string) bool {
	origin = prunePort(origin)

	if origin == "" {
		return false
	}

	return hasKeyOrWildcard(o.hosts, origin)
}

// AllowedRegistry ...
type AllowedRegistry struct {
	sync.Mutex
	hosts map[string]struct{}
}

// NewAllowedRegistry ...
func NewAllowedRegistry(valid ...string) *AllowedRegistry {
	m := make(map[string]struct{})
	applyAndAddToMap(m, prunePort, valid...)

	return &AllowedRegistry{
		hosts: m,
	}
}

// Add ...
func (o *AllowedRegistry) Add(valid ...string) {
	if valid == nil {
		return
	}

	o.Lock()
	defer o.Unlock()

	applyAndAddToMap(o.hosts, prunePort, valid...)
}

// Rmv ...
func (o *AllowedRegistry) Rmv(invalid ...string) {
	if invalid == nil {
		return
	}

	o.Lock()
	defer o.Unlock()

	applyAndDeleteFromMap(o.hosts, prunePort, invalid...)
}

// IsAllowed ...
func (o *AllowedRegistry) IsAllowed(origin string) bool {
	origin = prunePort(origin)

	if origin == "" {
		return false
	}

	o.Lock()
	defer o.Unlock()

	return hasKeyOrWildcard(o.hosts, origin)
}
