package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleExec() {
	err := Exec("/bin/echo", "abc", "123")

	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}

func ExampleSudoExec() {
	err := SudoExec("/bin/echo", "abc", "123")

	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}

func ExampleRunAsUser() {
	// run echo as user some user and redirect output to /var/log/output.log
	err := RunAsUser("someuser", "/var/log/output.log", "/bin/echo", "abc", "123")

	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}

func ExampleGetFSInfo() {
	fsInfo, err := GetFSInfo()

	if err != nil {
		return
	}

	// info is slice path -> info
	for path, info := range fsInfo {
		fmt.Printf(
			"Path: %s Type: %s Device: %s Used: %d Free: %d Total: %d\n",
			path, info.Type, info.Device, info.Used, info.Free, info.Total,
		)
	}
}

func ExampleGetIOStats() {
	ioStats, err := GetIOStats()

	if err != nil {
		return
	}

	// print info for each device
	for device, info := range ioStats {
		fmt.Printf("Device: %s", device)
		fmt.Printf(
			"  ReadComplete: %d ReadMerged: %d ReadSectors: %d ReadMs: %d\n",
			info.ReadComplete, info.ReadMerged, info.ReadSectors, info.ReadMs,
		)

		fmt.Printf(
			"  WriteComplete: %d WriteMerged: %d WriteSectors: %d WriteMs: %d\n",
			info.WriteComplete, info.WriteMerged, info.WriteSectors, info.WriteMs,
		)

		fmt.Printf(
			"  IOPending: %d IOMs: %d IOQueueMs: %d\n\n",
			info.IOPending, info.IOMs, info.IOQueueMs,
		)
	}
}

func ExampleGetIOUtil() {
	// get 5 sec IO utilization
	ioUtil, err := GetIOUtil(5 * time.Second)

	if err != nil {
		return
	}

	// print utilization for each device
	for device, utilization := range ioUtil {
		fmt.Printf("Device: %s Utilization: %g\n", device, utilization)
	}
}

func ExampleGetNetworkSpeed() {
	input, output, err := GetNetworkSpeed(5 * time.Second)

	if err != nil {
		return
	}

	// print input and output speed for all interfaces
	fmt.Printf("Input: %d kb/s\n Output: %d kb/s\n", input/1024, output/1024)
}

func ExampleGetUptime() {
	uptime, err := GetUptime()

	if err != nil {
		return
	}

	// print uptime
	fmt.Printf("Uptime: %d seconds\n", uptime)
}

func ExampleGetLA() {
	la, err := GetLA()

	if err != nil {
		return
	}

	// print 1, 5 and 15 min load average
	fmt.Printf("Min1: %g Min5: %g Min15: %g\n", la.Min1, la.Min5, la.Min15)
}

func ExampleGetMemInfo() {
	memInfo, err := GetMemInfo()

	if err != nil {
		return
	}

	// print all available memory info
	fmt.Printf("MemTotal: %d\n", memInfo.MemTotal)
	fmt.Printf("MemFree: %d\n", memInfo.MemFree)
	fmt.Printf("MemUsed: %d\n", memInfo.MemUsed)
	fmt.Printf("Buffers: %d\n", memInfo.Buffers)
	fmt.Printf("Cached: %d\n", memInfo.Cached)
	fmt.Printf("Active: %d\n", memInfo.Active)
	fmt.Printf("Inactive: %d\n", memInfo.Inactive)
	fmt.Printf("SwapTotal: %d\n", memInfo.SwapTotal)
	fmt.Printf("SwapFree: %d\n", memInfo.SwapFree)
	fmt.Printf("SwapUsed: %d\n", memInfo.SwapUsed)
	fmt.Printf("SwapCached: %d\n", memInfo.SwapCached)
	fmt.Printf("Dirty: %d\n", memInfo.Dirty)
	fmt.Printf("Slab: %d\n", memInfo.Slab)
}

func ExampleGetCPUInfo() {
	cpuInfo, err := GetCPUInfo()

	if err != nil {
		return
	}

	// print all available CPU info
	fmt.Printf("User: %f\n", cpuInfo.User)
	fmt.Printf("System: %f\n", cpuInfo.System)
	fmt.Printf("Nice: %f\n", cpuInfo.Nice)
	fmt.Printf("Idle: %f\n", cpuInfo.Idle)
	fmt.Printf("Wait: %f\n", cpuInfo.Wait)
	fmt.Printf("CPU Count: %d\n", cpuInfo.Count)
}

func ExampleGetSystemInfo() {
	sysInfo, err := GetSystemInfo()

	if err != nil {
		return
	}

	// print all available system info
	fmt.Printf("Hostname: %s\n", sysInfo.Hostname)
	fmt.Printf("OS: %s\n", sysInfo.OS)
	fmt.Printf("Kernel: %s\n", sysInfo.Kernel)
	fmt.Printf("Arch: %s\n", sysInfo.Arch)
}

func ExampleWho() {
	sessions, err := Who()

	if err != nil {
		return
	}

	// print info about all active sessions
	for _, session := range sessions {
		fmt.Printf(
			"User: %s LoginTime: %v LastActivityTime: %v\n",
			session.User.Name, session.LoginTime, session.LastActivityTime,
		)
	}
}

func ExampleCurrentUser() {
	user, err := CurrentUser()

	if err != nil {
		return
	}

	// print info about current user
	fmt.Printf("UID: %d\n", user.UID)
	fmt.Printf("GID: %d\n", user.GID)
	fmt.Printf("Name: %s\n", user.Name)
	fmt.Printf("Groups: %v\n", user.GroupList())
	fmt.Printf("Comment: %s\n", user.Comment)
	fmt.Printf("Shell: %s\n", user.Shell)
	fmt.Printf("HomeDir: %s\n", user.HomeDir)
	fmt.Printf("RealUID: %d\n", user.RealUID)
	fmt.Printf("RealGID: %d\n", user.RealGID)
	fmt.Printf("RealName: %s\n", user.RealName)
	fmt.Printf("IsRoot: %t\n", user.IsRoot())
	fmt.Printf("IsSudo: %t\n", user.IsSudo())
}
