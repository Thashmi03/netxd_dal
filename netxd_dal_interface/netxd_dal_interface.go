package netxddalinterface

import netxddalmodels "github.com/Thashmi03/netxd_dal/netxd_dal_models"

type ICustomer interface{
	CreateCustomer(detail * netxddalmodels.Customer)(*netxddalmodels.DbResponse,error)
}

