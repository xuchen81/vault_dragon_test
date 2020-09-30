package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"vault_dragon_test/repositories"

	"github.com/gin-gonic/gin"

	"vault_dragon_test/models"
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
	for key, value := range KVmapping {
		values := []models.Value{
			{
				Value: value,
			},
		}
		k := models.Key{
			Key:    key,
			Values: values,
		}
		res = repositories.KeyRepo.AddKey(&k)
	}

	if res.Error != nil {
		utils.JSONResponse(http.StatusOK, gin.H{
			"error": res.Error.Error(),
		}, w)
		return
	}
	createdKey := res.Value.(*models.Key)

	utils.JSONResponse(http.StatusOK, gin.H{
		"key":       createdKey.Key,
		"value":     createdKey.Values[0].Value,
		"timestamp": createdKey.CreatedAt.Unix(),
	}, w)
}
