package handlerequest

import (
	"context"
	"net/http"
	"time"

	errorhandling "github.com/mauuBi/NewScraperGo/errorHandling"
)

func CreateRequest(ctx *context.Context) *http.Response{
	url := "https://techcrunch.com/latest/"

	clicos := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequestWithContext(*ctx, http.MethodGet, url, nil);
	errorhandling.CheckError(err)

	resp, err := clicos.Do(req)
	errorhandling.CheckError(err)
	return resp
}




