package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var redisHost = "127.0.0.1:6379"

func newRedis() redis.Conn {
	c, err := redis.Dial("tcp", redisHost)
	if err != nil {
		fmt.Printf("Redis Connected Error: %s", err)
		return nil
	} else {
		fmt.Println("Redis Connected Succssful")
		return c
	}
}

// type People struct {
// 	Name string `json:"name_tILE"`
// 	Age  int    `json:"AGE_SIZE"`
// }

// func TestStructToJson() string {
// 	p := People{
// 		Name: "Liu xx",
// 		Age:  18,
// 	}

// 	//Person 结构体转换为对应的 Json
// 	jsonBytes, err := json.Marshal(p)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return string(jsonBytes)
// }

// func main() {
// 	c := newRedis()
// 	s := TestStructToJson()
// 	v, err := c.Do("SET", "news", s)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(v)
// 	v, err = redis.String(c.Do("GET", "news"))
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(v)
// }
