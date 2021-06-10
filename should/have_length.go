package should

import "reflect"

// HaveLength uses reflection to verify that len(actual) == 0.
func HaveLength(actual interface{}, expected ...interface{}) error {
	err := validateExpected(1, expected)
	if err != nil {
		return err
	}

	err = validateKind(actual, kindsWithLength...)
	if err != nil {
		return err
	}

	err = validateKind(expected[0], numericKinds...)
	if err != nil {
		return err
	}

	expectedLength := reflect.ValueOf(expected[0]).Int()
	actualLength := int64(reflect.ValueOf(actual).Len())
	if actualLength == expectedLength {
		return nil
	}

	TYPE := reflect.TypeOf(actual).String()
	return failure("got length of %d, want %d", TYPE, actualLength, expectedLength)
}

var numericKinds = []reflect.Kind{
	reflect.Int,
	reflect.Int8,
	reflect.Int16,
	reflect.Int32,
	reflect.Int64,
}
