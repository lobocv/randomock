package randomock

// Policy configures the policy that RandoMock objects adhere to when they are called more times than they have return
// values.
type Policy int

const (
	// ErrorOutPolicy is the default policy and causes RandoMock to panic when called more times
	// than there are return values for a specific key. This only applies if there are more than one return value.
	ErrorOutPolicy Policy = iota

	// WrapAroundPolicy causes the return values to repeat from the beginning after being exhausted.
	WrapAroundPolicy

	// RepeatLastPolicy causes the return values to repeat the last return value after being exhausted.
	RepeatLastPolicy
)

var defaultPolicy = ErrorOutPolicy

// OutOfBound is the structure returned when an ErrorOutPolicy panics.
type OutOfBound struct {
}

func (err OutOfBound) Error() string {
	return "more calls to randomock than return values"
}

// OutOfBoundsError error is the error value returned when the policy is broken
var OutOfBoundsError = OutOfBound{}

// SetDefaultPolicy sets the default policy of all RandoMock instances going forward.
// This does not modify already existing instances.
func SetDefaultPolicy(p Policy) {
	defaultPolicy = p
}

// results keeps track of return values for a specific key of RandoMock
type results struct {
	values []float64
	count  int
	policy Policy
}

func (r *results) Add(values ...float64) {
	r.values = append(r.values, values...)
}

func (r *results) Get() float64 {
	if len(r.values) > 1 {
		if r.count >= len(r.values) {
			switch r.policy {
			case WrapAroundPolicy:
				r.count = 0
			case RepeatLastPolicy:
				return r.values[len(r.values)-1]
			case ErrorOutPolicy:
				panic(OutOfBoundsError)
			}
		}
		v := r.values[r.count]
		r.count++
		return v
	}

	return r.values[0]
}

// RandoMock is a mock of Random which keeps track of the return values for each key. Use this in your tests.
type RandoMock struct {
	ret    map[string]*results
	policy Policy
}

// NewRandoMock creates a new RandoMock instance with the default policy.
func NewRandoMock() *RandoMock {
	r := &RandoMock{ret: make(map[string]*results), policy: defaultPolicy}
	return r
}

// Add adds return values to a specific key.
func (r *RandoMock) Add(key string, values ...float64) *RandoMock {
	res, ok := r.ret[key]
	if !ok {
		res = &results{policy: r.policy}
		r.ret[key] = res
	}
	res.Add(values...)

	return r
}

// SetPolicy sets the policy of the specific key.
func (r *RandoMock) SetPolicy(key string, p Policy) *RandoMock {
	r.ret[key].policy = p
	return r
}

// Policy returns the policy of a specific key
func (r *RandoMock) Policy(key string) Policy {
	return r.ret[key].policy
}

// rand functions:

// ExpFloat64 is a mocked version of rand.ExpFloat64
func (r *RandoMock) ExpFloat64(key string) float64 {
	return r.ret[key].Get()
}

// Float32 is a mocked version of rand.Float32
func (r *RandoMock) Float32(key string) float32 {
	return float32(r.ret[key].Get())
}

// Float64 is a mocked version of rand.Float64
func (r *RandoMock) Float64(key string) float64 {
	return r.ret[key].Get()
}

// Int is a mocked version of rand.Int
func (r *RandoMock) Int(key string) int {
	return int(r.ret[key].Get())
}

// Int31 is a mocked version of rand.Int31
func (r *RandoMock) Int31(key string) int32 {
	return int32(r.ret[key].Get())
}

// Int31n is a mocked version of rand.Int31n
func (r *RandoMock) Int31n(key string, n int32) int32 {
	return int32(r.ret[key].Get())
}

// Int63 is a mocked version of rand.Int63
func (r *RandoMock) Int63(key string) int64 {
	return int64(r.ret[key].Get())
}

// Int63n is a mocked version of rand.Int63n
func (r *RandoMock) Int63n(key string, n int64) int64 {
	return int64(r.ret[key].Get())
}

// Intn is a mocked version of rand.Intn
func (r *RandoMock) Intn(key string, n int) int {
	return int(r.ret[key].Get())
}

// NormFloat64 is a mocked version of rand.NormFloat64
func (r *RandoMock) NormFloat64(key string) float64 {
	return r.ret[key].Get()
}

// Uint32 is a mocked version of rand.Uint32
func (r *RandoMock) Uint32(key string) uint32 {
	return uint32(r.ret[key].Get())
}

// Uint64 is a mocked version of rand.Uint64
func (r *RandoMock) Uint64(key string) uint64 {
	return uint64(r.ret[key].Get())
}
