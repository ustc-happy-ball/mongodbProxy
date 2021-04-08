package response

type BaseResponse struct {
	*AddResponse
	*DeleteResponse
	*FindResponse
	*UpdateResponse
}
