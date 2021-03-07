package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* Varible de conexion a mongo  */
var MongoCN = ConectarBD()

/* URL Mongo Atlas */
var clientOptions = options.Client().ApplyURI("mongodb+srv://root:terro1975@cluster0.dq8mv.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

/* ConectarBD nos permite conectar la bd */
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion exitosa")
	return client

}

/* Checqueo connection */
func ChequeoConnection() int {

	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
