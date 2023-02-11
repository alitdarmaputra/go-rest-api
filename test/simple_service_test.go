package test

import (
	"testing"

	"github.com/alitdarmaputra/belajar-golang-rest-api/simple"
	"github.com/stretchr/testify/assert"
)

func TestSimpleServicePass(t *testing.T) {
	simpleSerivce, err := simple.InitializedService(false)
	assert.Nil(t, err)
	assert.NotNil(t, simpleSerivce)
}

func TestSimpleServiceErro(t *testing.T) {
	simpleSerivce, err := simple.InitializedService(true)
	assert.Nil(t, simpleSerivce)
	assert.NotNil(t, err)
}
