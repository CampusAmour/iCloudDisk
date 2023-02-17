package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"sync"
)

// GlobalConfig 基础配置类
type GlobalConfig struct {
	LogCfg *LogConfig `yaml:"log"`
}

var (
	cfg  *GlobalConfig
	lock = new(sync.RWMutex)
)

// Parse 解析配置文件内容
func Parse(path string) error {
	cfg = new(GlobalConfig)
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer func() {
		err := file.Close()
		if err != nil {
			return
		}
	}()

	decoder := yaml.NewDecoder(file)
	lock.Lock()
	defer lock.Unlock()
	return decoder.Decode(cfg)
}

// Config 获取配置类对象(单例模式)
func Config() *GlobalConfig {
	lock.RLock()
	defer lock.RUnlock()
	return cfg
}
