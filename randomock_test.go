package randomock

import (
	"math/rand"
	"testing"
)

var methods = []string{"ExpFloat64", "Float32", "Float64", "Int", "Int31", "Int31n", "Int63", "Int63n", "Intn", "NormFloat64", "Uint32", "Uint64"}

type callpair struct {
	key string
	val float64
}

var testCases = []struct {
	calls []callpair
}{
	{
		[]callpair{
			{"a", 1.5},
			{"b", 5.5},
			{"c", -35.5},
		},
	},
}

func setSeed() {
	rand.Seed(1)
}

func setCalls(calls []callpair) *RandoMock {
	r := NewRandoMock()
	for _, cp := range calls {
		r.Add(cp.key, cp.val)
	}
	return r
}

func TestRandoMock_Add(t *testing.T) {
	t.Parallel()
	for _, tc := range testCases {
		r := setCalls(tc.calls)
		for _, call := range tc.calls {
			results, ok := r.ret[call.key]
			if !ok {
				t.Fatalf("cannot find key %s", call.key)
			}
			got := results.Get()
			expected := call.val
			if got != expected {
				t.Fatalf("inserted call value of %f does not equal expected %f", got, expected)
			}
		}
	}
}

func callMethod(r Randomizer, method string, key string, args ...interface{}) interface{} {
	var result interface{}
	switch method {
	case "ExpFloat64":
		result = r.ExpFloat64(key)
	case "Float32":
		result = r.Float32(key)
	case "Float64":
		result = r.Float64(key)
	case "Int":
		result = r.Int(key)
	case "Int31":
		result = r.Int31(key)
	case "Int31n":
		result = r.Int31n(key, int32(args[0].(int)))
	case "Int63":
		result = r.Int63(key)
	case "Int63n":
		result = r.Int63n(key, int64(args[0].(int)))
	case "NormFloat64":
		result = r.NormFloat64(key)
	case "Intn":
		result = r.Intn(key, args[0].(int))
	case "Uint32":
		result = r.Uint32(key)
	case "Uint64":
		result = r.Uint64(key)
	}
	return result
}

func callRealMethod(method string, args ...interface{}) interface{} {
	var result interface{}
	switch method {
	case "ExpFloat64":
		result = rand.ExpFloat64()
	case "Float32":
		result = rand.Float32()
	case "Float64":
		result = rand.Float64()
	case "Int":
		result = rand.Int()
	case "Int31":
		result = rand.Int31()
	case "Int31n":
		result = rand.Int31n(int32(args[0].(int)))
	case "Int63":
		result = rand.Int63()
	case "Int63n":
		result = rand.Int63n(int64(args[0].(int)))
	case "NormFloat64":
		result = rand.NormFloat64()
	case "Intn":
		result = rand.Intn(args[0].(int))
	case "Uint32":
		result = rand.Uint32()
	case "Uint64":
		result = rand.Uint64()
	}
	return result
}

func convertResult(method string, r float64) interface{} {
	var result interface{}
	switch method {
	case "ExpFloat64", "Float64", "NormFloat64":
		result = float64(r)
	case "Float32":
		result = float32(r)
	case "Int", "Intn":
		result = int(r)
	case "Int31", "Int31n":
		result = int32(r)
	case "Int63", "Int63n":
		result = int64(r)
	case "Uint32":
		result = uint32(r)
	case "Uint64":
		result = uint64(r)
	}
	return result
}

func TestRandoMockReturns(t *testing.T) {
	t.Parallel()
	for _, tc := range testCases {
		for _, meth := range methods {
			r := setCalls(tc.calls)
			for _, call := range tc.calls {

				t.Run(meth,
					func(t *testing.T) {
						var expected = convertResult(meth, call.val)
						v := callMethod(r, meth, call.key, 123)
						if v != expected {
							t.Fatalf("RandoMock call value of %v does not equal expected %v", v, expected)
						}
					})
			}
		}
	}
}

