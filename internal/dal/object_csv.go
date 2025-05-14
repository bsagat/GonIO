package dal

import "GonIO/internal/domain"

type ObjectCSV struct {
}

func NewObjectCSVRepo() *ObjectCSV {
	return &ObjectCSV{}
}

var _ domain.ObjectDal = (*ObjectCSV)(nil)

func (xml ObjectCSV) RetrieveObject() error {
	return nil
}

/*
func List_Object(bucketname string) error {
	dir := *models.Dir
	isUnique, err := u.CheckUniqueCSV(dir+"/buckets.csv", bucketname, true)
	if err != nil {
		return err
	}
	if isUnique {
		return errors.New("bucket file is not exist")
	}
	metafile, err := os.OpenFile(dir+"/"+bucketname+"/objects.csv", os.O_RDWR, 0o666)
	if err != nil {
		return err
	}
	defer metafile.Close()
	reader := csv.NewReader(metafile)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}
	var Objects []models.Object
	for i, record := range records {
		if i == 0 {
			continue
		}
		if len(record) < 4 {
			continue
		}
		Objects = append(Objects, models.Object{
			ObjectKey:    record[0],
			Size:         record[1],
			ContentType:  record[2],
			LastModified: record[3],
		})
	}
	objectlist := models.ObjectsList{Objects: Objects}
	w.Header().Set("Content-Type", "application/xml")
	if err := xml.NewEncoder(w).Encode(objectlist); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return nil

	}
	return nil
}*/
