// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: auth/rpc/get_labeled_images.proto

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

// Validate checks the field values on GetLabeledImagesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetLabeledImagesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetLabeledImagesRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetLabeledImagesRequestMultiError, or nil if none found.
func (m *GetLabeledImagesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetLabeledImagesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for LabelerId

	// no validation rules for Page

	if m.GetSize() <= 0 {
		err := GetLabeledImagesRequestValidationError{
			field:  "Size",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := model.Category_name[int32(m.GetCategory())]; !ok {
		err := GetLabeledImagesRequestValidationError{
			field:  "Category",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetLabeledImagesRequestMultiError(errors)
	}

	return nil
}

// GetLabeledImagesRequestMultiError is an error wrapping multiple validation
// errors returned by GetLabeledImagesRequest.ValidateAll() if the designated
// constraints aren't met.
type GetLabeledImagesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetLabeledImagesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetLabeledImagesRequestMultiError) AllErrors() []error { return m }

// GetLabeledImagesRequestValidationError is the validation error returned by
// GetLabeledImagesRequest.Validate if the designated constraints aren't met.
type GetLabeledImagesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetLabeledImagesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetLabeledImagesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetLabeledImagesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetLabeledImagesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetLabeledImagesRequestValidationError) ErrorName() string {
	return "GetLabeledImagesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetLabeledImagesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetLabeledImagesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetLabeledImagesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetLabeledImagesRequestValidationError{}

// Validate checks the field values on GetLabeledImagesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetLabeledImagesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetLabeledImagesResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetLabeledImagesResponseMultiError, or nil if none found.
func (m *GetLabeledImagesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetLabeledImagesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetImages() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetLabeledImagesResponseValidationError{
						field:  fmt.Sprintf("Images[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetLabeledImagesResponseValidationError{
						field:  fmt.Sprintf("Images[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetLabeledImagesResponseValidationError{
					field:  fmt.Sprintf("Images[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetLabeledImagesResponseMultiError(errors)
	}

	return nil
}

// GetLabeledImagesResponseMultiError is an error wrapping multiple validation
// errors returned by GetLabeledImagesResponse.ValidateAll() if the designated
// constraints aren't met.
type GetLabeledImagesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetLabeledImagesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetLabeledImagesResponseMultiError) AllErrors() []error { return m }

// GetLabeledImagesResponseValidationError is the validation error returned by
// GetLabeledImagesResponse.Validate if the designated constraints aren't met.
type GetLabeledImagesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetLabeledImagesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetLabeledImagesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetLabeledImagesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetLabeledImagesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetLabeledImagesResponseValidationError) ErrorName() string {
	return "GetLabeledImagesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetLabeledImagesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetLabeledImagesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetLabeledImagesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetLabeledImagesResponseValidationError{}
