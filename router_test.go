package cli_test

import (
	"fmt"
	"github.com/gregoryv/cli"
	"testing"
)

var nop = func() error { return nil }

func TestNewRouter(t *testing.T) {
	if r := cli.NewRouter(); r == nil {
		t.Error("NewRouter() returned nil")
	}
}

func TestAdd(t *testing.T) {
	name := "help"
	if err := cli.Add(name, nop); err != nil {
		t.Error("should return a Command")
	}
	if err := cli.Add(name, nop); err == nil {
		t.Errorf("should fail if trying to register under same name")
	}
}

func TestRouter_List(t *testing.T) {
	router := cli.NewRouter()
	if result := router.List(); len(result) != 0 {
		t.Errorf("should be empty by default")
	}
	router.Add("b", nop)
	router.Add("c", nop)
	router.Add("a", nop)
	exp := []string{"a", "b", "c"}
	res := router.List()
	if fmt.Sprintf("%v", res) != fmt.Sprintf("%v", exp) {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}
