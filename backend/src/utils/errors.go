package utils

type ErrorCode int

const JwtPrefix = "GatorStore_"

// Server Request Errors
const (
	UnknownInternalErrCode ErrorCode = iota + 800
	MissingParamsCode
	InvalidParamsCode
)

// DB errors
const (
	UnknownDbErrCode ErrorCode = iota + 900
	UnableToGetDbObj
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
