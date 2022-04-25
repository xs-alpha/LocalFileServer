package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(Appconfig)

type Appconfig struct {
	Name       string `mapstructure:"name"`
	Mode       string `mapstructure:"mode"`
	Version    string `mapstructure:"version"`
	Port       int    `mapstructure:"port"`
	*LogConfig `mapstructure:"log"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackUps int    `mapstructure:"max_backups"`
}

func Init(file string) (err error) {
	// 这个指定的是main可执行文件所在的目录与yaml文件相对关系。不是setting.go文件和yaml相对关系
	fmt.Println("file:", file)
	if len(file) == 0 {
		viper.SetConfigFile(file)
	} else {
		viper.SetConfigFile("./config.yaml")
	}
	// 设置默认值
	//viper.SetDefault("fileDir", "")
	//viper.SetConfigName("config") // 配置文件名称中没有扩展名

	// 指定配置文件类型，专用于从远程获取配置信息时指定配置文件类型
	//viper.SetConfigType("yaml")           // 如果配置文件没有扩展名，这需要配置此项
	viper.AddConfigPath(".")              // 查找配置文件所在路径
	viper.AddConfigPath("$HOME/.appname") // 多次调用以添加多个搜索路径
	err = viper.ReadInConfig()            // 查找并读取配置文件
	if err != nil {
		//fmt.Println("main: viper initial failed:", err.Error())
		return err
		//panic(fmt.Errorf("main: config initial failed"))
	}

	// 反序列化
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Println("[*]setting: viper unmarshal failed:", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("[*]config changed:...:%v\n", in.Name)
		// 更新
		if err2 := viper.Unmarshal(Conf); err2 != nil {
			fmt.Println("[*]setting: config update failed:", err.Error())
		}
	})
	return
}
