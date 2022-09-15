package model

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/jdxj/cyber-wagon/config"
	"github.com/jdxj/cyber-wagon/internal/util"
)

var (
	path string
)

func TestMain(t *testing.M) {
	config.Init("../../../config/test.yaml")
	util.InitDB(config.GetDB())
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	path = filepath.Join(homeDir, "tmp")
	os.Exit(t.Run())
}

func TestStorage_WriteFile(t *testing.T) {
	stg := &Storage{
		path: path,
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

func TestStorage_ReadFile(t *testing.T) {
	stg := &Storage{path: path}
	fi, err := stg.ReadFile(context.Background(), 1, 2)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	r, err := fi.Open()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	defer r.Close()

	_, err = io.Copy(os.Stdout, r)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}