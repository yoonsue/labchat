package inmem

import (
	"github.com/garyburd/redigo/redis"
	"github.com/yoonsue/labchat/model/menu"
)

// MenuRepository struct definition
type MenuRepository struct {
	key   string
	value []string
}

// // Menus ...
// type Menus []string

// NewMenuRepository does several services according to InMemoryDB
func NewMenuRepository() menu.Repository {
	return &MenuRepository{
		key:   "",
		value: []string{""},
	}
}

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

var pool = newPool()

func setCli(key string, value string) {
	c := pool.Get()
	c.Send("SET", key, value)
	_, err := c.Receive()
	if err != nil {
		panic(err.Error())
	}
}

func getCli(key string) (value interface{}) {
	c := pool.Get()
	c.Send("GET", key)
	r, err := c.Receive()
	if err != nil {
		panic(err.Error())
	}
	return r
}

// ///////////////////////////////////////////////////////////
// var client *redisA.Client

// func init() {
// 	client = redisA.NewClient(&redisA.Options{
// 		Addr:         ":6379",
// 		DialTimeout:  10 * time.Second,
// 		ReadTimeout:  10 * time.Second,
// 		WriteTimeout: 10 * time.Second,
// 		PoolSize:     10,
// 		PoolTimeout:  10 * time.Second,
// 	})
// 	client.FlushDB()
// }

// // Client ...
// type Client struct {
// 	// contains filtered or unexported fields

// }

// // ExampleNewClient ...
// func ExampleNewClient() {
// 	client := redisA.NewClient(&redisA.Options{
// 		Addr:     "localhost:6379",
// 		Password: "",
// 		DB:       0,
// 	})
// 	pong, err := client.Ping().Result()
// 	fmt.Println(pong, err)
// }

// // ExampleClient ...
// func ExampleClient(cafeteria string, menuList []string) {
// 	for _, menu := range menuList {
// 		err := client.Set(cafeteria, menu, 0).Err()
// 		if err != nil {
// 			panic(err)
// 		}
// 	}

// 	val, err := client.Get(cafeteria).Result()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(cafeteria, ": ", val)

// }
