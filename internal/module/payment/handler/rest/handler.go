package rest

import (
	integration "product-service/internal/integration/midtrans"
	"product-service/internal/integration/midtrans/entity"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type paymentHandler struct {
	midtrans integration.MidtransContract
}

func NewHandler() *paymentHandler {
	midtrans := integration.NewMidtransContract()

	return &paymentHandler{
		midtrans: midtrans,
	}
}

func (h *paymentHandler) Register(router fiber.Router) {
	router.Post("/payments", h.createPayment)
}

func (h *paymentHandler) createPayment(c *fiber.Ctx) error {
	var (
		req = entity.CreatePaymentRequest{}
		ctx = c.Context()
	)

	if err := c.BodyParser(&req); err != nil {
		log.Error().Err(err).Msg("failed to parse request body")
		return err
	}

	resp, err := h.midtrans.CreatePayment(ctx, &req)
	if err != nil {
		log.Error().Err(err).Msg("failed to create payment")
		return err
	}

	statusCodeMap := map[string]int{
		"05":  fiber.StatusInternalServerError,
		"201": fiber.StatusCreated,
		"200": fiber.StatusOK,
		"400": fiber.StatusBadRequest,
		"401": fiber.StatusUnauthorized,
		"406": fiber.StatusNotAcceptable,
		"500": fiber.StatusInternalServerError,
		"503": fiber.StatusServiceUnavailable,
		"900": fiber.StatusInternalServerError,
	}

	statusCode, ok := statusCodeMap[resp.StatusCode]
	if !ok {
		statusCode = fiber.StatusInternalServerError
	}

	return c.Status(statusCode).JSON(resp)
}