func TestRandomReturns(t *testing.T) {
	t.Parallel()
	for _, tc := range testCases {
		r := &Random{}
		for _, call := range tc.calls {

			for _, meth := range methods {

				t.Run(meth,
					func(t *testing.T) {

						// Note we need to reset the seed after each call to a randomizing function
						// in order to get the same result.
						setSeed()
						var expected = callRealMethod(meth, 123)
						setSeed()
						v := callMethod(r, meth, call.key, 123)
						if v != expected {
							t.Fatalf("Random call value of %d does not equal expected %d", v, expected)
						}
					})
			}
		}
	}
}

func TestHelpers(t *testing.T) {
	t.Parallel()
	t.Run("RoundTo", func(t *testing.T) {
		testCases := []struct {
			v        float64
			prec     int
			expected float64
		}{
			{0.123, 2, 0.12},
			{-245.128, 2, -245.13},
			{5.1, 3, 5.1},
			{9.99, 1, 10.0},
			{0.0000123, 6, 0.000012},
			{0.001, 2, 0},
			{0.05, 1, 0.1},
			{0.6, 0, 1},
			{0.008, 2, 0.01},
			{0.02, 1, 0},
		}
		for _, tc := range testCases {
			r := RoundTo(tc.v, tc.prec)
			if r != tc.expected {
				t.Fatalf("expected %f, got %f", tc.expected, r)
			}
		}

	})

	t.Run("RandBetweenFloat64", func(t *testing.T) {
		attempts := 10
		testCases := []struct {
			a, b float64
		}{
			{0, 5},
			{-100, 100},
			{7, 7},
			{-10, -10},
		}
		for _, tc := range testCases {
			for attempt := 0; attempt < attempts; attempt++ {
				r := RandBetweenFloat64(tc.a, tc.b)
				if r < tc.a || r > tc.b {
					t.Fatalf("expected value to be between %f and %f", tc.a, tc.b)
				}
			}
		}
	})
}

func TestPolicies(t *testing.T) {
	allExpected := []float64{0.5, 1, 2.5, 3.5, 4}

	testCases := []struct {
		name         string
		policy       Policy
		expectedFunc func(callNum int) (float64, error)
	}{
		{
			name:   "WrapAroundPolicy",
			policy: WrapAroundPolicy,
			expectedFunc: func(count int) (float64, error) {
				return allExpected[count%len(allExpected)], nil
			},
		},
		{
			name:   "ErrorOutPolicy",
			policy: ErrorOutPolicy,
			expectedFunc: func(count int) (float64, error) {
				if count >= len(allExpected) {
					return 0, OutOfBoundsError
				}
				return allExpected[count], nil
			},
		},
		{
			name:   "RepeatLastPolicy",
			policy: RepeatLastPolicy,
			expectedFunc: func(count int) (float64, error) {
				if count >= len(allExpected) {
					count = len(allExpected) - 1
				}
				return allExpected[count], nil
			},
		},
	}

	tries := 50
	for _, tc := range testCases {
		r := NewRandoMock().Add("k", allExpected...)
		r.SetPolicy("k", tc.policy)

		t.Run(tc.name, func(t *testing.T) {
			for ii := 0; ii < tries; ii++ {
				expected, errExpected := tc.expectedFunc(ii)

				defer func(t *testing.T) {
					err := recover()
					if errExpected != err {
						t.Fatal(errExpected)
					}
				}(t)

				got := r.Float64("k")
				if got != expected {
					t.Fatalf("expected %f, got %f", expected, got)
				}
			}
		})

	}

	t.Run("DefaultPolicy", func(t *testing.T) {
		before := defaultPolicy
		policy := RepeatLastPolicy

		SetDefaultPolicy(policy)
		now := defaultPolicy
		if now != RepeatLastPolicy || now == before {
			t.Fatalf("failed to change default policy")
		}

		r := NewRandoMock().Add("k", allExpected...)
		if r.Policy("k") != policy {
			t.Fatalf("failed to change default policy of results struct")
		}

	})
}

func TestErrors(t *testing.T) {
	msg := OutOfBoundsError.Error()
	if msg == "" {
		t.Fatalf("no error message")
	}
}
