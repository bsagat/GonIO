package dal

type ObjectXML struct {
}

func NewObjectXMLRepo() *ObjectXML {
	return &ObjectXML{}
}

func (xml ObjectXML) RetrieveObject() error {
	return nil
}
