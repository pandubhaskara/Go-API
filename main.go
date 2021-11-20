package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

// func close(client *mongo.Client, ctx context.Context,
// 	cancel context.CancelFunc) {
// 	defer cancel()
// 	defer func() {
// 		if err := client.Disconnect(ctx); err != nil {
// 			panic(err)
// 		}
// 	}()
// }

// func connect(uri string) (*mongo.Client, context.Context,
// 	context.CancelFunc, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
// 	return client, ctx, cancel, err
// }

// func ping(client *mongo.Client, ctx context.Context) error {
// 	if err := client.Ping(ctx, readpref.Primary()); err != nil {
// 		return err
// 	}
// 	fmt.Println("connected successfully")
// 	return nil
// }

func main() {
	r := httprouter.New()
	uc := controllers.NewProductController(getSession())
	r.GET("/product", uc.GetProduct)
	r.POST("/product", uc.PostProduct)
	r.DELETE("/product/:id", uc.DeleteProduct)
	http.ListenAndServe("localhost:5000", r)

	// client, ctx, cancel, err := connect("mongodb://localhost:27017")
	// if err != nil {
	// 	panic(err)
	// }
	// defer close(client, ctx, cancel)
	// ping(client, ctx)
}
func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return s
}
