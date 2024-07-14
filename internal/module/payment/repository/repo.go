package repository

import "product-service/internal/module/payment/ports"

type repo struct {
}

func NewPaymentRepository() ports.PaymentRepository {
	return &repo{}
}
