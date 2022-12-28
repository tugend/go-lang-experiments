package lib

import "example/web-service-gin/persistence/types"

func Find[T comparable](elements []T, predicate func(T) bool) (T, error) {
	for _, element := range elements {
		if predicate(element) {
			return element, nil
		}
	}

	var result T
	return result, &types.NotFoundError{}
}
