package collection

func AppendUnique[T comparable](coll []T, obj T) []T {
	for _, c := range coll {
		if c == obj {
			return coll
		}
	}
	return append(coll, obj)
}
