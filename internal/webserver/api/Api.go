package api

import (
	"Gokapi/internal/configuration"
	"Gokapi/internal/helper"
	"Gokapi/internal/models"
	"Gokapi/internal/storage"
	"Gokapi/internal/webserver/sessionmanager"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

// Process parses the request and executes the API call or returns an error message to the sender
func Process(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("cache-control", "no-store")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	request := parseRequest(r)
	if !isAuthorised(w, request) {
		return
	}
	switch request.requestUrl {
	case "/files/list":
		list(w)
	case "/files/add":
		upload(w, request)
	case "/files/delete":
		deleteFile(w, request)
	case "/auth/friendlyname":
		changeFriendlyName(w, request)
	default:
		sendError(w, http.StatusBadRequest, "Invalid request")
	}
}

func DeleteApiKey(id string) bool {
	if !isValidApiKey(id, false) {
		return false
	}
	settings := configuration.GetServerSettings()
	delete(settings.ApiKeys, id)
	configuration.ReleaseAndSave()
	return true
}

func NewApiKey(name string) {
	if name == "" {
		name = "Unnamed key"
	}
	settings := configuration.GetServerSettings()
	newKey := models.ApiKey{
		Id:           helper.GenerateRandomString(30),
		FriendlyName: name,
		LastUsed:     0,
	}
	settings.ApiKeys[newKey.Id] = newKey
	configuration.ReleaseAndSave()
}

func changeFriendlyName(w http.ResponseWriter, request apiRequest) {
	if !isValidApiKey(request.apiKeyToModify, false) {
		sendError(w, http.StatusBadRequest, "Invalid api key provided.")
		return
	}
	if request.friendlyName == "" {
		request.friendlyName = "Unnamed key"
	}
	settings := configuration.GetServerSettings()
	key := settings.ApiKeys[request.apiKeyToModify]
	if key.FriendlyName != request.friendlyName {
		key.FriendlyName = request.friendlyName
		settings.ApiKeys[request.apiKeyToModify] = key
		configuration.ReleaseAndSave()
	} else {
		configuration.Release()
	}
	sendOk(w)
}

func deleteFile(w http.ResponseWriter, request apiRequest) {
	ok := storage.DeleteFile(request.headerId)
	if ok {
		sendOk(w)
	} else {
		sendError(w, http.StatusBadRequest, "Invalid id provided.")
	}
}

func list(w http.ResponseWriter) {
	sendOk(w)
	settings := configuration.GetServerSettings()
	result, err := json.Marshal(settings.Files)
	configuration.Release()
	helper.Check(err)
	_, _ = w.Write(result)
}

func upload(w http.ResponseWriter, request apiRequest) {
	sendOk(w)
	// TODO
}

func isValidApiKey(key string, modifyTime bool) bool {
	if key == "" {
		return false
	}
	settings := configuration.GetServerSettings()
	defer func() {
		configuration.Release()
	}()
	savedKey, ok := settings.ApiKeys[key]
	if ok && savedKey.Id != "" {
		if modifyTime {
			savedKey.LastUsed = time.Now().Unix()
			settings.ApiKeys[key] = savedKey
		}
		return true
	}
	return false
}

func isAuthorised(w http.ResponseWriter, request apiRequest) bool {
	if isValidApiKey(request.apiKey, true) || sessionmanager.IsValidSession(w, request.request) {
		return true
	}
	sendError(w, http.StatusUnauthorized, "Unauthorized")
	return false
}

func sendError(w http.ResponseWriter, errorInt int, errorMessage string) {
	w.WriteHeader(errorInt)
	_, _ = w.Write([]byte("{\"Result\":\"error\",\"ErrorMessage\":\"" + errorMessage + "\"}"))
}

func sendOk(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

type apiRequest struct {
	apiKey         string
	requestUrl     string
	headerId       string
	friendlyName   string
	apiKeyToModify string
	request        *http.Request
}

func parseRequest(r *http.Request) apiRequest {
	return apiRequest{
		apiKey:         r.Header.Get("apikey"),
		headerId:       r.Header.Get("id"),
		friendlyName:   r.Header.Get("friendlyName"),
		apiKeyToModify: r.Header.Get("apiKeyToModify"),
		requestUrl:     strings.Replace(r.URL.String(), "/api", "", 1),
		request:        r,
	}
}
