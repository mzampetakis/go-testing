package assert

import (
	"strings"
	"testing"

	"github.com/ppapapetrou76/go-testing/internal/pkg/values"
)

// StringOpt is a configuration option to initialize an AssertableString
type StringOpt func(*AssertableString)

// AssertableString is the implementation of CommonAssertable for string types
type AssertableString struct {
	t      *testing.T
	actual values.StringValue
}

// IgnoringCase sets underlying value to lower case
func IgnoringCase() StringOpt {
	return func(c *AssertableString) {
		c.actual = c.actual.AddDecorator(strings.ToLower)
	}
}

// ThatString returns an AssertableString structure initialized with the test reference and the actual value to assert
func ThatString(t *testing.T, actual string, opts ...StringOpt) AssertableString {
	assertable := &AssertableString{
		t:      t,
		actual: values.NewStringValue(actual),
	}
	for _, opt := range opts {
		opt(assertable)
	}
	return *assertable
}

// IsEqualTo asserts if the expected string is equal to the assertable string value
// It errors the tests if the compared values (actual VS expected) are not equal
func (a AssertableString) IsEqualTo(expected interface{}) AssertableString {
	if !a.actual.IsEqualTo(expected) {
		a.t.Error(shouldBeEqual(a.actual, expected))
	}
	return a
}

// IsNotEqualTo asserts if the expected string is not equal to the assertable string value
// It errors the tests if the compared values (actual VS expected) are equal
func (a AssertableString) IsNotEqualTo(expected interface{}) AssertableString {
	if a.actual.IsEqualTo(expected) {
		a.t.Error(shouldNotBeEqual(a.actual, expected))
	}
	return a
}

// IsEmpty asserts if the expected string is empty
// It errors the tests if the string is not empty
func (a AssertableString) IsEmpty() AssertableString {
	if a.actual.IsNotEmpty() {
		a.t.Error(shouldBeEmpty(a.actual))
	}
	return a
}

// IsNotEmpty asserts if the expected string is not empty
// It errors the tests if the string is empty
func (a AssertableString) IsNotEmpty() AssertableString {
	if a.actual.IsEmpty() {
		a.t.Error(shouldNotBeEmpty(a.actual))
	}
	return a
}

// Contains asserts if the assertable string contains the given element(s)
// It errors the test if it does not contain it
func (a AssertableString) Contains(substring string) AssertableString {
	if a.actual.DoesNotContain(substring) {
		a.t.Error(shouldContain(a.actual, substring))
	}
	return a
}

// ContainsIgnoringCase asserts if the assertable string contains the given element(s) case insensitively
// It errors the test if it does not contain it
func (a AssertableString) ContainsIgnoringCase(substring string) AssertableString {
	if !a.actual.ContainsIgnoringCase(substring) {
		a.t.Error(shouldContainIgnoringCase(a.actual, substring))
	}
	return a
}

// ContainsOnly asserts if the assertable string only contains the given substring
// It errors the test if it does not contain it
func (a AssertableString) ContainsOnly(substring string) AssertableString {
	if !a.actual.ContainsOnly(substring) {
		a.t.Error(shouldContainOnly(a.actual, substring))
	}
	return a
}

// ContainsOnlyOnce asserts if the assertable string contains the given substring only once
// It errors the test if it does not contain it or contains more than once
func (a AssertableString) ContainsOnlyOnce(substring string) AssertableString {
	if !a.actual.ContainsOnlyOnce(substring) {
		a.t.Error(shouldContainOnlyOnce(a.actual, substring))
	}
	return a
}

// DoesNotContain asserts if the assertable string does not contain the given substring
// It errors the test if it contains it
func (a AssertableString) DoesNotContain(substring string) AssertableString {
	if a.actual.Contains(substring) {
		a.t.Error(shouldNotContain(a.actual, substring))
	}
	return a
}

// StartsWith asserts if the assertable string starts with the given substring
// It errors the test if it doesn't start with the given substring
func (a AssertableString) StartsWith(substring string) AssertableString {
	if !a.actual.StartsWith(substring) {
		a.t.Error(shouldStartWith(a.actual, substring))
	}
	return a
}

// EndsWith asserts if the assertable string ends with the given substring
// It errors the test if it doesn't end with the given substring
func (a AssertableString) EndsWith(substring string) AssertableString {
	if !a.actual.EndsWith(substring) {
		a.t.Error(shouldEndWith(a.actual, substring))
	}
	return a
}

// HasSameSizeAs asserts if the assertable string has the same size with the given string
// It errors the test if they don't have the same size
func (a AssertableString) HasSameSizeAs(substring string) AssertableString {
	if !(a.actual.Size() == len(substring)) {
		a.t.Error(shouldHaveSameSizeAs(a.actual, substring))
	}
	return a
}

// ContainsOnlyDigits asserts if the expected string contains only digits
// It errors the tests if the string has other characters than digits
func (a AssertableString) ContainsOnlyDigits() AssertableString {
	if !(a.actual.HasDigitsOnly()) {
		a.t.Error(shouldContainOnlyDigits(a.actual))
	}
	return a
}
