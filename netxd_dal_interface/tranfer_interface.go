package netxddalinterface

import model "github.com/Thashmi03/transfer_model"
type Itransact interface{
	Transfer(detail *model.Transaction)(string,error)
}