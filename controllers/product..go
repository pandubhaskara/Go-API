package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test3/models"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProductController struct{
	session *mgo.Session
}

func NewProductController(s *mgo.Session) *ProductController{
	return &ProductController{s}
}

func (uc ProductController) GetProduct (w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id := p.ByName("id")

	if *bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound)
	}
	oid := bson.ObjectHex(id)

	u:= models.Product{}

	if err := uc.Session.DB("mongo-golang").C("product").findId(oid).One(&u); err != nil{
		w.WriteHeader(404)
		return
	}
	json.Marshal(u)

	uj, err := json.Marshal(u)
	if err != nil{
		fmt.Println((err))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Println(w, "%s\n", uj)
}

func (uc ProductController) PostProduct (w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	u := models.Product{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()

	uc.session.DB("mongo-golang").C("product").Insert(u)

	uj, err := json.Marshal(u)

	if err != nil{
		 fmt.Println(err)
	 }

	 w.Header().Set("Content-Type", "application/json")
	 w.WriteHeader(http.StatusCreated)
	 fmt.Fprintf(w, "%s\n", uj)
}

func (uc ProductController) DeleteProduct (w httprouter, r *http.Request, p httprouter.Params){
	id := p.Byname("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(404)
		return
	}

	oid:= bson.ObjectIdHex(id)

	if err := uc.session.DB("mongo-golang").C("product").Remove(u){
		w.WriteHeader(404)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted product", old, "\n")


}