package netxddalinterface

import tmodel "github.com/Thashmi03/transfer_model"
type Itransact interface{
	Transfer(detail * tmodel.Transaction)(string,error)
}