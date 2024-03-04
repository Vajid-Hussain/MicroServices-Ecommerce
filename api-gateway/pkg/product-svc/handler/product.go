package handler_product_svc

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/product-svc/pb"
	requestmodel_product_svc_clind "github.com/vajid-hussain/mobile-mart-microservice/pkg/product-svc/requestmodel"
)

type ProductHandler struct {
	Clind pb.ProductServiceClient
}

func NewProductHandler(clind pb.ProductServiceClient) *ProductHandler {
	return &ProductHandler{Clind: clind}
}

func (h *ProductHandler) CreateCategory(c *gin.Context) {
	var categoryDetails requestmodel_product_svc_clind.Category

	err := c.BindJSON(&categoryDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
	}

	result, err := h.Clind.CreateCatogory(context.Background(), &pb.CreateCategoryRequest{
		Name: categoryDetails.Name,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *ProductHandler) GetAllCategory(c *gin.Context) {

	result, err := h.Clind.GetAllCategory(context.Background(), &pb.GetAllCategoryRequest{
		Ofset: c.Query("page"),
		Limit: c.DefaultQuery("limit", "1"),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *ProductHandler) CreateBrand(c *gin.Context) {
	var brandDetails requestmodel_product_svc_clind.Brand

	err := c.BindJSON(&brandDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
	}

	result, err := h.Clind.CreateBrand(context.Background(), &pb.CreateBrandRequest{
		Name: brandDetails.Name,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *ProductHandler) GetAllBrand(c *gin.Context) {

	result, err := h.Clind.GetAllBrand(context.Background(), &pb.GetAllBrandRequest{
		Offset: c.Query("page"),
		Limit:  c.DefaultQuery("limit", "1"),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, result)
}
