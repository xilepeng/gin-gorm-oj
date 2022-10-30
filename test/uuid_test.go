package test

import (
	"github.com/satori/go.uuid"
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	s := uuid.NewV4().String()
	println(s, len(s))
	println(uuid.NewV4().String(), len(uuid.NewV4().String()))

}
