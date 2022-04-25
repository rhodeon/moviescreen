// Package testhelpers provides functions for assisting with tests.
package testhelpers

import (
	"reflect"
	"testing"
)

func AssertEqual[T comparable](t *testing.T, got T, want T) {
	t.Helper()

	if got != want {
		t.Errorf("\nGot:\t%v\nWant:\t%v", got, want)
	}

}

func AssertError(t *testing.T, got error, want error) {
	t.Helper()

	if got != want {
		t.Errorf("\nGot Error:\t%+v\nWant Error:\t%+v", got, want)
	}
}

func AssertFatalError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatal(err)
	}
}

func AssertStruct(t *testing.T, got any, want any) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\nGot:\t%+v\nWant:\t%+v", got, want)
	}
}
