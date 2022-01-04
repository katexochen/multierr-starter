package main

import (
	"testing"

	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/assert"
)

func TestFoo(t *testing.T) {
	assert := assert.New(t)
	err := foo()
	assert.Error(err)
	assert.ErrorIs(err, ErrFirst)
	assert.ErrorIs(err, ErrThird)
	var secondErr *SecondErr
	assert.ErrorAs(err, &secondErr)
	assert.Equal(1, secondErr.code) // errors.As returns first error
}

func TestBar(t *testing.T) {
	assert := assert.New(t)
	err := bar()
	assert.Error(err)
	var multierr *multierror.Error
	assert.ErrorAs(err, &multierr)
}

func TestBaz(t *testing.T) {
	assert := assert.New(t)
	err := baz()
	assert.NoError(err)
}

func TestLupz(t *testing.T) {
	assert := assert.New(t)
	err := lupz()
	assert.Error(err)
	assert.ErrorIs(err, ErrFirst)
	assert.ErrorIs(err, ErrThird)
}

func TestBulp1(t *testing.T) {
	assert := assert.New(t)
	err := bulp1()
	assert.Error(err)
	assert.ErrorIs(err, ErrFirst)
	assert.ErrorIs(err, ErrThird)
}

func TestBulp2(t *testing.T) {
	assert := assert.New(t)
	err := bulp2()
	assert.NoError(err)
	assert.Nil(err)
}
