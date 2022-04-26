package DeviceInfo

import (
	"fmt"
	"os/exec"
	"strings"
)

func DeviceNumber(num string) string {
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
