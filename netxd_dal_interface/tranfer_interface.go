package netxddalinterface

import netxddalmodels "github.com/Thashmi03/netxd_dal/netxd_dal_models"

type Itransact interface{
	Transfer(detail *netxddalmodels.Transaction)(string,error)
}