package base

import (
	"fmt"
)

// Error 类型
const (
	// 请求错误
	ErrInvalidInput      = "INVALID_INPUT"
	ErrSerializationFail = "SERIALIZATION_FAIL"

	// 网络错误
	ErrNetworkFail     = "NETWORK_FAIL"
	ErrHTTPRequestFail = "HTTP_REQUEST_FAIL"

	// 响应错误
	ErrInvalidResponse     = "INVALID_RESPONSE"
	ErrDeserializationFail = "DESERIALIZATION_FAIL"

	// API 错误
	ErrAPIError       = "API_ERROR"
	ErrAuthError      = "AUTH_ERROR"
	ErrRateLimitError = "RATE_LIMIT_ERROR"
)

// BingAdsError 表示 Bing Ads API 错误
type BingAdsError struct {
	Code    string
	Message string
	Cause   error
}

// Error 实现 error 接口
func (e *BingAdsError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("BingAds API 错误 [%s]: %s - %v", e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("BingAds API 错误 [%s]: %s", e.Code, e.Message)
}

// NewError 创建一个新的 BingAdsError
func NewError(code, message string, cause error) *BingAdsError {
	return &BingAdsError{
		Code:    code,
		Message: message,
		Cause:   cause,
	}
}

// IsAuthError 检查是否为认证错误
func IsAuthError(err error) bool {
	if bingErr, ok := err.(*BingAdsError); ok {
		return bingErr.Code == ErrAuthError
	}
	return false
}

// IsRateLimitError 检查是否为速率限制错误
func IsRateLimitError(err error) bool {
	if bingErr, ok := err.(*BingAdsError); ok {
		return bingErr.Code == ErrRateLimitError
	}
	return false
}

// IsAPIError 检查是否为 API 错误
func IsAPIError(err error) bool {
	if bingErr, ok := err.(*BingAdsError); ok {
		return bingErr.Code == ErrAPIError
	}
	return false
}
