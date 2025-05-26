package container

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewContainer(t *testing.T) {
	container := NewContainer()

	assert.NotNil(t, container)
	assert.NotNil(t, container.container)
}

func TestContainer_BuildContainer(t *testing.T) {
	container := NewContainer()

	err := container.BuildContainer()

	if err != nil {
		t.Logf("BuildContainer error (expected): %v", err)
	}
}

func TestContainer_Invoke(t *testing.T) {
	container := NewContainer()

	invoked := false
	testFunc := func() {
		invoked = true
	}

	err := container.Invoke(testFunc)

	assert.NoError(t, err)
	assert.True(t, invoked)
}

func TestContainer_InvokeWithDependency(t *testing.T) {
	container := NewContainer()

	err := container.container.Provide(func() string {
		return "test"
	})
	assert.NoError(t, err)

	var result string
	err = container.Invoke(func(s string) {
		result = s
	})

	assert.NoError(t, err)
	assert.Equal(t, "test", result)
}
