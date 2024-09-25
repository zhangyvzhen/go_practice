# go 练习项目

## 项目结构
```sh
my-go-project/
│
├── cmd/                # 各个命令的入口
│   └── myapp/         # 主应用程序
│       └── main.go    # 主函数
|
├── bin/               # 编译后的主程序
│   └── myapp/         # 主应用程序
│       └── main.exe    # 主函数
│
├── internal/          # 仅在本项目内部使用的包
│   ├── mypackage1/
│   │   └── mypackage1.go
│   └── mypackage2/
│       └── mypackage2.go
│
├── pkg/               # 可供外部使用的库（如果需要）
│   └── mylib/
│       └── mylib.go
│
├── api/               # API 定义（如 proto 文件）
│   └── api.proto
│
│
├── configs/           # 配置文件
│   └── config.yaml
│
├── scripts/           # 辅助脚本
│   └── setup.sh
│
├── tests/             # 测试文件
│   ├── mypackage1_test.go
│   └── mypackage2_test.go
│
├── go.mod             # Go 模块文件
└── README.md          # 项目说明


```
