package route

import (
	account_controller "digitalwallet-service/src/webapp/controller/account"
	"net/http"
)

var accoutRoutes = []routeModel{
	{
		uri:          "/accounts",
		method:       http.MethodPost,
		function:     account_controller.Create,
		authRequired: false,
	},
	{
		uri:          "/accounts/{account-id}/cashin",
		method:       http.MethodPost,
		function:     account_controller.Cashin,
		authRequired: false,
	},
}
