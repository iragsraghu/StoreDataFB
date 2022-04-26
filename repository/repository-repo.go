package repository

import (
	"StoreDataFB/entity"
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type ApiDataRepository interface {
	Save(post *entity.ApiData) (*entity.ApiData, error)
	FindAll() ([]entity.ApiData, error)
	FindAllDeviceIDs() ([]string, error)
}

type repo struct{}

// NewRepository
func NewRepository() ApiDataRepository {
	return &repo{}
}

const (
	projectId      string = "gofrontier2"
	collectionName string = "api_datas"
)

func (*repo) Save(referral *entity.ApiData) (*entity.ApiData, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create firesotre client: %v", err)
		return nil, err
	}

	defer client.Close()
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"DeviceID":       referral.DeviceID,
		"UniqueID":       referral.UniqueID,
		"ReferrerID":     referral.ReferrerID,
		"WalletAddress":  referral.WalletAddress,
		"ReferralsCount": referral.ReferralsCount,
		"RewardsEarned":  referral.RewardsEarned,
	})

	if err != nil {
		log.Fatalf("Failed to generate device info: %v", err)
		return nil, err
	}

	return referral, nil
}

func (*repo) FindAll() ([]entity.ApiData, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create firesotre client: %v", err)
		return nil, err
	}

	defer client.Close()
	var records []entity.ApiData
	iter := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of datas: %v", err)
			return nil, err
		}
		record := entity.ApiData{
			DeviceID:       doc.Data()["DeviceID"].(string),
			UniqueID:       doc.Data()["UniqueID"].(string),
			ReferrerID:     doc.Data()["ReferrerID"].(string),
			WalletAddress:  doc.Data()["WalletAddress"].(string),
			ReferralsCount: doc.Data()["ReferralsCount"].(int64),
			RewardsEarned:  doc.Data()["RewardsEarned"].(int64),
		}
		records = append(records, record)
	}
	return records, nil
}

func (*repo) FindAllDeviceIDs() ([]string, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create firesotre client: %v", err)
		return nil, err
	}

	defer client.Close()
	var deviceIDs []string
	iter := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of datas: %v", err)
			return nil, err
		}
		record := entity.ApiData{
			DeviceID: doc.Data()["DeviceID"].(string),
		}
		deviceIDs = append(deviceIDs, record.DeviceID)
	}
	return deviceIDs, nil
}
