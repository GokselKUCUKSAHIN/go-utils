package constants

import "golang.org/x/exp/constraints"

type KeyValueStore map[string]struct{}

func Keys[K constraints.Ordered, V any](source map[K]V) []K {
	resultKeys := make([]K, 0, len(source))
	for k := range source {
		resultKeys = append(resultKeys, k)
	}
	return resultKeys
}

func Values[K constraints.Ordered, V any](source map[K]V) []V {
	resultValues := make([]V, 0, len(source))
	for _, v := range source {
		resultValues = append(resultValues, v)
	}
	return resultValues
}

func CreateConstMap(first string, rest ...string) KeyValueStore {
	resultMap := map[string]struct{}{
		first: {},
	}
	for _, v := range rest {
		resultMap[v] = struct{}{}
	}
	return resultMap
}

func (base KeyValueStore) IsExist(key string) bool {
	_, exist := base[key]
	return exist
}
