package deviceInfo

import (
	"fmt"
	"os/exec"
	"strings"
)

// MAC device serial number
func DarwinDevice(num string) string {
	var uniq_number string
	out, _ := exec.Command("/usr/sbin/ioreg", "-l").Output() // err ignored for brevity
	for _, l := range strings.Split(string(out), "\n") {
		if strings.Contains(l, "IOPlatformSerialNumber") {
			s := strings.Split(l, " ")
			uniq_number = s[len(s)-1]
			fmt.Println(uniq_number)
			break
		}
	}
	return uniq_number
}

// Windows device serial number
func WindowsDevice(num string) string {
	var uniq_number string
	out, _ := exec.Command("wmic", "bios", "get", "serialnumber").Output() // err ignored for brevity
	for _, l := range strings.Split(string(out), "\n") {
		if strings.Contains(l, "SerialNumber") {
			s := strings.Split(l, " ")
			uniq_number = s[len(s)-1]
			fmt.Println(uniq_number)
			break
		}
	}
	return uniq_number
}

// Linux device serial number
func LinuxDevice(num string) string {
	var uniq_number string
	out, _ := exec.Command("cat", "/sys/class/dmi/id/product_serial").Output() // err ignored for brevity
	for _, l := range strings.Split(string(out), "\n") {
		if strings.Contains(l, "product_serial") {
			s := strings.Split(l, " ")
			uniq_number = s[len(s)-1]
			fmt.Println(uniq_number)
			break
		}
	}
	return uniq_number
}

// Android device serial number
func AndroidDevice(num string) string {
	var uniq_number string
	out, _ := exec.Command("getprop", "ro.serialno").Output() // err ignored for brevity
	for _, l := range strings.Split(string(out), "\n") {
		if strings.Contains(l, "ro.serialno") {
			s := strings.Split(l, " ")
			uniq_number = s[len(s)-1]
			fmt.Println(uniq_number)
			break
		}
	}
	return uniq_number
}

// IOS device serial number
func IOSDevice(num string) string {
	var uniq_number string
	out, _ := exec.Command("system_profiler", "SPHardwareDataType").Output() // err ignored for brevity
	for _, l := range strings.Split(string(out), "\n") {
		if strings.Contains(l, "Serial Number") {
			s := strings.Split(l, " ")
			uniq_number = s[len(s)-1]
			fmt.Println(uniq_number)
			break
		}
	}
	return uniq_number
}
