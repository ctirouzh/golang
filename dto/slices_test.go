package dto_test

import (
	"github.com/ctioruzh/golang/dto"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestList(t *testing.T) {
	var (
		inputs   = []int{1, 3, 5, 7}
		expected = []string{"1", "3", "5", "7"}
	)
	// List (type zero): []I -> []O
	list := dto.List[string](inputs, strconv.Itoa)
	assert.NotNil(t, list)
	for i := range list {
		assert.Equal(t, expected[i], list[i])
	}
	// List (type I): []I -> []*O
	list1 := dto.List[*string](inputs, dto.Mapper[int, string](strconv.Itoa).PtrO())
	assert.NotNil(t, list1)
	for i := range list1 {
		assert.Equal(t, expected[i], *list1[i])
	}

	// Convert inputs to pointer inputs: []I -> []*I
	inputs2 := dto.List[*int](inputs, dto.PointerFunc[int]())
	assert.NotNil(t, inputs2)
	for i := range inputs2 {
		assert.Equal(t, &inputs[i], inputs2[i])
		assert.Equal(t, inputs[i], *inputs2[i])
	}
	// List (type II): []*I -> []O
	list2 := dto.List[string](inputs2, dto.Mapper[int, string](strconv.Itoa).PtrI())
	assert.NotNil(t, list2)
	for i := range list2 {
		assert.Equal(t, expected[i], list2[i])
	}
	// List (type III): []*I -> []*O
	list3 := dto.List[*string](inputs2, dto.Mapper[int, string](strconv.Itoa).PtrIO())
	assert.NotNil(t, list3)
	for i := range list3 {
		assert.Equal(t, expected[i], *list3[i])
	}
}
