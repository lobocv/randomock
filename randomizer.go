package randomock

// Randomizer interface describes the RandoMock and Random structs and exposes all random functions provided by
// the standard library's rand package
type Randomizer interface {
	ExpFloat64(key string) float64
	Float32(key string) float32
	Float64(key string) float64
	Int(key string) int
	Int31(key string) int32
	Int31n(key string, n int32) int32
	Int63(key string) int64
	Int63n(key string, n int64) int64
	Intn(key string, n int) int
	NormFloat64(key string) float64
	Uint32(key string) uint32
	Uint64(key string) uint64
}

var _ Randomizer = (*RandoMock)(nil)
var _ Randomizer = (*Random)(nil)
