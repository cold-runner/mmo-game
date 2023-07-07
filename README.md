# mmo-game
基于zinx-v1.0版本开发的简单mmo游戏服务器
## 功能
1. 玩家上线广播
2. 世界聊天
3. 同步玩家信息（移动）
4. 玩家下线广播
## 快速开始
### 克隆项目
`git clone https://github.com/cold-runner/mmo-game`
### 编译文件
进入到项目主目录后编译运行
`go run main.go`
### 启动客户端文件
打开mmo_game_u3d_client文件夹中的`client.exe`
![image](https://github.com/cold-runner/mmo-game/assets/86516312/03d90694-1171-4b93-823b-ea6ad920165b)

IP地址填：`127.0.0.1`（或者你自己本地的局域网地址，服务器默认监听0.0.0.0，需要修该的话只需在`zinx.json`中修改）

Port填：`7777`（默认监听端口，修改同上）

可以启动多个客户端以实现多人在线

### 效果展示
![image](https://github.com/cold-runner/mmo-game/assets/86516312/c49a7b0d-5b52-4891-b9f1-0931601af65e)
