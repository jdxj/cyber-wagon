package model

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/jdxj/cyber-wagon/config"
	"github.com/jdxj/cyber-wagon/internal/util"
)

func TestStorage_WriteFile(t *testing.T) {
	config.Init("../../../config/test.yaml")
	util.InitDB(config.GetDB())
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("home: %s\n", homeDir)

	stg := &Storage{
		path: filepath.Join(homeDir, "tmp"),
	}

	filename := "hello.test"
	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	defer f.Close()

	fi, err := stg.WriteFile(context.Background(), 2, 2, "hello.test", f)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("file: %+v", fi)
}
