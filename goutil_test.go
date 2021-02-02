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

func TestRandSeq(t *testing.T) {
	fmt.Println(RandSeq(15))
}
func TestVersion(t *testing.T) {
	fmt.Println(Version())
}
func TestGetPhysicalID(t *testing.T) {
	fmt.Println(GetPhysicalID())
}

func TestRedis_Set(t *testing.T) {
	Redis_json_SAdd("online", GetPhysicalID())
	fmt.Println(Redis_json_SMembers("online"))
	Redis_json_SRem("online", GetPhysicalID())
	fmt.Println(Redis_json_SMembers("online"))
}
