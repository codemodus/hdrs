package hedrs

import "strings"

func valuesToString(vs []string) string {
	return strings.Join(vs, ", ")
}

func applyAndAddToMap(m map[string]struct{}, fn func(string) string, keys ...string) {
	if fn == nil {
		fn = passThrough
	}

	for _, v := range keys {
		v = fn(v)

		if v == "" {
			continue
		}

		m[v] = struct{}{}
	}
}

func applyAndDeleteFromMap(m map[string]struct{}, fn func(string) string, keys ...string) {
	if fn == nil {
		fn = passThrough
	}

	for _, v := range keys {
		v = fn(v)

		if v == "" {
			continue
		}

		delete(m, v)
	}
}

func hasKeyOrWildcard(m map[string]struct{}, key string) bool {
	_, ok := m[key]
	if ok {
		return true
	}

	_, ok = m["*"]

	return ok
}

func passThrough(s string) string {
	return s
}

func prunePort(host string) string {
	for k := range host {
		i := len(host) - k - 1
		v := host[i]
		if v == ':' && len(host) >= i+2 && host[i+1] != '/' {
			return host[:i]
		}
	}

	return host
}
