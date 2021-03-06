package Beans

import (
	"github.com/dengpju/higo-gin/higo"
	"github.com/gomodule/redigo/redis"
	"higo/app/Controllers/V3"
	"higo/app/Services"
)

type MyBean struct {
	higo.Bean
}

func NewMyBean() *MyBean {
	return &MyBean{}
}

func (this *MyBean) DemoService() *Services.DemoService {
	return Services.NewDemoService()
}

func (this *MyBean) NewGorm() *higo.Gorm {
	return higo.NewGorm()
}

func (this *MyBean) NewRedisPool() *redis.Pool {
	return higo.RedisPool
}

func (this *MyBean) NewRedisAdapter() *higo.RedisAdapter {
	return higo.NewRedisAdapter()
}

func (this *MyBean) NewRedisController() *V3.RedisController {
	return V3.NewRedisController()
}

func (this *MyBean) NewDemoController() *V3.DemoController {
	return V3.NewDemoController()
}
