package main

import (
	"errors"
	"fmt"

	"github.com/hashicorp/go-multierror"
)

func foo() error {
	var multierr error
	multierr = multierror.Append(multierr, ErrFirst)
	multierr = multierror.Append(multierr, &SecondErr{code: 1})
	multierr = multierror.Append(multierr, &SecondErr{code: 2})
	multierr = multierror.Append(multierr, ErrThird)
	multierr = multierror.Append(multierr, errors.New("other"))

	return multierr
}

func bar() error {
	var multierr error
	multierr = multierror.Append(multierr, nil)
	return multierr
}

func baz() error {
	var multierr *multierror.Error
	multierr = multierror.Append(multierr, nil)
	return multierr.ErrorOrNil()
}

func lupz() (err error) {
	var multierr *multierror.Error
	defer returnErr(&multierr, &err)

	multierr = multierror.Append(multierr, nil)
	defer appendDeferredErr(&multierr, ErrFirst)

	multierr = multierror.Append(multierr, ErrThird)
	return
}

func appendDeferredErr(multiErr **multierror.Error, err error) {
	*multiErr = multierror.Append(*multiErr, err)
}

func returnErr(err **multierror.Error, target *error) {
	*target = (*err).ErrorOrNil()
}

func bulp1() (err error) {
	defer appendDeferredErr2(&err, ErrFirst)
	return ErrThird
}

func bulp2() (err error) {
	defer appendDeferredErr2(&err, nil)
	return nil
}

func appendDeferredErr2(target *error, err error) {
	*target = multierror.Append(*target, err).ErrorOrNil()
}

var ErrFirst error = errors.New("first")

type SecondErr struct {
	code int
}

func (e *SecondErr) Error() string {
	return fmt.Sprintf("second, code: %d", e.code)
}

var ErrThird error = errors.New("third")
