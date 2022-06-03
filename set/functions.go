package set

func Union[T comparable](sets ...Set[T]) (res Set[T]) {
	res = make(Set[T])
	for _, set := range sets {
		for v := range set {
			res.Add(v)
		}
	}
	return
}

func Difference[T comparable](set1 Set[T], set2 Set[T]) (res Set[T]) {
	res = set1.Copy()

	for el := range set2 {
		res.Remove(el)
	}
	return
}

func Intersection[T comparable](sets ...Set[T]) (res Set[T]) {
	res = make(Set[T])

	for el := range sets[0] {
		ok := true
		for _, set := range sets {
			if !set.Contains(el) {
				ok = false
				break
			}
		}
		if ok {
			res.Add(el)
		}
	}
	return
}

func SymmetricalDifference[T comparable](set1 Set[T], set2 Set[T]) Set[T] {
	return Difference(Union(set1, set2), Intersection(set1, set2))
}
