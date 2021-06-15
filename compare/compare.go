// Package compare facilitates comparisons of any two values according to a set of specifications.
package compare

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"runtime/debug"
	"strings"
	"time"
)

func Compare(a, b interface{}) error {
	if check(a, b) {
		return nil
	}
	return errors.New(report(resolveFormatter(a), a, b))
}

func check(a, b interface{}) bool {
	for _, spec := range []specification{numericEquality{}, timeEquality{}, deepEquality{}} {
		if !spec.IsSatisfiedBy(a, b) {
			continue
		}
		if spec.Compare(a, b) {
			return true
		}
		break
	}
	return false
}
func resolveFormatter(v interface{}) formatter {
	formatVerb := func(verb string) formatter {
		return func(v interface{}) string {
			return fmt.Sprintf(verb, v)
		}
	}
	if isNumeric(v) || isTime(v) {
		return formatVerb("%v")
	} else {
		return formatVerb("%#v")
	}
}
func report(format formatter, a, b interface{}) string {
	aType := fmt.Sprintf("(%v)", reflect.TypeOf(a))
	bType := fmt.Sprintf("(%v)", reflect.TypeOf(b))
	longestType := int(math.Max(float64(len(aType)), float64(len(bType))))
	aType += strings.Repeat(" ", longestType-len(aType))
	bType += strings.Repeat(" ", longestType-len(bType))
	aFormat := format(a)
	bFormat := format(b)
	typeDiff := diff(bType, aType)
	valueDiff := diff(bFormat, aFormat)

	builder := new(strings.Builder)
	_, _ = fmt.Fprintf(builder, "\n")
	_, _ = fmt.Fprintf(builder, "A: %s %s\n", aType, aFormat)
	_, _ = fmt.Fprintf(builder, "B: %s %s\n", bType, bFormat)
	_, _ = fmt.Fprintf(builder, "   %s %s\n", typeDiff, valueDiff)
	_, _ = fmt.Fprintf(builder, "Stack (filtered):\n%s\n", stack())

	return builder.String()
}
func diff(a, b string) string {
	result := new(strings.Builder)

	for x := 0; ; x++ {
		if x >= len(a) && x >= len(b) {
			break
		}
		if x >= len(a) || x >= len(b) || a[x] != b[x] {
			result.WriteString("^")
		} else {
			result.WriteString(" ")
		}
	}
	return result.String()
}
func stack() string {
	lines := strings.Split(string(debug.Stack()), "\n")
	var filtered []string
	for x := 1; x < len(lines)-1; x += 2 {
		if strings.Contains(lines[x+1], "_test.go:") {
			filtered = append(filtered, lines[x], lines[x+1])
		}
	}
	return "> " + strings.Join(filtered, "\n> ")
}

type formatter func(interface{}) string

type specification interface {
	IsSatisfiedBy(a, b interface{}) bool
	Compare(a, b interface{}) bool
}

// deepEquality compares any two values using reflect.DeepEqual.
// https://golang.org/pkg/reflect/#DeepEqual
type deepEquality struct{}

func (this deepEquality) IsSatisfiedBy(a, b interface{}) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}
func (this deepEquality) Compare(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

// numericEquality compares numeric values using the built-in equality
// operator (`==`). Values of differing numeric reflect.Kind are each
// converted to the type of the other and are compared with `==` in both
// directions. https://golang.org/pkg/reflect/#Kind
type numericEquality struct{}

func (this numericEquality) IsSatisfiedBy(a, b interface{}) bool {
	return isNumeric(a) && isNumeric(b)
}
func (this numericEquality) Compare(a, b interface{}) bool {
	if a == b {
		return true
	}
	aValue := reflect.ValueOf(a)
	bValue := reflect.ValueOf(b)
	aAsB := aValue.Convert(bValue.Type()).Interface()
	bAsA := bValue.Convert(aValue.Type()).Interface()
	return a == bAsA && b == aAsB
}
func isNumeric(v interface{}) bool {
	kind := reflect.TypeOf(v).Kind()
	return kind == reflect.Int ||
		kind == reflect.Int8 ||
		kind == reflect.Int16 ||
		kind == reflect.Int32 ||
		kind == reflect.Int64 ||
		kind == reflect.Uint ||
		kind == reflect.Uint8 ||
		kind == reflect.Uint16 ||
		kind == reflect.Uint32 ||
		kind == reflect.Uint64 ||
		kind == reflect.Float32 ||
		kind == reflect.Float64
}

// timeEquality compares values both of type time.Time using their Equal method.
// https://golang.org/pkg/time/#Time.Equal
type timeEquality struct{}

func (this timeEquality) IsSatisfiedBy(a, b interface{}) bool {
	return isTime(a) && isTime(b)
}
func (this timeEquality) Compare(a, b interface{}) bool {
	return a.(time.Time).Equal(b.(time.Time))
}

func isTime(v interface{}) bool {
	_, ok := v.(time.Time)
	return ok
}
