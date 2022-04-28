package demo_controller

import (
	"StoreDataFB/DeviceInfo"
	"StoreDataFB/Utility"
	"StoreDataFB/repository"
	"html/template"
	"net/http"
)

func Index(response http.ResponseWriter, request *http.Request) {
	tmplt, err1 := template.ParseFiles("template/demo_controller/index.html")
	if err1 != nil {
		panic(err1)
	}
	tmplt.Execute(response, nil)
}

func ReferralsPage(response http.ResponseWriter, request *http.Request) {
	var current_device_id string
	device_number := DeviceInfo.DeviceNumber("Apple") // Generate serial number
	device_id := Utility.TrimString(device_number)    // Trim the string
	tmplt, err1 := template.ParseFiles("template/demo_controller/referrals.html")
	if err1 != nil {
		panic(err1)
	}
	records, err2 := repository.NewRepository().FindAll()
	if err2 != nil {
		panic(err2)
	}
	for _, record := range records {
		if record.DeviceID == device_id {
			current_device_id = record.UniqueID
		}
	}
	if current_device_id != "" {
		tmplt.Execute(response, current_device_id)
	} else {
		tmplt.Execute(response, "Device not found")
	}
}
