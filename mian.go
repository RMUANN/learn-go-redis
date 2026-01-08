package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	// 1. 创建客户端 (连接门店)
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // 你的地址
		Password: "",               // 没有密码就留空
		DB:       0,                // 默认数据库是0
	})
	//
	//
	//

	// 2. 必须创建一个 Context (上下文)
	// 理解：这是 Go 语言控制请求超时的机制，Redis 操作必须带上它
	ctx := context.Background()

	// 3. 写入数据 (Set)
	// 参数：上下文, Key, Value, 过期时间(0表示永不过期)
	err := rdb.Set(ctx, "my_name", "Gemini", 0).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("写入成功！")

	// 4. 读取数据 (Get)
	val, err := rdb.Get(ctx, "my_name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("读出来的名字是:", val)
}
