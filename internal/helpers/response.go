package helpers

type ResponseStruct struct {
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}

func Response(data interface{}, description string) ResponseStruct {
	return ResponseStruct{Description: description, Data: data}
}
