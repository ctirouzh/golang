package dto_test

import (
	"github.com/ctioruzh/golang/dto"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestMapper(t *testing.T) {
	var (
		input    = 3
		expected = "3"
	)
	mapper := dto.Mapper[int, string](strconv.Itoa)
	if got := mapper(input); assert.NotEmpty(t, got) {
		assert.Equal(t, expected, got)
	}

	mapper1 := mapper.PtrO()
	if got := mapper1(input); assert.NotNil(t, got) {
		assert.Equal(t, expected, *got)
	}

	mapper2 := mapper.PtrI()
	if got := mapper2(&input); assert.NotEmpty(t, got) {
		assert.Equal(t, expected, got)
	}

	mapper3 := mapper.PtrIO()
	if got := mapper3(&input); assert.NotNil(t, got) {
		assert.Equal(t, expected, *got)
	}
}

func TestPointerAndIndirectFunc(t *testing.T) {
	x := "str"
	// Pointer to x
	pX := dto.PointerFunc[string]()(x)
	if assert.NotNil(t, pX) {
		assert.Equal(t, x, *pX)
	}
	// value of the pointer to x
	vpX := dto.IndirectFunc[string](false)(pX)
	assert.Equal(t, x, vpX)
	// pointer of pointer to x
	ppX := dto.PointerFunc[*string]()(pX)
	if assert.NotNil(t, ppX) {
		assert.Equal(t, x, **ppX)
		assert.Equal(t, pX, dto.IndirectFunc[*string](false)(ppX))
	}
	// safe mode && nil input >>> zero value
	assert.Equal(t, nil, dto.IndirectFunc[any](true)(nil))
	assert.Equal(t, false, dto.IndirectFunc[bool](true)(nil))
	assert.Equal(t, "", dto.IndirectFunc[string](true)(nil))
	assert.Equal(t, struct{}{}, dto.IndirectFunc[struct{}](true)(nil))

	// unsafe mode && nil input >>> panics
	assert.Panics(t, func() { dto.IndirectFunc[any](false)(nil) })
}
