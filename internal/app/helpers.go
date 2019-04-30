package app

import "net/http"

func (a *App) internalServerError(w http.ResponseWriter, err error) {
	// TODO: log, use zap.AddCallerSkip(1)
	code := http.StatusInternalServerError
	if a.c.Debug {
		http.Error(w, err.Error(), code)
	} else {
		http.Error(w, http.StatusText(code), code)
	}
}

func (a *App) badRequest(w http.ResponseWriter, reason string) {
	// TODO: log, use zap.AddCallerSkip(1)
	code := http.StatusBadRequest
	if a.c.Debug {
		http.Error(w, reason, code)
	} else {
		http.Error(w, http.StatusText(code), code)
	}
}
