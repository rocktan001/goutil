package goutil

import (
    "errors"
    "math/rand"
    "reflect"
    "time"

    "github.com/go-redis/redis"
)

func Version() string {
    return "v1.5.1"
}
func Contain(obj interface{}, target interface{}) (bool, error) {
    targetValue := reflect.ValueOf(target)
    switch reflect.TypeOf(target).Kind() {
    case reflect.Slice, reflect.Array:
        for i := 0; i < targetValue.Len(); i++ {
            if targetValue.Index(i).Interface() == obj {
                return true, nil
            }
        }
    case reflect.Map:
        if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
            return true, nil
        }
    }

    return false, errors.New("not in array")
}

func RandNumberString(length int) string {
    var letterRunes = []rune("0123456789")
    rand.Seed(time.Now().UnixNano())
    b := make([]rune, length)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}

func RandSeq(length int) string {
    var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    rand.Seed(time.Now().UnixNano())
    b := make([]rune, length)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}
func Redis_json_set(key string, obj interface{}) {
    client := redis.NewClient(&redis.Options{
        Addr:     "www.rocktan001.com:6379",
        Password: "F96AEB124C", // no password set
        DB:       0,            // use default DB
    })
    defer client.Close()
    // client.Set(key, obj, 30*time.Second)
    if _, err := client.Ping().Result(); err != nil {
        panic(err)
    }

    err := client.Set(key, obj, -1).Err()
    if err != nil {
        panic(err)
    }
}

func Redis_json_get(key string) string {
    client := redis.NewClient(&redis.Options{
        Addr:     "www.rocktan001.com:6379",
        Password: "F96AEB124C", // no password set
        DB:       0,            // use default DB
    })
    defer client.Close()
    val, err := client.Get(key).Result()
    if err != nil {
        panic(err)
    }
    return val
}
