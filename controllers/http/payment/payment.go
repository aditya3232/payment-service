package controllers

import (
	"net/http"
	"payment-service/common/response"
	"payment-service/domain/dto"
	"payment-service/services"

	errWrap "payment-service/common/error"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PaymentController struct {
	service services.IServiceRegistry
}

type IPaymentController interface {
	Create(*gin.Context)
	FindAllWithoutPagination(*gin.Context)
}

func NewPaymentController(service services.IServiceRegistry) IPaymentController {
	return &PaymentController{service: service}
}

func (c *PaymentController) Create(ctx *gin.Context) {
	request := &dto.PaymentRequest{}
	err := ctx.ShouldBindJSON(request)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := errWrap.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusUnprocessableEntity,
			Message: &errMessage,
			Data:    errResponse,
			Err:     err,
			Gin:     ctx,
		})
		return
	}

	invoice, err := c.service.GetPayment().Create(ctx, request)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: invoice,
		Gin:  ctx,
	})
}

func (c *PaymentController) FindAllWithoutPagination(ctx *gin.Context) {
	var params dto.PaymentRequestParam
	err := ctx.ShouldBindQuery(&params)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	validate := validator.New()
	if err = validate.Struct(params); err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errorResponse := errWrap.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResp{
			Err:     err,
			Code:    http.StatusUnprocessableEntity,
			Message: &errMessage,
			Data:    errorResponse,
			Gin:     ctx,
		})
		return
	}

	invoices, err := c.service.GetPayment().FindAllWithoutPagination(ctx, &params)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: invoices,
		Gin:  ctx,
	})
}
