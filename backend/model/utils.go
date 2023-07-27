package model

func toBoolMap[K comparable](list []K) map[K]bool {
	m := map[K]bool{}
	for _, k := range list {
		m[k] = true
	}
	return m
}
