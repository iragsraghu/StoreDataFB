package utility

import (
	"StoreDataFB/deviceInfo"
	"runtime"
)

func TrimString(s string) string {
	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}
	return s
}

func DeviceTypes() string {
	var device_id string
	if runtime.GOOS == "darwin" {
		device_number := deviceInfo.DarwinDevice("Apple") // Generate serial number
		device_id = TrimString(device_number)             // Trim the string
	}
	if runtime.GOOS == "linux" {
		device_number := deviceInfo.LinuxDevice("Linux") // Generate serial number
		device_id = TrimString(device_number)            // Trim the string
	}
	if runtime.GOOS == "windows" {
		device_number := deviceInfo.WindowsDevice("Windows") // Generate serial number
		device_id = TrimString(device_number)                // Trim the string
	}
	if runtime.GOOS == "android" {
		device_number := deviceInfo.AndroidDevice("Android") // Generate serial number
		device_id = TrimString(device_number)                // Trim the string
	}
	if runtime.GOOS == "ios" {
		device_number := deviceInfo.IOSDevice("IOS") // Generate serial number
		device_id = TrimString(device_number)        // Trim the string
	}
	return device_id
}
