package response

import "encoding/json"

//Response ... The basic response structure
type Response struct {
	Status  uint16 `json:"status"`
	Id      string `json:"id"`
	Message string `json:"message"`
}

//CreateErrorResponse ... This function will create error response
func CreateErrorResponse(message string, status uint16, data interface{}) string {
	// response := Response{Message: message,
	// 	Status: status,
	// 	Data:   data}

	// responseByte, err := json.Marshal(response)

	// if err != nil {
	// 	return "Something went wrong while parsing json"
	// }

	// return string(responseByte)

	return message
}

//CreateSuccessResponse ... This function will create success response
func CreateSuccessResponse(status uint16, id string, message string) []byte {
	response := Response{Status: status,
		Id:      id,
		Message: message}

	responseByte, err := json.Marshal(response)

	if err != nil {
		return []byte("Something went wrong while parsing json")
	}

	return responseByte
}
