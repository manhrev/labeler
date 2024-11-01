// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: auth/rpc/rollback_labeled_image.proto

package rpc

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"

	model "github.com/manhrev/labeler/pkg/api/go/auth/model"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort

	_ = model.Category(0)
)

// Validate checks the field values on RollbackLabeledImageRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RollbackLabeledImageRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RollbackLabeledImageRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RollbackLabeledImageRequestMultiError, or nil if none found.
func (m *RollbackLabeledImageRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RollbackLabeledImageRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := model.Category_name[int32(m.GetCategory())]; !ok {
		err := RollbackLabeledImageRequestValidationError{
			field:  "Category",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetId()) < 1 {
		err := RollbackLabeledImageRequestValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return RollbackLabeledImageRequestMultiError(errors)
	}

	return nil
}

// RollbackLabeledImageRequestMultiError is an error wrapping multiple
// validation errors returned by RollbackLabeledImageRequest.ValidateAll() if
// the designated constraints aren't met.
type RollbackLabeledImageRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RollbackLabeledImageRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RollbackLabeledImageRequestMultiError) AllErrors() []error { return m }

// RollbackLabeledImageRequestValidationError is the validation error returned
// by RollbackLabeledImageRequest.Validate if the designated constraints
// aren't met.
type RollbackLabeledImageRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RollbackLabeledImageRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RollbackLabeledImageRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RollbackLabeledImageRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RollbackLabeledImageRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RollbackLabeledImageRequestValidationError) ErrorName() string {
	return "RollbackLabeledImageRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RollbackLabeledImageRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRollbackLabeledImageRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RollbackLabeledImageRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RollbackLabeledImageRequestValidationError{}

// Validate checks the field values on RollbackLabeledImageResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RollbackLabeledImageResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RollbackLabeledImageResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RollbackLabeledImageResponseMultiError, or nil if none found.
func (m *RollbackLabeledImageResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *RollbackLabeledImageResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return RollbackLabeledImageResponseMultiError(errors)
	}

	return nil
}

// RollbackLabeledImageResponseMultiError is an error wrapping multiple
// validation errors returned by RollbackLabeledImageResponse.ValidateAll() if
// the designated constraints aren't met.
type RollbackLabeledImageResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RollbackLabeledImageResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RollbackLabeledImageResponseMultiError) AllErrors() []error { return m }

// RollbackLabeledImageResponseValidationError is the validation error returned
// by RollbackLabeledImageResponse.Validate if the designated constraints
// aren't met.
type RollbackLabeledImageResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RollbackLabeledImageResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RollbackLabeledImageResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RollbackLabeledImageResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RollbackLabeledImageResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RollbackLabeledImageResponseValidationError) ErrorName() string {
	return "RollbackLabeledImageResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RollbackLabeledImageResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRollbackLabeledImageResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RollbackLabeledImageResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RollbackLabeledImageResponseValidationError{}
