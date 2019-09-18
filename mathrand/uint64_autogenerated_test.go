package mathrand_test

import (
	"testing"

	"github.com/xaionaro-go/rand/mathrand"
)

func TestUint64AddRotateMultiply(t *testing.T) {
	//testUint64(t, mathrand.GlobalPRNG.Uint64AddRotateMultiply)
}
func BenchmarkUint64AddRotateMultiply(b *testing.B) {
	prng := mathrand.New()
	b.SetBytes(8)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prng.Uint64AddRotateMultiply()
	}
}

func TestUint64AddNRotateMultiply(t *testing.T) {
	//testUint64(t, mathrand.GlobalPRNG.Uint64AddNRotateMultiply)
}
func BenchmarkUint64AddNRotateMultiply(b *testing.B) {
	prng := mathrand.New()
	b.SetBytes(8)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prng.Uint64AddNRotateMultiply()
	}
}

func TestUint64MultiplyAdd(t *testing.T) {
	//testUint64(t, mathrand.GlobalPRNG.Uint64MultiplyAdd)
}
func BenchmarkUint64MultiplyAdd(b *testing.B) {
	prng := mathrand.New()
	b.SetBytes(8)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prng.Uint64MultiplyAdd()
	}
}

func TestUint64AddRotate(t *testing.T) {
	//testUint64(t, mathrand.GlobalPRNG.Uint64AddRotate)
}
func BenchmarkUint64AddRotate(b *testing.B) {
	prng := mathrand.New()
	b.SetBytes(8)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prng.Uint64AddRotate()
	}
}

func TestUint64Xorshift(t *testing.T) {
	//testUint64(t, mathrand.GlobalPRNG.Uint64Xorshift)
}
func BenchmarkUint64Xorshift(b *testing.B) {
	prng := mathrand.New()
	b.SetBytes(8)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prng.Uint64Xorshift()
	}
}

func TestUint64Xoshiro256(t *testing.T) {
	//testUint64(t, mathrand.GlobalPRNG.Uint64Xoshiro256)
}
func BenchmarkUint64Xoshiro256(b *testing.B) {
	prng := mathrand.New()
	b.SetBytes(8)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prng.Uint64Xoshiro256()
	}
}

func TestUint64LFG(t *testing.T) {
	//testUint64(t, mathrand.GlobalPRNG.Uint64LFG)
}
func BenchmarkUint64LFG(b *testing.B) {
	prng := mathrand.New()
	b.SetBytes(8)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prng.Uint64LFG()
	}
}

func TestUint64MSWS(t *testing.T) {
	//testUint64(t, mathrand.GlobalPRNG.Uint64MSWS)
}
func BenchmarkUint64MSWS(b *testing.B) {
	prng := mathrand.New()
	b.SetBytes(8)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prng.Uint64MSWS()
	}
}
