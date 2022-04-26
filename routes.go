package main

import (
	"StoreDataFB/DeviceInfo"
	"StoreDataFB/Utility"
	"StoreDataFB/entity"
	"StoreDataFB/repository"
	"StoreDataFB/uniqueNumber"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var (
	repo repository.ApiDataRepository = repository.NewRepository()
)

func StoreDeviceData(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	var record entity.ApiData

	rand.Seed(time.Now().UnixNano())
	unique_id := uniqueNumber.RandomString()          // 6 digit random string referral code
	device_number := DeviceInfo.DeviceNumber("Apple") // Generate serial number
	device_id := Utility.TrimString(device_number)    // Trim the string

	// Getting all serial numbers from the database
	device_ids, err := repo.FindAllDeviceIDs()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error"}`))
		return
	}
	// checking current serial number is present in the database or not
	if contains(device_ids, device_id) {
		response.WriteHeader(409) // data already exists
		fmt.Println(device_id, "this Device already exists")
	} else {
		record.DeviceID = device_id                  // Serial Number of the device
		record.UniqueID = unique_id                  // Referral Code for particular user referral
		record.ReferrerID = "IRAGSRAGHU"             // Referrer ID
		record.WalletAddress = "0x12WE12233EDDJJKJJ" // Wallet Address
		record.ReferralsCount = 0                    // Referrals Count
		record.RewardsEarned = 0                     // Rewards Earned
		repo.Save(&record)                           // Save all data to firestore
		response.WriteHeader(http.StatusOK)          // Send response
		json.NewEncoder(response).Encode(record)
	}
}

func ListStoreData(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	var records []entity.ApiData
	records, err := repo.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error Generating code"}`))
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(records)
}

// Record exists or not
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
