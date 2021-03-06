/*
 * Copyright 2018 Intel Corporation, Inc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/onap/multicloud-k8s/src/k8splugin/internal/rb"

	"github.com/gorilla/mux"
)

// Used to store backend implementations objects
// Also simplifies mocking for unit testing purposes
type rbDefinitionHandler struct {
	// Interface that implements bundle Definition operations
	// We will set this variable with a mock interface for testing
	client rb.DefinitionManager
}

// createOrUpdateHandler handles creation of the definition entry in the database
func (h rbDefinitionHandler) createHandler(w http.ResponseWriter, r *http.Request) {
	var v rb.Definition

	err := json.NewDecoder(r.Body).Decode(&v)
	switch {
	case err == io.EOF:
		http.Error(w, "Empty body", http.StatusBadRequest)
		return
	case err != nil:
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	h.createOrUpdateHandler(v, w, false)
}

// createOrUpdateHandler handles creation of the definition entry in the database
func (h rbDefinitionHandler) updateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["rbname"]
	version := vars["rbversion"]

	var v rb.Definition

	err := json.NewDecoder(r.Body).Decode(&v)
	switch {
	case err == io.EOF:
		http.Error(w, "Empty body", http.StatusBadRequest)
		return
	case err != nil:
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if v.RBVersion != "" && v.RBVersion != version {
		http.Error(w, "RB version mismatch", http.StatusBadRequest)
		return
	}

	if v.RBName != "" && v.RBName != name {
		http.Error(w, "RB name mismatch", http.StatusBadRequest)
		return
	}

	v.RBVersion = version
	v.RBName = name

	h.createOrUpdateHandler(v, w, true)
}

// createOrUpdateHandler handles creation of the definition entry in the database
func (h rbDefinitionHandler) createOrUpdateHandler(v rb.Definition, w http.ResponseWriter, update bool) {
	// Name is required.
	if v.RBName == "" {
		http.Error(w, "Missing name in request", http.StatusBadRequest)
		return
	}

	// Version is required.
	if v.RBVersion == "" {
		http.Error(w, "Missing version in request", http.StatusBadRequest)
		return
	}

	var ret rb.Definition
	var err error
	if update {
		ret, err = h.client.Update(v)
	} else {
		ret, err = h.client.Create(v)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(ret)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// uploadHandler handles upload of the bundle tar file into the database
func (h rbDefinitionHandler) uploadHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["rbname"]
	version := vars["rbversion"]

	inpBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read body", http.StatusBadRequest)
		return
	}

	if len(inpBytes) == 0 {
		http.Error(w, "Empty body", http.StatusBadRequest)
		return
	}

	err = h.client.Upload(name, version, inpBytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// listVersionsHandler handles GET (list) operations on the endpoint
// Returns a list of rb.Definitions
func (h rbDefinitionHandler) listVersionsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["rbname"]

	ret, err := h.client.List(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(ret)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// listVersionsHandler handles GET (list) operations on the endpoint
// Returns a list of rb.Definitions
func (h rbDefinitionHandler) listAllHandler(w http.ResponseWriter, r *http.Request) {

	ret, err := h.client.List("")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(ret)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// getHandler handles GET operations on a particular ids
// Returns a rb.Definition
func (h rbDefinitionHandler) getHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["rbname"]
	version := vars["rbversion"]

	ret, err := h.client.Get(name, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(ret)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// deleteHandler handles DELETE operations on a particular bundle definition id
func (h rbDefinitionHandler) deleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["rbname"]
	version := vars["rbversion"]

	err := h.client.Delete(name, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
