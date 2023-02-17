package config

// LogConfig 日志相关配置类
type LogConfig struct {
	Dir           string `yaml:"log_dir"`
	FileName      string `yaml:"log_file"`
	Level         string `yaml:"log_level"`
	CallStack     bool   `yaml:"log_call_stack"`
	Rotate        int64  `yaml:"log_rotate_interval"`
	RotationCount uint   `yaml:"log_rotation_count"`
}
