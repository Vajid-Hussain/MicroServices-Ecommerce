package handler_product_svc

import (
	"context"
	"fmt"
	"io/ioutil"
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

func (h *ProductHandler) AddProduct(c *gin.Context) {
	var inventoryDetails requestmodel_product_svc_clind.InventoryReq

	if err := c.ShouldBind(&inventoryDetails); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	fmt.Println("--", inventoryDetails)

	image, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	inventoryDetails.Image = image

	fileOpened, err := image.Open()
	if err != nil {
		fmt.Println("err", err)
	}
	defer fileOpened.Close()

	// Read the contents of the file into a byte slice
	fileBytes, err := ioutil.ReadAll(fileOpened)
	if err != nil {
		fmt.Println("--", err)
	}

	result, err := h.Clind.AddProduct(context.Background(), &pb.AddProductRequest{
		Productname:        inventoryDetails.Productname,
		Description:        inventoryDetails.Description,
		BrandId:            uint32(inventoryDetails.BrandID),
		CategoryId:         uint32(inventoryDetails.CategoryID),
		Mrp:                uint32(inventoryDetails.Mrp),
		Discount:           uint32(inventoryDetails.Discount),
		Units:              uint32(inventoryDetails.Units),
		Os:                 inventoryDetails.Os,
		CellularTechnology: inventoryDetails.CellularTechnology,
		Ram:                uint32(inventoryDetails.Ram),
		Screensize:         float32(inventoryDetails.Screensize),
		BatteryCapacity:    uint32(inventoryDetails.Batterycapacity),
		Processor:          inventoryDetails.Processor,
		Image:              fileBytes,
		Saleprice:          uint32(inventoryDetails.Saleprice),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *ProductHandler) GetAllProduct(c *gin.Context) {
	page := c.Query("page")
	limit := c.DefaultQuery("limit", "1")

	result, err := h.Clind.GetAllProduct(context.Background(), &pb.GetAllProductRequest{
		Offset: page,
		Limit:  limit,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, result)
}

func (h *ProductHandler) CrateCart(c *gin.Context) {
	var cart *requestmodel_product_svc_clind.Cart

	err := c.BindJSON(&cart)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
	}

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"err":"can't fetch user id from context"})
		return
	}

	cart.UserID = userID
}