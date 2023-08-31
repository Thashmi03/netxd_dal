package netxddalservices

import (
	
	"context"
	"fmt"
	"log"
	model "github.com/Thashmi03/transfer_model"
	tinterface "github.com/Thashmi03/transfer_interface"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


type Transaction struct{
	ctx context.Context
	client *mongo.Client
	mongoCollection *mongo.Collection
	Customercollection *mongo.Collection
}

func InitTransaction (ccollection *mongo.Collection,tcollection *mongo.Collection,ctx context.Context,client *mongo.Client)(tinterface.Itransact){
	return &Transaction{ctx,client,tcollection,ccollection}
}

func (t*Transaction)Transfer(detail *model.Transaction)(string,error){
	session,err:=t.client.StartSession()
	if err!=nil{
		log.Fatal(err)
	}
	defer session.EndSession(context.Background())
	_,err=session.WithTransaction(context.Background(),func(ctx mongo.SessionContext) (interface{}, error){
		_, err := t.Customercollection.UpdateOne(ctx,
			bson.M{"customer_id": detail.From_account},
			bson.M{"$inc": bson.M{"balance": -detail.Amount}})
		if err!=nil{
			fmt.Println("failed1")
		}
		_,err2:=t.Customercollection.UpdateOne(context.Background(),
		bson.M{"customer_id":detail.To_account},
		bson.M{"$inc":bson.M{"balance":detail.Amount}},)
		if err2!=nil{
			fmt.Println("failed2")
		}
		trans:=&model.Transaction{
			Transaction_id:"T001",
			From_account:detail.From_account,
			To_account :detail.To_account,
			Amount :detail.Amount,
		}
		res,err:=t.mongoCollection.InsertOne(context.Background(),&trans)
		if err!=nil{
			return "nil",err
		}
		
	var newUser *model.TResponse
	query := bson.M{"_id": res.InsertedID}
	
	err3 := t.mongoCollection.FindOne(t.ctx, query).Decode(&newUser)
	if err3 != nil {
	return nil, err3
	}
	return newUser, nil
	})
	if err != nil {
	return "failed", err
	}

	return "yes",nil
}