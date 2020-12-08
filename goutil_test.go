package goutil

import (
	"fmt"
	"testing"
)

func TestContain(t *testing.T) {
	target := []string{"rock", "rocktan"}
	b, err := Contain("rock", target)
	fmt.Println(b, err)
}

func TestRandNumberString(t *testing.T) {
	fmt.Println(RandNumberString(10))
}

func TestRedis_json_set(t *testing.T) {
	Redis_json_set("hello", "hello")
}

func TestRedis_json_get(t *testing.T) {
	fmt.Println(Redis_json_get("hello"))
}
