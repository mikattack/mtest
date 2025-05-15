package assert

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// True asserts that an expression is true.
func True(t testing.TB, ok bool) {
	if ok {
		return
	}
	t.Helper()
	t.Fatal("Expected expression to be true")
}

// False asserts that an expression is false.
func False(t testing.TB, ok bool) {
	if !ok {
		return
	}
	t.Helper()
	t.Fatal("Expected expression to be false")
}

// Equal asserts that "expected" and "actual" are equal.
func Equal[T any](t testing.TB, expected, actual T) {
	if objectsAreEqual(expected, actual) {
		return
	}
	t.Helper()
	t.Fatalf("%s\n%s", "Expected values to be equal:", diff(expected, actual))
}

// Error asserts that an error is not nil.
func Error(t testing.TB, err error) {
	if err != nil {
		return
	}
	t.Helper()
	t.Fatal("Expected an error")
}

// NoError asserts that an error is nil.
func NoError(t testing.TB, err error) {
	if err == nil {
		return
	}
	t.Helper()
	t.Fatalf("%s\n%+v", "Unexpected error:", err)
}

// Panics asserts that the given function panics.
func Panics(t testing.TB, fn func()) {
	t.Helper()
	defer func() {
		if recover() == nil {
			t.Fatal("Expected function to panic")
		}
	}()
	fn()
}

func diff(expected, actual any) string {
	lines := []string{
		"expected:",
		fmt.Sprintf("%v", expected),
		"actual:",
		fmt.Sprintf("%v", actual),
	}
	return strings.Join(lines, "\n")
}

func objectsAreEqual(expected, actual any) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}
	if exp, eok := expected.([]byte); eok {
		if act, aok := actual.([]byte); aok {
			return bytes.Equal(exp, act)
		}
	}
	if exp, eok := expected.(string); eok {
		if act, aok := actual.(string); aok {
			return exp == act
		}
	}

	return reflect.DeepEqual(expected, actual)
}
