package main

import "github.com/cold-runner/zinx-learn/zinx-v1.0/znet"

func main() {
	//创建服务器句柄
	s := znet.NewServer()

	//启动服务
	s.Serve()
}
