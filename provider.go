package provider

// provider interface.
type Provider interface {
	Ping() error  //ping连接是否还有效
	Close() error //关闭连接
}
