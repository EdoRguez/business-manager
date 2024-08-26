package services

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/EdoRguez/business-manager/product-svc/pkg/models"
	product "github.com/EdoRguez/business-manager/product-svc/pkg/pb"
	repo "github.com/EdoRguez/business-manager/product-svc/pkg/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductService struct {
	Repo *repo.ProductRepo
	product.UnimplementedProductServiceServer
}

func (s *ProductService) CreateProduct(ctx context.Context, req *product.CreateProductRequest) (*product.CreateProductResponse, error) {
	fmt.Println("Product Service :  CreateProduct")
	fmt.Println("Product Service :  CreateProduct - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	createProductParams := models.Product{
		CompanyId:     req.CompanyId,
		Name:          req.Name,
		Description:   req.Description,
		Sku:           req.Sku,
		Quantity:      req.Quantity,
		Price:         req.Price,
		Images:        req.Images,
		ProductStatus: req.ProductStatus,
		CreatedAt:     time.Now(),
		ModifiedAt:    time.Now(),
	}

	c, err := s.Repo.CreateProduct(ctx, createProductParams)
	if err != nil {
		fmt.Println("Product Service :  CreateProduct - ERROR")
		fmt.Println(err.Error())
		return &product.CreateProductResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("Product Service :  CreateProduct - SUCCESS")
	fmt.Println(c)
	fmt.Println("------------------")

	return &product.CreateProductResponse{
		Status: http.StatusCreated,
		Id:     c.Hex(),
	}, nil
}

func (s *ProductService) GetProduct(ctx context.Context, req *product.GetProductRequest) (*product.GetProductResponse, error) {
	fmt.Println("Product Service :  GetProduct")
	fmt.Println("Product Service :  GetProduct - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return &product.GetProductResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	p, err := s.Repo.GetProduct(ctx, objID)
	if err != nil {
		fmt.Println("Product Service :  GetProduct - ERROR")
		fmt.Println(err.Error())

		return &product.GetProductResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("Product Service :  GetProduct - SUCCESS")
	return &product.GetProductResponse{
		Id:            p.Id.Hex(),
		CompanyId:     p.CompanyId,
		Name:          p.Name,
		Description:   p.Description,
		Sku:           p.Sku,
		Quantity:      p.Quantity,
		Price:         p.Price,
		Images:        p.Images,
		ProductStatus: p.ProductStatus,
		Status:        http.StatusOK,
	}, nil
}

func (s *ProductService) GetProducts(ctx context.Context, req *product.GetProductsRequest) (*product.GetProductsResponse, error) {
	fmt.Println("Product Service :  GetProducts")
	fmt.Println("Product Service :  GetProducts - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	params := repo.GetProductsParams{
		CompanyId: req.CompanyId,
		Name:      req.Name,
		Sku:       req.Sku,
		Limit:     req.Limit,
		Offset:    req.Offset,
	}

	p, err := s.Repo.GetProducts(ctx, params)
	if err != nil {
		fmt.Println("Product Service :  GetProducts - ERROR")
		fmt.Println(err.Error())
		return &product.GetProductsResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	products := make([]*product.GetProductResponse, 0, len(p))
	for _, v := range p {
		products = append(products, &product.GetProductResponse{
			Id:            v.Id.Hex(),
			CompanyId:     v.CompanyId,
			Name:          v.Name,
			Description:   v.Description,
			Sku:           v.Sku,
			Quantity:      v.Quantity,
			Price:         v.Price,
			Images:        v.Images,
			ProductStatus: v.ProductStatus,
		})
	}

	fmt.Println("Product Service :  GetProducts - SUCCESS")
	return &product.GetProductsResponse{
		Products: products,
		Status:   http.StatusOK,
	}, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, req *product.UpdateProductRequest) (*product.UpdateProductResponse, error) {
	fmt.Println("Product Service :  UpdateProduct")
	fmt.Println("Product Service :  UpdateProduct - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	params := models.Product{
		Name:        req.Name,
		Description: req.Description,
		Sku:         req.Sku,
		Quantity:    req.Quantity,
		Price:       req.Price,
		Images:      req.Images,
	}

	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return &product.UpdateProductResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	err = s.Repo.UpdateProduct(ctx, objID, params)
	if err != nil {
		fmt.Println("Product Service :  UpdateProduct - ERROR")
		fmt.Println(err.Error())

		return &product.UpdateProductResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("Product Service :  UpdateProduct - SUCCESS")
	return &product.UpdateProductResponse{
		Status: http.StatusNoContent,
	}, nil
}

func (s *ProductService) UpdateProductStatus(ctx context.Context, req *product.UpdateProductStatusRequest) (*product.UpdateProductStatusResponse, error) {
	fmt.Println("Product Service :  UpdateProductStatus")
	fmt.Println("Product Service :  UpdateProductStatus - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return &product.UpdateProductStatusResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	err = s.Repo.UpdateProductStatus(ctx, objID, req.ProductStatus)
	if err != nil {
		fmt.Println("Product Service :  UpdateProductStatus - ERROR")
		fmt.Println(err.Error())

		return &product.UpdateProductStatusResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("Product Service :  UpdateProductStatus - SUCCESS")
	return &product.UpdateProductStatusResponse{
		Status: http.StatusNoContent,
	}, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, req *product.DeleteProductRequest) (*product.DeleteProductResponse, error) {
	fmt.Println("Product Service :  DeleteProduct")
	fmt.Println("Product Service :  DeleteProduct - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return &product.DeleteProductResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	err = s.Repo.DeleteProduct(ctx, objID)
	if err != nil {
		fmt.Println("Product Service :  DeleteProduct - ERROR")
		fmt.Println(err.Error())
		return &product.DeleteProductResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("Product Service :  DeleteProduct - SUCCESS")
	return &product.DeleteProductResponse{
		Status: http.StatusNoContent,
	}, nil
}
