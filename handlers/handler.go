package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

type svc interface {
	CreatePayment(req string) error
}

type Handler struct {
	svc       svc
	validator *validator.Validate
}

func NewHandler(svc svc, validator *validator.Validate) *Handler {
	return &Handler{svc: svc, validator: validator}
}

func (h *Handler) CreatePayment(w http.ResponseWriter, r *http.Request) {
	// req := r.Context().Value("req").(string)
	req := "req"
	err := h.svc.CreatePayment(req)
	if err != nil {
		// helper.HandleResponse(w, http.StatusInternalServerError, helper.Response{
		// 	Message: helper.ERROR_MESSAGE,
		// 	Error:   err.Error(),
		// })
		return
	}

	// helper.HandleResponse(w, http.StatusOK, helper.Response{
	// 	Message: helper.SUCCESS_MESSSAGE,
	// })
}
