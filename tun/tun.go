package tun

import (
	"github.com/songgao/water"
	"log"
)

// CreateTun 创建 tun接口，默认运行结束后删除
func CreateTun() *water.Interface {
	conf := water.Config{DeviceType: water.TAP}
	conf.Name = "tap-tunvs0"
	ifce, err := water.New(conf)
	if err != nil {
		log.Fatalln("create tun interface failed: ", err)
	}
	log.Println("interface created: ", ifce.Name())
	return ifce
}
