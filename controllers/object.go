package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	"vault_dragon_test/models"
	"vault_dragon_test/repositories"
	"vault_dragon_test/utils"
)

// CreateKeyValueHandler : create a key value pair in DB
func CreateKeyValueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	requestBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.JSONResponse(http.StatusInternalServerError, gin.H{"error": err.Error()}, w)
		return
	}
	var KVmapping map[string]string
	if e := json.Unmarshal(requestBytes, &KVmapping); e != nil {
		utils.JSONResponse(http.StatusInternalServerError, gin.H{"error": e.Error()}, w)
		return
	}

	var res repositories.RepositoryResult
	var key, value string
	for k, v := range KVmapping {
		key = k
		value = v
	}

	check := repositories.KeyRepo.GetKeyByKey(key)
	if check.Error != nil && !errors.Is(check.Error, gorm.ErrRecordNotFound) {
		utils.JSONResponse(http.StatusOK, gin.H{
			"error": check.Error.Error(),
		}, w)
		return
	}
	// it the key is found, add a new value
	if check.RowsAffected == 1 {
		found := check.Value.(*models.Key)
		res := repositories.ValueRepo.AddValue(&models.Value{
			Name:  value,
			KeyID: found.ID,
		})

		if res.Error != nil {
			utils.JSONResponse(http.StatusOK, gin.H{
				"error": check.Error.Error(),
			}, w)
			return
		}

		newValue := res.Value.(*models.Value)

		utils.JSONResponse(http.StatusOK, gin.H{
			"key":       key,
			"value":     newValue.Name,
			"timestamp": newValue.CreatedAt.Unix(),
		}, w)
		return
	}

	// if the key does not exist, create it
	values := []models.Value{
		{
			Name: value,
		},
	}
	k := models.Key{
		Name:   key,
		Values: values,
	}
	res = repositories.KeyRepo.AddKey(&k)

	if res.Error != nil {
		utils.JSONResponse(http.StatusOK, gin.H{
			"error": res.Error.Error(),
		}, w)
		return
	}
	createdKey := res.Value.(*models.Key)

	utils.JSONResponse(http.StatusOK, gin.H{
		"key":       createdKey.Name,
		"value":     createdKey.Values[0].Name,
		"timestamp": createdKey.CreatedAt.Unix(),
	}, w)
}

func GetKeyValueHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	keyName := params["keyname"]
	timestamp := r.URL.Query().Get("timestamp")

	var timeBy time.Time
	// if timestamp is not specified, use current time
	if timestamp == "" {
		timeBy = time.Now().UTC()
	} else {
		timestampInt, _ := strconv.ParseInt(timestamp, 10, 64)
		timeBy = time.Unix(timestampInt, 0)
	}

	res := repositories.ValueRepo.GetLatestValueByKeyBefore(keyName, timeBy)
	if res.Error != nil {
		utils.JSONResponse(http.StatusOK, gin.H{
			"error": res.Error.Error(),
		}, w)
		return
	}
	v := res.Value.(*models.Value)
	utils.JSONResponse(http.StatusOK, gin.H{
		"value": v.Name,
	}, w)
}
