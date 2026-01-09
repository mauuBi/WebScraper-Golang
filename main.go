package main

import (
	"context"
	"time"

	handlegoquery "github.com/mauuBi/NewScraperGo/handleGoQuery"
	handlerequest "github.com/mauuBi/NewScraperGo/handleRequest"
)


func main(){
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	resp := handlerequest.CreateRequest(&ctx)
	defer resp.Body.Close()
	handlegoquery.GoQueryParsing(resp)
}