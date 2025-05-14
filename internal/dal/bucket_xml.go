package dal

import (
	"GonIO/internal/domain"
	"encoding/csv"
	"os"
)

type BucketXML struct {
	BucketMetaPath string
}

func NewBucketXMLRepo() *BucketXML {
	return &BucketXML{BucketMetaPath: domain.BucketsPath}
}

func (xml BucketXML) GetBucketList() ([]domain.Bucket, error) {
	file, err := os.Open(xml.BucketMetaPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var buckets []domain.Bucket
	for _, record := range records {
		buckets = append(buckets, domain.Bucket{
			Name:             record[0],
			CreationTime:     record[1],
			LastModifiedTime: record[2],
			Status:           record[3],
		})
	}
	return buckets, nil
}
