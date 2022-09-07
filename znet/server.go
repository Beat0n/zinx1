package znet

import "fmt"

type Server struct {
	//Server Name
	Name string
	//tcp4 or other
	IPVersion string
	//IP Address
	IP string
	//Port
	Port int
}

func (s *Server) Start() {
	fmt.Printf("[START] Server listenner at IP: %s, Port %d, is starting\n", s.IP, s.Port)
	//开启一个go去做服务端Linster业务
	go func() {
		//1 获取TCP地址
		addr, err :=
		//2 监听服务器地址

		   //监听成功

		//3 启动server网络连接
	}
}
