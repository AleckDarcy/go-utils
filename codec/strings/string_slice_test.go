package strings

import (
	"reflect"
	"testing"
)

func DeepCopyStringSliceOld(slice []string) []string {
	newSlice := make([]string, len(slice))

	for i, str := range slice {
		newSlice[i] = DeepCopyString(str)
	}

	return newSlice
}

func ShallowCopyStringSlice(slice []string) []string {
	newSlice := make([]string, len(slice))

	for i, str := range slice {
		newSlice[i] = str
	}

	return newSlice
}

func BenchmarkDeepCopyStringSliceOld(b *testing.B) {
	var sample = []string{"string0", "", "string1"}
	var copy []string

	for i := 0; i < b.N; i++ {
		copy = DeepCopyStringSliceOld(sample)
	}

	if !reflect.DeepEqual(copy, sample) {
		b.Error("fail", copy)
	}
}

func BenchmarkShallowCopyStringSlice(b *testing.B) {
	var sample = []string{"string0", "", "string1"}
	var copy []string

	for i := 0; i < b.N; i++ {
		copy = ShallowCopyStringSlice(sample)
	}

	if !reflect.DeepEqual(copy, sample) {
		b.Error("fail", copy)
	}
}

func BenchmarkDeepCopyStringSlice(b *testing.B) {
	var sample = []string{"string0", "", "string1"}
	var copy []string

	for i := 0; i < b.N; i++ {
		copy = DeepCopyStringSlice(sample)
	}

	if !reflect.DeepEqual(copy, sample) {
		b.Error("fail", copy)
	}
}
