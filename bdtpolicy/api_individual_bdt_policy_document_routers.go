// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

/*
 * Npcf_BDTPolicyControl Service API
 *
 * The Npcf_BDTPolicyControl Service is used by an NF service consumer to retrieve background data transfer policies from the PCF and to update the PCF with the background data transfer policy selected by the NF service consumer.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package bdtpolicy

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omec-project/openapi"
	"github.com/omec-project/openapi/models"
	"github.com/omec-project/pcf/logger"
	"github.com/omec-project/pcf/producer"
	"github.com/omec-project/util/httpwrapper"
)

// GetBDTPolicy - Read an Individual BDT policy
func HTTPGetBDTPolicy(c *gin.Context) {
	req := httpwrapper.NewRequest(c.Request, nil)
	req.Params["bdtPolicyId"] = c.Params.ByName("bdtPolicyId")

	rsp := producer.HandleGetBDTPolicyContextRequest(req)

	responseBody, err := openapi.Serialize(rsp.Body, "application/json")
	if err != nil {
		logger.Bdtpolicylog.Errorln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(rsp.Status, "application/json", responseBody)
	}
}

// UpdateBDTPolicy - Update an Individual BDT policy
func HTTPUpdateBDTPolicy(c *gin.Context) {
	var bdtPolicyDataPatch models.BdtPolicyDataPatch
	// step 1: retrieve http request body
	requestBody, err := c.GetRawData()
	if err != nil {
		problemDetail := models.ProblemDetails{
			Title:  "System failure",
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
			Cause:  "SYSTEM_FAILURE",
		}
		logger.Bdtpolicylog.Errorf("Get Request Body error: %+v", err)
		c.JSON(http.StatusInternalServerError, problemDetail)
		return
	}

	// step 2: convert requestBody to openapi models
	err = openapi.Deserialize(&bdtPolicyDataPatch, requestBody, "application/json")
	if err != nil {
		problemDetail := "[Request Body] " + err.Error()
		rsp := models.ProblemDetails{
			Title:  "Malformed request syntax",
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		logger.Bdtpolicylog.Errorln(problemDetail)
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	req := httpwrapper.NewRequest(c.Request, bdtPolicyDataPatch)
	req.Params["bdtPolicyId"] = c.Params.ByName("bdtPolicyId")

	rsp := producer.HandleUpdateBDTPolicyContextProcedure(req)

	responseBody, err := openapi.Serialize(rsp.Body, "application/json")
	if err != nil {
		logger.Bdtpolicylog.Errorln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(rsp.Status, "application/json", responseBody)
	}
}
