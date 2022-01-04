package main

import (
	"errors"
	"fmt"

	"go.uber.org/multierr"
)

func foo() error {
	var err error
	err = multierr.Append(err, ErrFirst)
	err = multierr.Append(err, &SecondErr{code: 1})
	err = multierr.Append(err, &SecondErr{code: 2})
	err = multierr.Append(err, ErrThird)
	err = multierr.Append(err, errors.New("other"))

	return err
}

func bar() error {
	var err error
	err = multierr.Append(err, nil)
	err = multierr.Append(err, nil)
	return err
}

func bulp1() (err error) {
	defer multierr.AppendInvoke(&err, multierr.Invoke(Close))
	return ErrThird
}

func bulp2() (err error) {
	defer multierr.AppendInvoke(&err, multierr.Invoke(Succeed))
	return nil
}

func Close() error {
	return ErrFirst
}

func Succeed() error {
	return nil
}

var ErrFirst error = errors.New("first")

type SecondErr struct {
	code int
}

func (e *SecondErr) Error() string {
	return fmt.Sprintf("second, code: %d", e.code)
}

var ErrThird error = errors.New("third")
