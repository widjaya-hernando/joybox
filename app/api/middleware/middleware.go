package middleware

type Config struct {
	// put middleware config here
}

type Middleware struct {
	config      Config
	routeGroups map[string]string
}

func New(
	cfg Config,
) *Middleware {
	routeGroups := map[string]string{
		"admin":  "/admin",
		"user":   "",
		"public": "/pblic",
		"server": "/server",
	}

	return &Middleware{
		config:      cfg,
		routeGroups: routeGroups,
	}
}
