package conf

type config struct {
	K8s  *kube `toml:"kube"`
	Http *http `toml:"http"`
	Grpc *grpc `toml:"grpc"`
}

func NewConfig() *config {
	return &config{
		K8s:  NewK8s(),
		Http: NewHttp(),
		Grpc: NewGrpc(),
	}
}

type kube struct {
	ConfigPath string `json:"config_path"`
}

func NewK8s() *kube {
	return &kube{
		ConfigPath: "./conf/config",
	}
}

type http struct {
	Host      string `toml:"host"`
	Port      string `toml:"port"`
	EnableSSL bool   `toml:"enable_ssl"`
	CertFile  string `toml:"cert_file"`
	KeyFile   string `toml:"key_file"`
}

func (h *http) Addr() string {
	return h.Host + ":" + h.Port
}

func NewHttp() *http {
	return &http{
		Host: "localhost",
		Port: "9999",
	}
}

type grpc struct {
	Host      string `toml:"host"`
	Port      string `toml:"port"`
	EnableSSL bool   `toml:"enable_ssl"`
	CertFile  string `toml:"cert_file"`
	KeyFile   string `toml:"key_file"`
}

func (g *grpc) Addr() string {
	return g.Host + ":" + g.Port
}

func NewGrpc() *grpc {
	return &grpc{
		Host: "localhost",
		Port: "8888",
	}
}
