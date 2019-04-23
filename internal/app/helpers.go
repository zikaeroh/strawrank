package app

import "net/http"

func httpError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}

func (a *App) internalServerError(w http.ResponseWriter, err error) {
	code := http.StatusInternalServerError
	if a.c.Debug {
		http.Error(w, err.Error(), code)
	} else {
		http.Error(w, http.StatusText(code), code)
	}
}
