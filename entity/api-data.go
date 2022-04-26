package entity

// Post data structure
type ApiData struct {
	DeviceID       string `json:"device_id"` // Serial number
	UniqueID       string `json:"unique_id"` // ReferralCode
	ReferrerID     string `json:"referrer_id"`
	WalletAddress  string `json:"wallet_address"`
	ReferralsCount int64  `json:"referrals_count"`
	RewardsEarned  int64  `json:"rewards_earned"`
}
