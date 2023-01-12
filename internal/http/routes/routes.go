package routes

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(
	items ItemsRoutes,
) Routes {
	return Routes{
		items,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
