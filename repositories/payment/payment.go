package repositories

import (
	"context"
	"errors"
	errWrap "payment-service/common/error"
	errConstant "payment-service/constants/error"
	"payment-service/domain/dto"
	"payment-service/domain/models"

	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

type IPaymentRepository interface {
	Create(context.Context, *dto.PaymentRequest) (*models.Payment, error)
	FindAllWithoutPagination(context.Context, *dto.PaymentRequestParam) ([]models.Payment, error)
	FindByReferenceNo(context.Context, string) (*models.Payment, error)
}

func NewPaymentRepository(db *gorm.DB) IPaymentRepository {
	return &PaymentRepository{db: db}
}

func (r *PaymentRepository) Create(ctx context.Context, req *dto.PaymentRequest) (*models.Payment, error) {
	payment := models.Payment{
		InvoiceID:   req.InvoiceID,
		Amount:      req.Amount,
		Method:      req.Method,
		ReferenceNo: req.ReferenceNo,
		PaidAt:      req.PaidAt,
	}

	err := r.db.WithContext(ctx).Create(&payment).Error
	if err != nil {
		return nil, err
	}

	return &payment, nil
}

func (r *PaymentRepository) FindAllWithoutPagination(ctx context.Context, req *dto.PaymentRequestParam) ([]models.Payment, error) {
	var payments []models.Payment
	query := r.db.WithContext(ctx)

	if req.InvoiceID != nil {
		query = query.Where("invoice_id = ?", req.InvoiceID)
	}

	if err := query.Find(&payments).Error; err != nil {
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}

	return payments, nil
}

func (r *PaymentRepository) FindByReferenceNo(ctx context.Context, referenceNo string) (*models.Payment, error) {
	var payment models.Payment

	err := r.db.WithContext(ctx).Where("reference_no = ?", referenceNo).First(&payment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}

	return &payment, nil
}
