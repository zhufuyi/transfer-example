// code generated by https://github.com/zhufuyi/sponge

package config

import (
	"github.com/zhufuyi/sponge/pkg/conf"
)

var config *Config

func Init(configFile string, fs ...func()) error {
	config = &Config{}
	return conf.Parse(configFile, config, fs...)
}

func Show(hiddenFields ...string) string {
	return conf.Show(config, hiddenFields...)
}

func Get() *Config {
	if config == nil {
		panic("config is nil")
	}
	return config
}

func Set(conf *Config) {
	config = conf
}

type Config struct {
	App        App          `yaml:"app" json:"app"`
	Consul     Consul       `yaml:"consul" json:"consul"`
	Etcd       Etcd         `yaml:"etcd" json:"etcd"`
	Grpc       Grpc         `yaml:"grpc" json:"grpc"`
	GrpcClient []GrpcClient `yaml:"grpcClient" json:"grpcClient"`
	Jaeger     Jaeger       `yaml:"jaeger" json:"jaeger"`
	Logger     Logger       `yaml:"logger" json:"logger"`
	NacosRd    NacosRd      `yaml:"nacosRd" json:"nacosRd"`
}

type Consul struct {
	Addr string `yaml:"addr" json:"addr"`
}

type Etcd struct {
	Addrs []string `yaml:"addrs" json:"addrs"`
}

type Jaeger struct {
	AgentHost string `yaml:"agentHost" json:"agentHost"`
	AgentPort int    `yaml:"agentPort" json:"agentPort"`
}

type ServerSecure struct {
	CaFile   string `yaml:"caFile" json:"caFile"`
	CertFile string `yaml:"certFile" json:"certFile"`
	KeyFile  string `yaml:"keyFile" json:"keyFile"`
	Type     string `yaml:"type" json:"type"`
}

type App struct {
	CacheType             string  `yaml:"cacheType" json:"cacheType"`
	EnableCircuitBreaker  bool    `yaml:"enableCircuitBreaker" json:"enableCircuitBreaker"`
	EnableHTTPProfile     bool    `yaml:"enableHTTPProfile" json:"enableHTTPProfile"`
	EnableLimit           bool    `yaml:"enableLimit" json:"enableLimit"`
	EnableMetrics         bool    `yaml:"enableMetrics" json:"enableMetrics"`
	EnableStat            bool    `yaml:"enableStat" json:"enableStat"`
	EnableTrace           bool    `yaml:"enableTrace" json:"enableTrace"`
	Env                   string  `yaml:"env" json:"env"`
	Host                  string  `yaml:"host" json:"host"`
	Name                  string  `yaml:"name" json:"name"`
	RegistryDiscoveryType string  `yaml:"registryDiscoveryType" json:"registryDiscoveryType"`
	TracingSamplingRate   float64 `yaml:"tracingSamplingRate" json:"tracingSamplingRate"`
	Version               string  `yaml:"version" json:"version"`
}

type GrpcClient struct {
	EnableLoadBalance     bool   `yaml:"enableLoadBalance" json:"enableLoadBalance"`
	Host                  string `yaml:"host" json:"host"`
	Name                  string `yaml:"name" json:"name"`
	Port                  int    `yaml:"port" json:"port"`
	RegistryDiscoveryType string `yaml:"registryDiscoveryType" json:"registryDiscoveryType"`
}

type Grpc struct {
	EnableToken  bool         `yaml:"enableToken" json:"enableToken"`
	HTTPPort     int          `yaml:"httpPort" json:"httpPort"`
	Port         int          `yaml:"port" json:"port"`
	ReadTimeout  int          `yaml:"readTimeout" json:"readTimeout"`
	ServerSecure ServerSecure `yaml:"serverSecure" json:"serverSecure"`
	WriteTimeout int          `yaml:"writeTimeout" json:"writeTimeout"`
}

type Logger struct {
	Format string `yaml:"format" json:"format"`
	IsSave bool   `yaml:"isSave" json:"isSave"`
	Level  string `yaml:"level" json:"level"`
}

type NacosRd struct {
	IPAddr      string `yaml:"ipAddr" json:"ipAddr"`
	NamespaceID string `yaml:"namespaceID" json:"namespaceID"`
	Port        int    `yaml:"port" json:"port"`
}
