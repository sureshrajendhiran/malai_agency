package controller

import (
	"database/sql"
	"malai_agency/backend/env"
	controller "malai_agency/backend/micro-servers/admin-management/controller"
	"malai_agency/backend/middlewares"
	"malai_agency/backend/services/db"

	"github.com/gorilla/mux"
)

var DB *sql.DB
var BaseURL = ""

// Init func used to init all router points
func Init() *mux.Router {

	BaseURL = env.AdminManagementURL

	router := mux.NewRouter()
	db.Init()

	// Middleware func calls
	router.Use(middlewares.LoggingMiddleware)

	// User profiles based
	router.HandleFunc(BaseURL+"login/", controller.LoginCtrl).Methods("POST")

	router.HandleFunc(BaseURL+"create_record/{table_name}/", controller.CreateNewRowCtrl).Methods("POST")
	router.HandleFunc(BaseURL+"delete/{table_name}/{id}/", controller.DeleteCtrl).Methods("DELETE")
	router.HandleFunc(BaseURL+"update/{table_name}/", controller.UpdateCtrl).Methods("PUT")

	router.HandleFunc(BaseURL+"get_qi/{type}", controller.GetInvoiceCtrl).Methods("POST")
	router.HandleFunc(BaseURL+"update_qi/{type}/{operation}/", controller.UpdateQICtrl).Methods("PUT")
	router.HandleFunc(BaseURL+"get_invoice/{id}", controller.GetInvoiceByIdCtrl).Methods("GET")
	router.HandleFunc(BaseURL+"qi/{type}/{operation}/{id}", controller.GetInvoiceByIdCtrl).Methods("GET")
	router.HandleFunc(BaseURL+"search_option/{type}/", controller.SearchOptionCtrl).Methods("GET")
	router.HandleFunc(BaseURL+"get_filter_count/{type}/", controller.FilterCountCtrl).Methods("GET")
	router.HandleFunc(BaseURL+"get_master/{type}/", controller.GetMasterCtrl).Methods("POST")

	return router
}
