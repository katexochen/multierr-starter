package main

import (
	"errors"
	"fmt"

	"github.com/hashicorp/go-multierror"
)

func foo() error {
	var err error
	err = multierror.Append(err, ErrFirst)
	err = multierror.Append(err, &SecondErr{code: 1})
	err = multierror.Append(err, &SecondErr{code: 2})
	err = multierror.Append(err, ErrThird)
	err = multierror.Append(err, errors.New("other"))

	return err
}

func bar() error {
	var err error
	err = multierror.Append(err, nil)
	err = multierror.Append(err, nil)
	return err
}

func baz() error {
	var multierr *multierror.Error
	multierr = multierror.Append(multierr, nil)
	return multierr.ErrorOrNil()
}

func bulp1() (err error) {
	defer appendDeferredErr(&err, ErrFirst)
	return ErrThird
}

func bulp2() (err error) {
	defer appendDeferredErr(&err, nil)
	return nil
}

func appendDeferredErr(target *error, err error) {
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
