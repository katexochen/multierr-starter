package main

import (
	"testing"

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
	assert.NoError(err)
	assert.Nil(err)
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
