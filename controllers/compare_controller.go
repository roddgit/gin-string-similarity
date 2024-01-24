package controllers

import (
	"fmt"
	"gin-string-similarity/configs"
	"gin-string-similarity/payloads"
	_ "io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	// jaro "github.com/toldjuuso/go-jaro-winkler-distance"
)

var start_process = configs.StartProcess()
var validateReq = validator.New()
var logs = configs.ZeroLogger()

func CompareHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var compare payloads.CompareRequest
		var logs_id = configs.LogsIdGenerator("GIN")

		// validate the request body
		if err := c.BindJSON(&compare); err != nil {
			errs := payloads.CompareFault{
				StatusCode:    "11",
				StatusMessage: "failed",
				Message:       err.Error(),
			}
			logs.Error().Str("Request", configs.ArrayToString(errs)).Msg(logs_id)

			c.JSON(http.StatusBadRequest, errs)
			return
		}

		// use validator lib to validate requires filds
		if validationErr := validateReq.Struct(&compare); validationErr != nil {
			errs := payloads.CompareFault{
				StatusCode:    "22",
				StatusMessage: "failed",
				Message:       validationErr.Error(),
			}
			logs.Error().Str("Request", configs.ArrayToString(errs)).Msg(logs_id)

			c.JSON(http.StatusBadRequest, errs)
			return
		}

		logs.Info().Str("Request", configs.ArrayToString(compare)).Msg(logs_id)

		name_pmo_raw := compare.NamePMO
		name_core_raw := compare.NameCORE

		name_pmo_clean := configs.ValidationString(name_pmo_raw)
		name_core_clean := configs.ValidationString(name_core_raw)

		cleansing := map[string]string{"name_pmo": name_pmo_clean, "name_core": name_core_clean}

		logs.Info().Str("Processing", configs.ArrayToString(cleansing)).Msg(logs_id)

		// compare string by jaro
		// distance := jaro.Calculate(name_pmo_clean, name_core_clean)
		distance := configs.JaroWinklerDistance(name_pmo_clean, name_core_clean)

		// simplified decimal
		result_jaro := configs.TrimFloat(distance)

		result := payloads.CompareResponse{
			StatusCode:           "00",
			StatusMessage:        "success",
			LogsId:               logs_id,
			NameMatchingTreshold: result_jaro,
		}

		if configs.EnvConfig()["DB_METHOD"] == "MONGO" {
			// save to mongo db
			row_count := configs.InsertTresholdMongo(logs_id, name_pmo_raw, name_core_raw, name_pmo_clean, name_core_clean, result_jaro)
			logs.Info().Str("Save result to db", fmt.Sprint(row_count)).Msg(logs_id)
		} else if configs.EnvConfig()["DB_METHOD"] == "ORACLE" {
			// save to oracle db
			row_count := configs.InsertTresholdORA(logs_id, name_pmo_raw, name_core_raw, name_pmo_clean, name_core_clean, result_jaro)
			logs.Info().Str("Save result to db", fmt.Sprint(row_count)).Msg(logs_id)
		}

		logs.Info().Str("Response", configs.ArrayToString(result)).Msg(logs_id)

		c.JSON(http.StatusOK, result)
		logs.Info().Str("Process finished elapsed time", configs.EndProcess(start_process)+" ms").Msg(logs_id)

	}
}
