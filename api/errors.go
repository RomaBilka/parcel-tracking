package api

type Error struct {
	Message string `json:"message"`
}

func IsWarn(httpCode int) bool {
	return httpCode >= 400 && httpCode <= 499
}

func IsErr(httpCode int) bool {
	return httpCode >= 500
}
