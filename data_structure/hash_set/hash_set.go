package hash_set

import (
	"fmt"
	"hash/fnv"
)

type IHashSet[T any] interface {
	Add(T) int
	AddSlice([]T) int
	AddItems(first T, rest ...T) int
	Remove(T) bool
	Items() []T
	IsExist(T) bool
	IsEmpty() bool
	Size() int
	String() string
}

type HashSet[T any] struct {
	items map[uint64]T
	size  int
}

func New[T any](items ...T) *HashSet[T] {
	hashSet := &HashSet[T]{items: make(map[uint64]T)}
	hashSet.AddSlice(items)
	return hashSet
}

func (base *HashSet[T]) Add(item T) int {
	key := hash(item)
	base.items[key] = item
	base.size++
	return base.size
}

func (base *HashSet[T]) AddSlice(items []T) int {
	for _, item := range items {
		base.Add(item)
	}
	return base.size
}

func (base *HashSet[T]) AddItems(first T, rest ...T) int {
	base.Add(first)
	base.AddSlice(rest)
	return base.size
}

func (base *HashSet[T]) Remove(item T) bool {
	key := hash(item)
	if _, exist := base.items[key]; exist {
		delete(base.items, key)
		_, found := base.items[key]
		return !found
	}
	return false
}

func (base *HashSet[T]) Items() []T {
	items := make([]T, 0, base.size)
	for _, item := range base.items {
		items = append(items, item)
	}
	return items
}

func (base *HashSet[T]) IsEmpty() bool {
	return base.size == 0
}

func (base *HashSet[T]) Size() int {
	return base.size
}

func (base *HashSet[T]) String() string {
	return fmt.Sprintf("%v", base.Items())
}

func hash(value any) uint64 {
	h := fnv.New64a()
	h.Write([]byte(fmt.Sprintf("%v_%T", value, value)))
	return h.Sum64()
}
