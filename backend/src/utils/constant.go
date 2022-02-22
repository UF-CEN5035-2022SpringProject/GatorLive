package utils

const JwtPrefix = "GatorStore_"

type ErrorCode int

const (
	MissingParams ErrorCode = iota + 800
	InvalidParams
)
