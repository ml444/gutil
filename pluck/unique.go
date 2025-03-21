package pluck

func ToUniqueString(list []string) []string {
	m := make(map[string]struct{})
	for _, v := range list {
		m[v] = struct{}{}
	}
	var newList []string
	for k := range m {
		newList = append(newList, k)
	}
	return newList
}

func ToUniqueNumbers[T TNumber](list []T) (newList []T) {
	m := make(map[T]struct{})
	for _, v := range list {
		m[v] = struct{}{}
	}
	for k := range m {
		newList = append(newList, k)
	}
	return
}
