package gocpu

import (
	"bufio"
	"os"
	"strings"
)

type CpuInfo struct {
	CPUNum        int    `json:"cpuNum"`
	Vendor        string `json:"vendor"`
	ModelName     string `json:"modelName"`
	CPUMHz        string `json:"cpum_hz"`
	LogicCpuCores int    `json:"logicCpuCores"`
}

func GetCpuInfo() (CpuInfo, error) {
	cpuinfo_file := "/proc/cpuinfo"
	return getPhysicalCPUInfo(cpuinfo_file)
}

func getPhysicalCPUInfo(cpuinfo_file string) (CpuInfo, error) {
	cpuInfo := CpuInfo{}
	file, err := os.Open(cpuinfo_file)
	if err != nil {
		return cpuInfo, err
	}
	defer file.Close()
	physicalCPUs := make(map[string]struct{})
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "physical id") {
			count++
			fields := strings.Split(line, ":")
			//log.Info(fields[1])
			if len(fields) >= 2 {
				physicalCPUs[fields[1]] = struct{}{}
			}
		}
		if strings.HasPrefix(line, "model name") {
			fields := strings.Split(line, ":")
			//log.Info(fields[1])
			if len(fields) >= 2 {
				cpuInfo.ModelName = strings.TrimSpace(fields[1])
			}
		}
		if strings.HasPrefix(line, "vendor_id") {
			fields := strings.Split(line, ":")
			//log.Info(fields[1])
			if len(fields) >= 2 {
				cpuInfo.Vendor = strings.TrimSpace(fields[1])
			}
		}
		if strings.HasPrefix(line, "cpu MHz") {
			fields := strings.Split(line, ":")
			//log.Info(fields[1])
			if len(fields) >= 2 {
				cpuInfo.CPUMHz = strings.TrimSpace(fields[1])
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return cpuInfo, err
	}
	cpuInfo.CPUNum = len(physicalCPUs)
	cpuInfo.LogicCpuCores = count
	return cpuInfo, nil
}
