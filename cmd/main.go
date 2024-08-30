package main

import (
	"encoding/json"
	"github.com/golangrustnode/gocpu/gocpu"
	"github.com/golangrustnode/log"
)

func main() {
	cpuinfo, err := gocpu.GetCpuInfo()
	if err != nil {
		log.Error(err)
	}
	d, _ := json.Marshal(cpuinfo)
	log.Info(string(d))
}
