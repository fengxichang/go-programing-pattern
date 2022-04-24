package slice_append

import "testing"

// 切片初始化长度不够，性能不好，每次append都会重新分配内存
func BenchmarkSliceAppend(t *testing.B) {
	a := make([]int, 0)
	for i := 0; i < t.N; i++ {
		a = append(a, i)
	}
}

// 性能好，每次复制都是对切片指针操作
func BenchmarkSliceSet(t *testing.B) {
	a := make([]int, t.N)
	for i := 0; i < t.N; i++ {
		a[i] = i
	}
}
