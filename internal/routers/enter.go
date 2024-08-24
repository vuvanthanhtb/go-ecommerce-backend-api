package routers

import (
	"github.com/vuvanthanhtb/go-ecommerce-backend-api/internal/routers/admin"
	"github.com/vuvanthanhtb/go-ecommerce-backend-api/internal/routers/user"
)

type RouterGroup struct {
	User   user.UserRouterGroup
	Manage admin.AdminRouterGroup
}

var RouterGroupApp = new(RouterGroup)
