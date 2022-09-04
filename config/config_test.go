package config

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	filename := "test.yaml"
	Init(filename)

	fmt.Printf("web: %+v\n", GetWeb())
	fmt.Printf("web: %+v\n", GetDB())
}