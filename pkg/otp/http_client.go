package otp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type OtpClient interface {
	GenerateOtp(req GenerateOtpRequest) (*GenerateOtpSuccessResponse, error)
	VerifyOtp(req VerifyOtpRequest) (*VerifyOtpSuccessResponse, error)
}
type otpClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewOtpClient(baseURL string) OtpClient {
	return &otpClient{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
	}
}
func postRequestWithJSON[T any](client *http.Client, url string, reqBody interface{}) (*T, error) {
	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	resp, err := client.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	// return resp, nil
	defer resp.Body.Close()
	var respData T
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return &respData, nil
}
func (c *otpClient) GenerateOtp(req GenerateOtpRequest) (*GenerateOtpSuccessResponse, error) {
	url := fmt.Sprintf("%s/otp/generate", c.BaseURL)
	resp, err := postRequestWithJSON[GenerateOtpSuccessResponse](c.HTTPClient, url, req)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, fmt.Errorf("received nil response from server")
	}
	return resp, nil
}

func (c *otpClient) VerifyOtp(req VerifyOtpRequest) (*VerifyOtpSuccessResponse, error) {
	url := fmt.Sprintf("%s/otp/verify", c.BaseURL)
	resp, err := postRequestWithJSON[VerifyOtpSuccessResponse](c.HTTPClient, url, req)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, fmt.Errorf("received nil response from server")
	}
	return resp, nil
}
