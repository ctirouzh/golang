package dto_test

import (
	"github.com/ctioruzh/golang/dto"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestAddress(t *testing.T) {
	x := "str"
	pX := dto.Address(x)
	assert.Equal(t, x, *pX)

	assert.NotPanics(t, func() { dto.Address[*bool](nil) })
	pBool := dto.Address[*bool](nil)
	log.Printf(">>>> pBool: %v, is nil: %v, *pBool: %v *pBool is nil: %v", pBool, pBool == nil, *pBool, *pBool == nil)
}
