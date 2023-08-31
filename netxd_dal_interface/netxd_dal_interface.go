package netxddalinterface

import netxddalmodels "banking_with_grpc/netxd_dal/netxd_dal_models"





type ICustomer interface{
	CreateCustomer(detail * netxddalmodels.Customer)(*netxddalmodels.DbResponse,error)
}