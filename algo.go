package z

import (
	"iter"
)

func AllOf[T any](vs iter.Seq[T], p func(v T) bool) bool {
	for v := range vs {
		if !p(v) {
			return false
		}
	}

	return true
}

func AnyOf[T any](vs iter.Seq[T], p func(v T) bool) bool {
	for v := range vs {
		if p(v) {
			return true
		}
	}

	return false
}

func NoneOf[T any](vs iter.Seq[T], p func(v T) bool) bool {
	for v := range vs {
		if p(v) {
			return false
		}
	}

	return true
}
