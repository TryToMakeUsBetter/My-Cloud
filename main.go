package main

import (
	"errors"
	"my-cloud/internal/cmd"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	var err error

	// 检查数据库是否能连接
	err = connDb()
	if err != nil {
		panic(err)
	}

	cmd.Main.Run(gctx.New())
}

// connDb 检查数据库连接是否正常
func connDb() error {
	err := g.DB().PingMaster()
	if err != nil {
		return errors.New("连接到数据库失败")
	}
	return nil
}
