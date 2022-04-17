package request

type InsertProduct struct {
	Title              string `json:"title"`
	Districtfk         uint32 `json:"districtfk"`
	AddressDetail      string `json:"address_detail"`
	ProductStateFk     uint8  `json:"product_statefk"`
	ProductDescription string `json:"product_description"`
	Price              uint   `json:"price"`
	CategoryFk         uint16 `json:"categoryfk"`
}

type UpdateProduct struct {
	ID                 uint   `json:"id"`
	Title              string `json:"title"`
	ProductStateFk     uint8  `json:"product_statefk"`
	ProductDescription string `json:"product_description"`
	Price              uint   `json:"price"`
	CategoryFk         uint16 `json:"categoryfk"`
}
