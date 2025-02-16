package dtos

type ResponseDto struct {
	ResponseCode int32
	Message      string
	Data         any
}
