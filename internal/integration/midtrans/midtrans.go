package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"payment-service/internal/infrastructure"
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
	var (
		response    entity.CreatePaymentResponse
		midtransCfg        = infrastructure.Envs.Midtrans
		envCfg      string = infrastructure.Envs.App.Environtment
		ChargeURL   string
	)

	if envCfg == "production" {
		ChargeURL = midtransCfg.Production.ChargeURL
	} else {
		ChargeURL = midtransCfg.Sandbox.ChargeURL
	}

	bytesReq, err := json.Marshal(req)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal request")
		return response, err
	}

	request, err := http.NewRequest(http.MethodPost, ChargeURL, bytes.NewBuffer(bytesReq))
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
