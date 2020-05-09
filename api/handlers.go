package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"io"
)




func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	io.WriteString(w, "Create User Hander")
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}