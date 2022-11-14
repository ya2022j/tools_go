package main

import (
 "context"
 "fmt"
 "github.com/gin-gonic/gin"
 "github.com/go-redis/redis/v8"
)

var (
 Redis *redis.Client
)

func InitRedis() *redis.Client {
 rdb := redis.NewClient(&redis.Options{
  Addr:     "127.0.0.1:6379",
  Password: "xxxx", // no password set
  DB:       10,                 // use default DB
  PoolSize: 10,
 })
 result := rdb.Ping(context.Background())
 fmt.Println("redis ping:", result.Val())
 if result.Val() != "PONG" {
  // 连接有问题
  return nil
 }
 return rdb
}

func indexHandler(c *gin.Context) {
 name, err := Redis.Get(c, "name").Result()
 if err != nil {
  c.JSON(500, gin.H{
   "msg": err.Error(),
  })
 } else {
  c.JSON(200, gin.H{
   "name": name,
  })
 }

}

func main() {
 e := gin.New()
 e.Use(gin.Recovery())

 Redis = InitRedis()
 e.GET("/", indexHandler)
 e.Run(":8080")
}
