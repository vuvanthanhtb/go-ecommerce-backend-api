package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddOne(t *testing.T) {
	var (
		input  = 1
		output = 2
	)

	actual := AddOne(input)

	if actual != output {
		t.Errorf("AddOne(%d); input %d; output %d", actual, input, output)
	}

	// assert.Equal(t, AddOne(input), output, fmt.Errorf("AddOne(%d); output %d", input, output).Error())
	// assert.NotEqual(t, AddOne(input+1), output)
	// assert.Nil(t, nil, nil)
}

func TestAddTwo(t *testing.T) {
	var (
		input  = 1
		output = 2
	)

	actual := AddTwo(input)

	if actual != output {
		t.Errorf("AddTwo(%d); input %d; output %d", actual, input, output)
	}
}

func TestRequire(t *testing.T) {
	require.Equal(t, 2, 2)
	fmt.Println("TestRequire")
}

func TestAssert(t *testing.T) {
	assert.Equal(t, 1, 1)
	fmt.Println("TestAssert")
}

// go test -coverprofile=coverage.out
// go tool cover -html=coverage.out -o coverage.html
// open coverage.html
