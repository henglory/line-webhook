package server

func errMissformat(raw string) errorResponse {
	var res errorResponse
	res.ResponseCode = -1
	res.Reason = "MISS FORMAT REQUEST"
	res.RawRequest = raw
	return res
}
