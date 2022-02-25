package utils

type ErrorCode int

const JwtPrefix = "GatorStore_"

// Server Request Errors
const (
	MissingParamsCode ErrorCode = iota + 800
	InvalidParamsCode
)

// 1000 JWT
const (
	MissingJwtTokenCode ErrorCode = iota + 1000
	InvalidJwtTokenCode
)

// 9000 Error with Youtube
const (
	MissingAccessTokenCode ErrorCode = iota + 9000
	InvalidAccessTokenCode
	InvalidGoogleCode
)

func SetErrorMsg(msg string) map[string]interface{} {
	returnMsg := make(map[string]interface{})
	returnMsg["msg"] = msg
	return returnMsg
}
