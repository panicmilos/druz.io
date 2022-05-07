package controllers

import (
	"net/http"
	"strconv"

	"github.com/panicmilos/druz.io/UserService/api_contracts"
	"github.com/panicmilos/druz.io/UserService/dto"
	"github.com/panicmilos/druz.io/UserService/errors"
	"github.com/panicmilos/druz.io/UserService/helpers"
	"github.com/panicmilos/druz.io/UserService/models"
	"github.com/panicmilos/druz.io/UserService/services"

	"github.com/gorilla/mux"
	"github.com/sarulabs/di"
)

var SearchReports = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	params := &dto.UserReportsSearchParams{
		Reported:   r.URL.Query().Get("Reported"),
		ReportedBy: r.URL.Query().Get("ReportedBy"),
		Reason:     r.URL.Query().Get("Reason"),
	}

	userReportsService := di.Get(r, services.UserReportService).(*services.UserReportsService)

	helpers.JSONResponse(w, 200, userReportsService.Search(params))
})

var ReportUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var request *api_contracts.ReportUserRequest
	err := helpers.ReadJSONBody(r, &request)

	if errors.Handle(err, w) {
		return
	}

	err = request.Validate()
	if errors.Handle(err, w) {
		return
	}

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	reportedById, _ := strconv.Atoi(r.Header.Get("name"))

	report := &models.UserReport{
		ReportedId:   uint(id),
		ReportedById: uint(reportedById),
		Reason:       request.Reason,
	}

	userReportsService := di.Get(r, services.UserReportService).(*services.UserReportsService)
	createdReport, err := userReportsService.Create(report)
	if errors.Handle(err, w) {
		return
	}

	helpers.JSONResponse(w, 200, createdReport)
})

var IgnoreReport = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	userReportsService := di.Get(r, services.UserReportService).(*services.UserReportsService)
	ignoredReport, err := userReportsService.Delete(uint(id))
	if errors.Handle(err, w) {
		return
	}

	helpers.JSONResponse(w, 200, ignoredReport)
})
