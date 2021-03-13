package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/superterro666/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* BuscoPerfil */
func BuscoPerfil(ID string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoCN.Database("twitter")

	col := db.Collection("usuarios")

	var perfil models.Usuario

	objID, _ := primitive.ObjectIDFromHex(ID)

	fmt.Println(objID)

	condicion := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condicion).Decode(&perfil)

	perfil.Password = ""

	if err != nil {
		fmt.Println("No se encontro el usuairo" + err.Error())
		return perfil, err
	}

	return perfil, nil

}
