package hedrs

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

func keysToString(m map[string]struct{}) string {
	s := ""

	for k := range m {
		if s == "" {
			s += k
			continue
		}

		s += ", " + k
	}

	return s
}
