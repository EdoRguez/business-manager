package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"github.com/EdoRguez/business-manager/gateway/pkg/product/contracts"
	"github.com/EdoRguez/business-manager/gateway/pkg/product/pb"
	"google.golang.org/grpc"
)

var productServiceClient pb.ProductServiceClient

func InitProductServiceClient(c *config.Config) error {
	fmt.Println("Product CLIENT :  InitProductServiceClient")
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.Product_Svc_Url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
		return err
	}

	productServiceClient = pb.NewProductServiceClient(cc)
	return nil
}

func CreateProduct(body contracts.CreateProductRequest, c context.Context) (*pb.CreateProductResponse, *contracts.Error) {
	fmt.Println("Proudct CLIENT :  CreateProduct")

	fmt.Println("Product CLIENT :  CreateProduct - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	createProductParams := &pb.CreateProductRequest{
		CompanyId:     body.CompanyId,
		Name:          body.Name,
		Description:   body.Description,
		Sku:           body.Sku,
		Quantity:      body.Quantity,
		Price:         body.Price,
		Images:        body.Images,
		ProductStatus: body.ProductStatus,
	}

	res, err := productServiceClient.CreateProduct(c, createProductParams)

	if err != nil {
		fmt.Println("Product CLIENT :  CreateProduct - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("API Gateway :  CreateProduct - SUCCESS")
	return res, nil
}

func GetProduct(id string, c context.Context) (*contracts.GetProductResponse, *contracts.Error) {
	fmt.Println("Product CLIENT :  GetProduct")

	params := &pb.GetProductRequest{
		Id: id,
	}

	res, err := productServiceClient.GetProduct(c, params)

	fmt.Println("product clinet")
	fmt.Println(res)

	if err != nil {
		fmt.Println("Product CLIENT :  GetProduct - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	if res.Status != http.StatusOK {
		error := &contracts.Error{
			Status: res.Status,
			Error:  res.Error,
		}
		return nil, error
	}

	fmt.Println("Product CLIENT :  GetProduct - SUCCESS")
	return &contracts.GetProductResponse{
		Id:            res.Id,
		CompanyId:     res.CompanyId,
		Name:          res.Name,
		Description:   res.Description,
		Sku:           res.Sku,
		Quantity:      res.Quantity,
		Price:         res.Price,
		Images:        res.Images,
		ProductStatus: res.ProductStatus,
	}, nil
}

func GetProducts(params *pb.GetProductsRequest, c context.Context) ([]*contracts.GetProductResponse, *contracts.Error) {
	fmt.Println("Product CLIENT :  GetProduct")

	if params.CompanyId <= 0 {
		error := &contracts.Error{
			Status: http.StatusBadRequest,
			Error:  "Company ID is required in order to get results",
		}

		return nil, error
	}

	res, err := productServiceClient.GetProducts(c, params)

	if err != nil {
		fmt.Println("Product CLIENT :  GetProducts - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	if res.Status != http.StatusOK {
		error := &contracts.Error{
			Status: res.Status,
			Error:  res.Error,
		}

		return nil, error
	}

	pr := make([]*contracts.GetProductResponse, 0, len(res.Products))
	for _, v := range res.Products {
		pr = append(pr, &contracts.GetProductResponse{
			Id:            v.Id,
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

	fmt.Println("Product CLIENT :  GetProducts - SUCCESS")
	return pr, nil
}

func UpdateProduct(id string, body contracts.UpdateProductRequest, c context.Context) (*pb.UpdateProductResponse, *contracts.Error) {
	fmt.Println("Product CLIENT :  UpdateProduct")

	fmt.Println("Product CLIENT :  UpdateProduct - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	updateProductParams := &pb.UpdateProductRequest{
		Id:            id,
		Name:          body.Name,
		Description:   body.Description,
		Sku:           body.Sku,
		Quantity:      body.Quantity,
		Price:         body.Price,
		Images:        body.Images,
		ProductStatus: body.ProductStatus,
	}

	res, err := productServiceClient.UpdateProduct(c, updateProductParams)

	if err != nil {
		fmt.Println("Product CLIENT :  UpdateProduct - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Product CLIENT :  UpdateProduct - SUCCESS")
	return res, nil
}

func DeleteProduct(id string, c context.Context) (*pb.DeleteProductResponse, *contracts.Error) {
	fmt.Println("Product CLIENT :  DeleteProduct")

	params := &pb.DeleteProductRequest{
		Id: id,
	}

	res, err := productServiceClient.DeleteProduct(c, params)

	if err != nil {
		fmt.Println("Product CLIENT :  DeleteProduct - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Product CLIENT :  DeleteProduct - SUCCESS")
	return res, nil
}
