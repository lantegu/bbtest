package scripttool

import (
	"time"
)

type Config struct {
	dataReadPattern int
	proportion      int
	timeOut         time.Duration
	retryCount      int
}

func (config *Config) GetDataReadPattern() int {
	return config.dataReadPattern
}

func (config *Config) GetProportion() int {
	return config.proportion
}

func (config *Config) GetTimeout() time.Duration {
	return config.timeOut
}

func (config *Config) SetTimeout(timeOut time.Duration)  {
    config.timeOut = timeOut
}

func (config *Config) GetRetryCount() int {
	return config.retryCount
}

//default config execute all datas in order, 
func NewDefaultConfig() *Config {
	return &Config{dataReadPattern: Order, proportion: 100, timeOut: time.Second,
	retryCount: 5}
}


type Option func(*Config)

//  随机选取脚本中固定比例的数据执行，常用于验证耗时较长的任务与小样本测试
// random select data in datafile and run script, it often used to verify time consuming tasks
func RandomSelectData(proportion int) Option {
	return func (c *Config) {
		c.dataReadPattern = Random
		c.proportion = proportion
	}
}

// 按顺序选取脚本中固定比例的数据执行，常用于小样本测试
// Select a fixed proportion of data in the script  to execution, 
// which is often used for small sample testing

func OrderSelectData(proportion int) Option {
	return func(c *Config) {
		c.dataReadPattern = Order
		c.proportion = proportion
	}
}

// 设置脚本超时时间, 秒级
// set script timeout, second unit
func SecondUnitTimeout(num int) Option {
	return func(c *Config) {
		c.timeOut = time.Duration(num) * time.Second
	}
}

//设置脚本超时时间
// set script timeout, uillisecond unit
func MisecondUnitTimeout(num int) Option {
	return func(c *Config) {
		c.timeOut = time.Duration(num) * time.Millisecond
	}
}

// 设置脚本重试次数
//set retry time

func SelectRetryTime(num int) Option {
	return func(c *Config) {
		c.retryCount = num
	}
}

func SetConfig(option []Option)  *Config{
	config := NewDefaultConfig()
	for _, opt := range option {
		opt(config)
	}
	return config
}