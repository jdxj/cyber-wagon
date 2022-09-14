package config

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	filename := "test.yaml"
	Init(filename)

	fmt.Printf("config: %+v\n", c)
}
