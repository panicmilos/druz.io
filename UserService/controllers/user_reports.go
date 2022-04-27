package controllers

import (
	"UserService/api_contracts"
	"UserService/errors"
	"UserService/helpers"
	"UserService/models"
	"UserService/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sarulabs/di"
)

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
	reportedBy, _ := strconv.Atoi(r.Header.Get("name"))

	report := &models.UserReport{
		Reported:   uint(id),
		ReportedBy: uint(reportedBy),
		Reason:     request.Reason,
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
