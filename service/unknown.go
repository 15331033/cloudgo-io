package service

import (
    "net/http"
)

func unknownHandler() http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        http.Error(w, "501 page not implemented",501)
    }
}