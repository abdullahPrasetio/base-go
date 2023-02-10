package constants

const (
	ErrorValidate     = "03"
	ErrorNotFound     = "05"
	ErrorDB           = "15"
	ErrorQueryDB      = "16"
	ErrorThirdParty   = "20"
	ErrorGeneral      = "99"
	Success           = "00"
	ErrorPageNotFound = "404"
)

var mappingError = map[string]string{
	"03":  "Error Validate Input Request",
	"05":  "Data Not Found",
	"404": "Page Not Found",
	"15":  "DB connection error",
	"16":  "DB Query error",
	"20":  "Third Party Error",
	"99":  "General Error",
}

type mapError struct {
	ResponseCode string
	Description  string
}

func GetError(key string) *mapError {
	errors := mappingError
	var mapError mapError
	newError := errors[key]
	if len(newError) <= 0 {
		mapError.ResponseCode = "99"
		mapError.Description = "General Error"
		return &mapError
	}

	mapError.ResponseCode = key
	mapError.Description = newError
	return &mapError
}
