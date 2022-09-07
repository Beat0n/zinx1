package ziface

type IServer interface {
	//Start server
	Start()
	//Stop server
	Stop()
	//Start serve
	Serve()
}
