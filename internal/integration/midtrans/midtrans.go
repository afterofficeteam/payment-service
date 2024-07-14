package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"payment-service/internal/integration/midtrans/entity"

	"github.com/rs/zerolog/log"
)

type MidtransContract interface {
	CreatePayment(ctx context.Context, req *entity.CreatePaymentRequest) (entity.CreatePaymentResponse, error)
}

type midtrans struct {
}

func NewMidtransContract() MidtransContract {
	return &midtrans{}
}

func (m *midtrans) CreatePayment(ctx context.Context, req *entity.CreatePaymentRequest) (entity.CreatePaymentResponse, error) {
	log.Debug().Any("request", req).Msg("create payment request")
	var response entity.CreatePaymentResponse

	bytesReq, err := json.Marshal(req)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal request")
		return response, err
	}
	request, err := http.NewRequest(http.MethodPost, "https://api.sandbox.midtrans.com/v2/charge", bytes.NewBuffer(bytesReq))
	if err != nil {
		log.Error().Err(err).Msg("failed to create request")
		return response, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Basic "+req.BasicAuthHeader)

	// create http client
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to do request")
		return response, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
		return response, err
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response")
		return response, err
	}

	return response, nil
}
