package gc_friendly

import (
	"testing"
)

/**
1. go test -bench="<bench_name>" -trace="<out_file>.out"
2. go tool trace <out_file>.out
*/

const NumOfElems = 1000

type Content struct {
	Detail [10000]int
}

func withValue(arr [NumOfElems]Content) int {
	//	fmt.Println(&arr[2])
	return 0
}

func withReference(arr *[NumOfElems]Content) int {
	//b := *arr
	//	fmt.Println(&arr[2])
	return 0
}

func TestFn(t *testing.T) {
	var arr [NumOfElems]Content
	//fmt.Println(&arr[2])
	withValue(arr)
	withReference(&arr)
}

func BenchmarkPassingArrayWithValue(b *testing.B) {
	var arr [NumOfElems]Content

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withValue(arr)
	}
	b.StopTimer()
}

func BenchmarkPassingArrayWithRef(b *testing.B) {
	var arr [NumOfElems]Content

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withReference(&arr)
	}
	b.StopTimer()
}
