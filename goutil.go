package goutil

import (
    "errors"
    "math/rand"
    "reflect"
    "time"

    "github.com/go-redis/redis"
)

func Version() string {
    return "v1.6.2"
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
    if _, err := client.Ping().Result(); err != nil {
        panic(err)
    }

    val, err := client.Get(key).Result()
    if err != nil {
        panic(err)
    }
    return val
}

func Redis_json_SAdd(key, value string) {
    client := redis.NewClient(&redis.Options{
        Addr:     "www.rocktan001.com:6379",
        Password: "F96AEB124C", // no password set
        DB:       0,            // use default DB
    })
    defer client.Close()
    if _, err := client.Ping().Result(); err != nil {
        panic(err)
    }

    err := client.SAdd(key, value).Err()
    if err != nil {
        panic(err)
    }

}

func Redis_json_SRem(key, value string) {
    client := redis.NewClient(&redis.Options{
        Addr:     "www.rocktan001.com:6379",
        Password: "F96AEB124C", // no password set
        DB:       0,            // use default DB
    })
    defer client.Close()
    if _, err := client.Ping().Result(); err != nil {
        panic(err)
    }

    err := client.SRem(key, value).Err()
    if err != nil {
        panic(err)
    }

}

func Redis_json_SMembers(key string) []string {
    client := redis.NewClient(&redis.Options{
        Addr:     "www.rocktan001.com:6379",
        Password: "F96AEB124C", // no password set
        DB:       0,            // use default DB
    })
    defer client.Close()
    if _, err := client.Ping().Result(); err != nil {
        panic(err)
    }

    es, _ := client.SMembers(key).Result()
    return es

}

func Redis_json_pub(key string, obj interface{}) {
    client := redis.NewClient(&redis.Options{
        Addr:     "www.rocktan001.com:6379",
        Password: "F96AEB124C", // no password set
        DB:       0,            // use default DB
    })
    defer client.Close()
    if _, err := client.Ping().Result(); err != nil {
        panic(err)
    }

    if err := client.Publish(key, obj).Err(); err != nil {
        panic(err)
    }

}

func Redis_json_sub(key string) string {
    // var message *redis.Message
    client := redis.NewClient(&redis.Options{
        Addr:     "www.rocktan001.com:6379",
        Password: "F96AEB124C", // no password set
        DB:       0,            // use default DB
    })
    defer client.Close()
    if _, err := client.Ping().Result(); err != nil {
        panic(err)
    }
    sub := client.Subscribe(key)
    iface, err := sub.Receive()
    if err != nil {
        // handle error
        panic(err)
    }

    // Should be *Subscription, but others are possible if other actions have been
    // taken on sub since it was created.
    switch iface.(type) {
    case *redis.Subscription:
        // fmt.Println("Subscription")
    case *redis.Message:
        // received first message
        // fmt.Println("Message")
    case *redis.Pong:
        // pong received
        // fmt.Println("Pong")
    default:
        // handle error
    }

    ch := sub.Channel()
    // fmt.Println(ch)
    for msg := range ch {
        // message = msg
        // fmt.Println(msg.Channel, ":", msg.Payload)
        return msg.Payload
    }
    return "error"
}
