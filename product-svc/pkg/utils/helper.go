package helper_product_svc

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	responsemodel_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/model/responsemodel"
	resCustomError_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/model/responsemodel/custom_error"
)

func Pagination(page string, limit string) (int, int, error) {

	pageNO, err := strconv.Atoi(page)
	if err != nil {
		return 0, 0, resCustomError_product_svc.ErrConversionOFPage
	}

	if pageNO < 1 {
		return 0, 0, resCustomError_product_svc.ErrPagination
	}

	limits, err := strconv.Atoi(limit)
	if err != nil {
		return 0, 0, resCustomError_product_svc.ErrConversionOfLimit
	}

	if limits <= 0 {
		return 0, 0, resCustomError_product_svc.ErrPageLimit
	}

	offSet := (pageNO * limits) - limits
	limits = pageNO * limits

	return offSet, limits, nil
}

func EasyValidataion(data interface{}) []responsemodel_product_svc.ValidatonError {
	var afterErrorCorection []responsemodel_product_svc.ValidatonError
	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				afterErrorCorection = append(afterErrorCorection, responsemodel_product_svc.ValidatonError{Error: fmt.Sprint(e.Error(), e.Param())})
				fmt.Println("---", e.Error())
			}
		}
	}
	return afterErrorCorection
}

func GenerateUUID() string {
	newUUID := uuid.New()

	uuidString := newUUID.String()
	return uuidString
}

func StringToUintConvertion(value string) (uint, error) {

	result, err := strconv.Atoi(value)
	if err != nil {
		return 0, errors.New("converition lead to error")
	}
	return uint(result), nil
}

func FindDiscount(originalPrice, percentageOffer float64) uint {
	discountAmount := (percentageOffer / 100) * originalPrice
	discountedPrice := originalPrice - discountAmount
	return uint(discountedPrice)
}

func GenerateReferalCode() (string, error) {
	bytes := make([]byte, 10)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes)[:5], nil
}
