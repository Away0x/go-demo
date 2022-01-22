package cmd

import (
	"gohub/pkg/console"
	"gohub/pkg/redis"
	"time"

	"github.com/spf13/cobra"
)

var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "Likes the Go Playground, but running at our application context",
	Run:   runPlay,
}

// 可在该处调试代码, 调试完成后请记得清除测试代码
// 有点像 go.dev/play/ ，但是运行在应用环境中，数据库、配置、Redis 等系统服务都已初始化，可以放心使用
func runPlay(cmd *cobra.Command, args []string) {
	// 存进去 redis 中
	redis.Redis.Set("hello", "hi from redis", 10*time.Second)
	// 从 redis 里取出
	console.Success(redis.Redis.Get("hello"))
}
