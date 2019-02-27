[![Go Report Card](https://goreportcard.com/badge/github.com/lobocv/randomock)](https://goreportcard.com/report/github.com/lobocv/randomock)

# Randomock

Randomock is a mocking library for the standard library rand package.

## Why

I was writing some simulation software that heavily utilized the rand package functions and 
found that it was difficult to write tests for most of my code. Since the standard library rand package
only provides functions and no interfaces, mocking the results of calls to the rand package was more of 
a pain than it needed to be.


## What Randomock Provides

Randomock provides the `Randomizer` interface, which describes two of the main structs in the package:
 
 `Random` - A thin wrapper around stdlib `rand` package functions. This is meant to be used in your non-test code.
 
 `RandoMock` - Mocked version of `Random` which allows you to specify the return values for certain function calls. 
 

## Usage

One of the main differences to using the `Randomizer` interface is that all calls must be given a key.
In the case of `RandomMock`, this key is used to keep track of which values to return. For `Random`, the key
does nothing.


### Writing test code

In your test code, create an instance of RandoMock:

```go
r := NewRandoMock()
```

Then add the mocked return values for a certain key:

```go
r.Add("dice", 4)
````

or, add many return values for a key at the same time:

```go
r.Add("dice", 4, 3, 1)
```

This will cause any call to a rand method to turn `4` the first time, `3` the second time and `1` the third time.

**For convenience, keys with only one return value always return that value no matter how many times they are called.**
 

For keys which have more than one return value, the values will be returned in the order they were added. 
When called more times than there are return values, the `RandoMock`'s policy defines what occurs.
 
The default policy is set to `ErrorOutPolicy`, which means the code will panic when called more times than there are 
return values. This is the safest solution as it will immediately indicate that you have not set up your 
tests with enough mock return values.

You can change the default policy with `SetDefaultPolicy(p Policy)` or the policy of individual RandoMock instances 
with the `SetPolicy(key string, p Policy)` method.


### Available policies

WrapAroundPolicy - When all return values are exhausted, start again from the first value.

RepeatLastPolicy - When all return values are exhausted, repeat the last return value.

ErrorOutPolicy   - When all return values are exhausted, panic on the next call.


### Writing non-test code

In your non-test code, simply create a `Random` instance and hold it inside of a `Randomizer` interface.

```go
    var r randomock.Randomizer = &randomock.Random{}

```

You can now call all your `rand` functions through this object.

```go 
    r.Int("roll")
    r.Float64("age")
    r.ExpFloat64("norm")
```

### Writing test code

Use a randomock.RandoMock instance in your Randomizer interface instead. Before calling your test, load up 
the mock return values and set the policy (if required).

```go
    var r randomock.Randomizer = randomock.NewRandoMock().Add("roll", 2.0, 3.0, 5.0, 2.0)
```

## Example

See the [example](example) directory for examples of how to use randomock in your tests. 


## Issues and Suggestions

Feel free to open a ticket if you find any issues or would like to request a feature. PR's are also welcome.
