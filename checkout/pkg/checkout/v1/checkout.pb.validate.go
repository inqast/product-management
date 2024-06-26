// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: checkout.proto

package checkout

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
)

// Validate checks the field values on AddToCartRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AddToCartRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddToCartRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddToCartRequestMultiError, or nil if none found.
func (m *AddToCartRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddToCartRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUser() < 1 {
		err := AddToCartRequestValidationError{
			field:  "User",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetSku() < 1 {
		err := AddToCartRequestValidationError{
			field:  "Sku",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetCount(); val < 1 || val >= 65536 {
		err := AddToCartRequestValidationError{
			field:  "Count",
			reason: "value must be inside range [1, 65536)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return AddToCartRequestMultiError(errors)
	}

	return nil
}

// AddToCartRequestMultiError is an error wrapping multiple validation errors
// returned by AddToCartRequest.ValidateAll() if the designated constraints
// aren't met.
type AddToCartRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddToCartRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddToCartRequestMultiError) AllErrors() []error { return m }

// AddToCartRequestValidationError is the validation error returned by
// AddToCartRequest.Validate if the designated constraints aren't met.
type AddToCartRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddToCartRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddToCartRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddToCartRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddToCartRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddToCartRequestValidationError) ErrorName() string { return "AddToCartRequestValidationError" }

// Error satisfies the builtin error interface
func (e AddToCartRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddToCartRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddToCartRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddToCartRequestValidationError{}

// Validate checks the field values on DeleteFromCartRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteFromCartRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteFromCartRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteFromCartRequestMultiError, or nil if none found.
func (m *DeleteFromCartRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteFromCartRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUser() < 1 {
		err := DeleteFromCartRequestValidationError{
			field:  "User",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetSku() < 1 {
		err := DeleteFromCartRequestValidationError{
			field:  "Sku",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetCount(); val < 1 || val >= 65536 {
		err := DeleteFromCartRequestValidationError{
			field:  "Count",
			reason: "value must be inside range [1, 65536)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteFromCartRequestMultiError(errors)
	}

	return nil
}

// DeleteFromCartRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteFromCartRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteFromCartRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteFromCartRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteFromCartRequestMultiError) AllErrors() []error { return m }

// DeleteFromCartRequestValidationError is the validation error returned by
// DeleteFromCartRequest.Validate if the designated constraints aren't met.
type DeleteFromCartRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteFromCartRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteFromCartRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteFromCartRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteFromCartRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteFromCartRequestValidationError) ErrorName() string {
	return "DeleteFromCartRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteFromCartRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteFromCartRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteFromCartRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteFromCartRequestValidationError{}

// Validate checks the field values on ListCartRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListCartRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListCartRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListCartRequestMultiError, or nil if none found.
func (m *ListCartRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListCartRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUser() < 1 {
		err := ListCartRequestValidationError{
			field:  "User",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ListCartRequestMultiError(errors)
	}

	return nil
}

// ListCartRequestMultiError is an error wrapping multiple validation errors
// returned by ListCartRequest.ValidateAll() if the designated constraints
// aren't met.
type ListCartRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListCartRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListCartRequestMultiError) AllErrors() []error { return m }

// ListCartRequestValidationError is the validation error returned by
// ListCartRequest.Validate if the designated constraints aren't met.
type ListCartRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCartRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCartRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCartRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCartRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCartRequestValidationError) ErrorName() string { return "ListCartRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListCartRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCartRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCartRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCartRequestValidationError{}

// Validate checks the field values on ListCartResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListCartResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListCartResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListCartResponseMultiError, or nil if none found.
func (m *ListCartResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListCartResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListCartResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListCartResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListCartResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for TotalPrice

	if len(errors) > 0 {
		return ListCartResponseMultiError(errors)
	}

	return nil
}

// ListCartResponseMultiError is an error wrapping multiple validation errors
// returned by ListCartResponse.ValidateAll() if the designated constraints
// aren't met.
type ListCartResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListCartResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListCartResponseMultiError) AllErrors() []error { return m }

// ListCartResponseValidationError is the validation error returned by
// ListCartResponse.Validate if the designated constraints aren't met.
type ListCartResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCartResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCartResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCartResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCartResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCartResponseValidationError) ErrorName() string { return "ListCartResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListCartResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCartResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCartResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCartResponseValidationError{}

// Validate checks the field values on Item with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Item) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Item with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ItemMultiError, or nil if none found.
func (m *Item) ValidateAll() error {
	return m.validate(true)
}

func (m *Item) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Sku

	// no validation rules for Count

	// no validation rules for Name

	// no validation rules for Price

	if len(errors) > 0 {
		return ItemMultiError(errors)
	}

	return nil
}

// ItemMultiError is an error wrapping multiple validation errors returned by
// Item.ValidateAll() if the designated constraints aren't met.
type ItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ItemMultiError) AllErrors() []error { return m }

// ItemValidationError is the validation error returned by Item.Validate if the
// designated constraints aren't met.
type ItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ItemValidationError) ErrorName() string { return "ItemValidationError" }

// Error satisfies the builtin error interface
func (e ItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ItemValidationError{}

// Validate checks the field values on PurchaseRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PurchaseRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PurchaseRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PurchaseRequestMultiError, or nil if none found.
func (m *PurchaseRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PurchaseRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUser() < 1 {
		err := PurchaseRequestValidationError{
			field:  "User",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return PurchaseRequestMultiError(errors)
	}

	return nil
}

// PurchaseRequestMultiError is an error wrapping multiple validation errors
// returned by PurchaseRequest.ValidateAll() if the designated constraints
// aren't met.
type PurchaseRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PurchaseRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PurchaseRequestMultiError) AllErrors() []error { return m }

// PurchaseRequestValidationError is the validation error returned by
// PurchaseRequest.Validate if the designated constraints aren't met.
type PurchaseRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PurchaseRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PurchaseRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PurchaseRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PurchaseRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PurchaseRequestValidationError) ErrorName() string { return "PurchaseRequestValidationError" }

// Error satisfies the builtin error interface
func (e PurchaseRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPurchaseRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PurchaseRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PurchaseRequestValidationError{}

// Validate checks the field values on PurchaseResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PurchaseResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PurchaseResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PurchaseResponseMultiError, or nil if none found.
func (m *PurchaseResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *PurchaseResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrderId

	if len(errors) > 0 {
		return PurchaseResponseMultiError(errors)
	}

	return nil
}

// PurchaseResponseMultiError is an error wrapping multiple validation errors
// returned by PurchaseResponse.ValidateAll() if the designated constraints
// aren't met.
type PurchaseResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PurchaseResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PurchaseResponseMultiError) AllErrors() []error { return m }

// PurchaseResponseValidationError is the validation error returned by
// PurchaseResponse.Validate if the designated constraints aren't met.
type PurchaseResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PurchaseResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PurchaseResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PurchaseResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PurchaseResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PurchaseResponseValidationError) ErrorName() string { return "PurchaseResponseValidationError" }

// Error satisfies the builtin error interface
func (e PurchaseResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPurchaseResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PurchaseResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PurchaseResponseValidationError{}
