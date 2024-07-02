package config

type Config struct {
	Env         string `yaml:"env" env-default:"local"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	//TODO https://www.youtube.com/watch?v=rCJvW2xgnk0&t=1246s
	//HTTPServer  `yaml:"http_server"`
}
