package maps

func InvertMap[K, V comparable](m map[K]V) map[V]K {
	var res = make(map[V]K)
	for k, v := range m {
		res[v] = k
	}
	return res
}

func Keys[K comparable, V any](m map[K]V) []K {
	var res []K
	for key := range m {
		res = append(res, key)
	}
	return res
}

func Values[K comparable, V any](m map[K]V) []V {
	var res []V
	for _, value := range m {
		res = append(res, value)
	}
	return res
}
