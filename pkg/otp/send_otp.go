package otp

type GenerateOtpRequest struct {
	// Định danh của người dùng (email, số điện thoại, username)
	Identifier string `json:"identifier"`
	// Loại định danh
	Type *string `json:"type,omitempty"`
	// Mục đích sử dụng OTP
	Purpose string `json:"purpose"`
}
type VerifyOtpRequest struct {
	// Định danh của người dùng (email, số điện thoại, username)
	Identifier string `json:"identifier"`
	// Mã OTP cần xác thực
	Otp string `json:"otp"`
	// Loại định danh
	Type *string `json:"type,omitempty"`
	// Mục đích sử dụng OTP
	Purpose string `json:"purpose"`
}

// response
type TooManyRequestsResponseData struct {
	// Thời gian còn lại (giây) trước khi có thể yêu cầu OTP mới
	RemainingTime *int32 `json:"remainingTime,omitempty"`
}
type TooManyRequestsResponse struct {
	Success *bool                        `json:"success,omitempty"`
	Message *string                      `json:"message,omitempty"`
	Data    *TooManyRequestsResponseData `json:"data,omitempty"`
}
type GenerateOtpSuccessResponse struct {
	Success *bool                           `json:"success,omitempty"`
	Message *string                         `json:"message,omitempty"`
	Data    *GenerateOtpSuccessResponseData `json:"data,omitempty"`
}
type GenerateOtpSuccessResponseData struct {
	// Mã OTP (chỉ trả về trong môi trường phát triển)
	Otp        *string `json:"otp,omitempty"`
	Identifier *string `json:"identifier,omitempty"`
	Type       *string `json:"type,omitempty"`
	Purpose    *string `json:"purpose,omitempty"`
	// Thời gian hết hạn (giây)
	ExpiresIn *int32 `json:"expiresIn,omitempty"`
}
type VerifyOtpSuccessResponse struct {
	Success *bool                         `json:"success,omitempty"`
	Message *string                       `json:"message,omitempty"`
	Data    *VerifyOtpSuccessResponseData `json:"data,omitempty"`
}
type VerifyOtpSuccessResponseData struct {
	Verified *bool `json:"verified,omitempty"`
}
type ErrorResponse struct {
	Success *bool   `json:"success,omitempty"`
	Message *string `json:"message,omitempty"`
	// Thông tin lỗi chi tiết
	Error *string `json:"error,omitempty"`
}
