package pluck

func ToUnique(list []string) []string {
	var m = make(map[string]struct{})
	for _, v := range list {
		m[v] = struct{}{}
	}
	var newList []string
	for k := range m {
		newList = append(newList, k)
	}
	return newList
}
