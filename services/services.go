package services

type paymentRepository interface {
	CreatePayment(string) (string, error)
}

type Svc struct {
	payment paymentRepository
}

func Newsvc(payment paymentRepository) *Svc {
	return &Svc{
		payment: payment,
	}
}

func (s *Svc) CreatePayment(req string) error {
	_, err := s.payment.CreatePayment(req)
	if err != nil {
		return err
	}

	return nil
}
