package global

var Config ProjectConfig

type ProjectConfig struct {
	Log LogConfig

	Service ServiceConfig

	Redis RedisConfig
}

type LogConfig struct {
	// 存储目录路径
	Dir string
	// 日志分类目录
	Category string
	// 日志级别
	Level string
	// 是否打印到控制台
	StdPrint bool `toml:"std_print"`
}

type ServiceConfig struct {
	LocalHttpAddr    string `toml:"local_http_addr"`
	MedTestAddr      string `toml:"med_test_addr"`
	MedTestStoreAddr string `toml:"med_test_store_addr"`
}

type RedisConfig struct {
	Pool map[string]RedisPoolConfig `toml:"pool"`
}

type RedisPoolConfig struct {
	Address      string `toml:"address"`
	DialPassword string `toml:"dialpassword"`
	Db           int    `toml:"db"`
	MaxIdle      int    `toml:"maxidle"`
	MaxActive    int    `toml:"maxactive"`
	Wait         bool   `toml:"wait"`
	Idletimeout  int    `toml:"idletimeout"` // 连接超时时间， 毫秒
}
