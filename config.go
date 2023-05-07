package common

import (
	"github.com/go-micro/plugins/v4/config/source/consul"
	"go-micro.dev/v4/config"
	"strconv"
)

// GetConsulConfig 设置配置中心
func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	consulSource := consul.NewSource(
		// 设置配置中心的地址
		consul.WithAddress(host+":"+strconv.FormatInt(port, 10)),
		// 设置前缀
		consul.WithPrefix(prefix),
		// 是否移除前缀,这里设置为true,表示可以不带前缀直接获取对应的配置
		consul.StripPrefix(true),
	)
	microConfig, err := config.NewConfig()
	if err != nil {
		return microConfig, err
	}

	// 加载配置
	err = microConfig.Load(consulSource)
	return microConfig, err
}
