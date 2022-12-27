package dblayer

import "testing"

func BenchmarkHashPassword(b *testing.B) {
	text := "A String to be hashed"
	b.ResetTimer() // time 초기화해서 좀더 정확한 측정 가능
	for i := 0; i < b.N; i++ {
		hashPassword(&text)
	}
}
