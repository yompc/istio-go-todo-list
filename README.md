### 为了学习使用Istio所制作的Demo 

```
├── istio                 Istio配置
│   ├── app               k8s部署应用配置
│   ├── gateway           Istio网关配置
│   └── route             Istio路由配置
├── oauth2-server         鉴权服务(Spring Boot)
│   
├── todo-consumer         todo-list消费者（Go Gin,Grpc Consumer）
├─────authorize           调用oauth2-server鉴权的地方
├───────ResourceServer.go    此处仅为测试应该要实现Trie Tree
│   
├── todo-list.sql         MySql初始化文件
│
├── todo-select-provider (todo 查询数据 Grpc Provider)
│  
├── todo-write-provider  (todo 写数据 Grpc Provider)
└── web                  前端
    └── todo-list        预览地址:  https://yom-todo.vercel.app/
```

![](https://raw.github.com/yompc/repositpry/master/istio-go-todo-list/DEMO.PNG)