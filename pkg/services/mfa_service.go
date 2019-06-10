package services

import (
	"encoding/json"
	"fmt"
	"github.com/ProtocolONE/qilin-store-api/pkg/api/dto"
	"github.com/ProtocolONE/qilin-store-api/pkg/common"
	"github.com/ProtocolONE/qilin-store-api/pkg/interfaces"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"time"
)

type endpoint struct {
	listURL   string
	removeURL string
	addURL    string
}

type mfaService struct {
	endpoint endpoint
	storage  *redis.Client
}

var cacheTime time.Duration

func NewMfaService(issuer string, storage *redis.Client) interfaces.MfaService {
	cacheTime, _ = time.ParseDuration("5m")
	conf := endpoint{
		listURL:   issuer + "/mfa/list",
		addURL:    issuer + "/mfa/add",
		removeURL: issuer + "/mfa/remove",
	}
	return &mfaService{storage: storage, endpoint: conf}
}

func (mfaService) Add(userId string, providerId string) error {
	panic("implement me")
}

func (service *mfaService) List(userId string) ([]dto.MfaProviderDTO, error) {
	storageId := fmt.Sprintf("mfa_list:%s", userId)
	getCmd := service.storage.Get(storageId)
	if err := getCmd.Err(); err != nil {
		return nil, common.NewServiceError(http.StatusInternalServerError, err)
	}

	var result []dto.MfaProviderDTO
	err := json.Unmarshal([]byte(getCmd.String()), &result)
	if err == nil {
		return result, nil
	}

	formData := url.Values{
		"client_id": {userId},
	}

	resp, err := http.PostForm(service.endpoint.listURL, formData)
	defer resp.Body.Close()

	if err != nil {
		return nil, common.NewServiceError(http.StatusInternalServerError, err)
	}

	if code := resp.StatusCode; code < 200 || code > 299 {
		return nil, common.NewServiceErrorf(http.StatusBadRequest, "Server return code: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, common.NewServiceError(http.StatusBadRequest, errors.Wrap(err, "Can't decode json"))
	}

	storageBytes, err := json.Marshal(result)

	if err := service.storage.Set(storageId, string(storageBytes), cacheTime).Err(); err != nil {
		zap.L().Error(fmt.Sprintf("Can't set `%s`", storageId), zap.Error(err))
	}

	return result, nil
}

func (service *mfaService) Remove(userId string, providerId string) error {
	formData := url.Values{
		"client_id": {userId},
		"provider_id": {providerId},
	}

	resp, err := http.PostForm(service.endpoint.removeURL, formData)
	defer resp.Body.Close()

	if err != nil {
		return common.NewServiceError(http.StatusInternalServerError, err)
	}

	if code := resp.StatusCode; code < 200 || code > 299 {
		return common.NewServiceErrorf(http.StatusBadRequest, "Server return code: %d", resp.StatusCode)
	}

	if err := service.storage.Del(fmt.Sprintf("mfa_list:%s", userId)).Err(); err != nil {
		zap.L().Error(fmt.Sprintf("Can't delete `mfa_list:%s` from storage", userId), zap.Error(err))
	}

	return nil
}
