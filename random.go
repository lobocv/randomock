package randomock

import (
	"math/rand"
)

// Random is the structure used for real random calls.
// It is a simple struct that exposes rand package functions as methods. All methods take a key as an argument, however
// this key is only used during mocking. Regardless, keys should be given meaningful and descriptive names.
type Random struct {
}

func (r *Random) ExpFloat64(key string) float64 {
	return rand.ExpFloat64()
}

func (r *Random) Float32(key string) float32 {
	return rand.Float32()
}

func (r *Random) Float64(key string) float64 {
	return rand.Float64()
}

func (r *Random) Int(key string) int {
	return rand.Int()
}

func (r *Random) Int31(key string) int32 {
	return rand.Int31()
}

func (r *Random) Int31n(key string, n int32) int32 {
	return rand.Int31n(n)
}

func (r *Random) Int63(key string) int64 {
	return rand.Int63()
}

func (r *Random) Int63n(key string, n int64) int64 {
	return rand.Int63n(n)
}

func (r *Random) Intn(key string, n int) int {
	return rand.Intn(n)
}

func (r *Random) NormFloat64(key string) float64 {
	return rand.NormFloat64()
}

func (r *Random) Uint32(key string) uint32 {
	return rand.Uint32()
}

func (r *Random) Uint64(key string) uint64 {
	return rand.Uint64()
}
