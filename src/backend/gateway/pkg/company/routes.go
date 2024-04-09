package company

import (
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/company/routes"
	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"github.com/gorilla/mux"
)

type CompanyRoutes struct {
	config *config.Config
}

func LoadRoutes(router *mux.Router, c *config.Config) {
	loadCompanyRoutes(router, c)
	loadPaymentRoutes(router, c)
}

func loadCompanyRoutes(router *mux.Router, c *config.Config) {
	baseRoute := router.PathPrefix("/companies").Subrouter()

	cr := &CompanyRoutes{
		config: c,
	}

	mw := MiddlewareConfig{}

	getRouter := baseRoute.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/{id:[0-9]+}", cr.GetCompany)
	getRouter.HandleFunc("", cr.GetCompanies)

	postRouter := baseRoute.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("", cr.CreateCompany)
	postRouter.Use(mw.MiddlewareValidateCreateCompany)

	putRouter := baseRoute.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", cr.UpdateCompany)
	putRouter.Use(mw.MiddlewareValidateUpdateCompany)

	deleteRouter := baseRoute.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", cr.DeleteCompany)
}

func loadPaymentRoutes(router *mux.Router, c *config.Config) {
	baseRoute := router.PathPrefix("/payments").Subrouter()

	cr := &CompanyRoutes{
		config: c,
	}

	mw := MiddlewareConfig{}

	getRouter := baseRoute.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/{id:[0-9]+}", cr.GetPayment)
	getRouter.HandleFunc("", cr.GetPayments)

	postRouter := baseRoute.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("", cr.CreatePayment)
	postRouter.Use(mw.MiddlewareValidateCreatePayment)

	putRouter := baseRoute.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", cr.UpdatePayment)
	putRouter.Use(mw.MiddlewareValidateUpdatePayment)

	deleteRouter := baseRoute.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", cr.DeletePayment)
}

func (cr *CompanyRoutes) CreateCompany(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  CreateCompany Called --> 1")
	routes.CreateCompany(w, r, cr.config)
}

func (cr *CompanyRoutes) GetCompany(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  GetCompany Called --> 1")
	routes.GetCompany(w, r, cr.config)
}

func (cr *CompanyRoutes) GetCompanies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  GetCompanies Called --> 1")
	routes.GetCompanies(w, r, cr.config)
}

func (cr *CompanyRoutes) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  UpdateCompany Called --> 1")
	routes.UpdateCompany(w, r, cr.config)
}

func (cr *CompanyRoutes) DeleteCompany(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  DeleteCompany Called --> 1")
	routes.DeleteCompany(w, r, cr.config)
}

func (cr *CompanyRoutes) CreatePayment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  CreatePayment Called --> 1")
	routes.CreatePayment(w, r, cr.config)
}

func (cr *CompanyRoutes) GetPayment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  GetPayment Called --> 1")
	routes.GetPayment(w, r, cr.config)
}

func (cr *CompanyRoutes) GetPayments(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  GetPayments Called --> 1")
	routes.GetPayments(w, r, cr.config)
}

func (cr *CompanyRoutes) UpdatePayment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  UpdatePayment Called --> 1")
	routes.UpdatePayment(w, r, cr.config)
}

func (cr *CompanyRoutes) DeletePayment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Gateway :  DeletePayment Called --> 1")
	routes.DeletePayment(w, r, cr.config)
}
