package main

import (
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type envelope map[string]any

func readInt64Param(paramKey string, r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	i, err := strconv.ParseInt(params.ByName(paramKey), 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func (app *application) readIDParam(r *http.Request) (int64, error) {
	id, err := readInt64Param("id", r)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
