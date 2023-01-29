package main

import (
	"api-gateway/pkg/config"
	"api-gateway/pkg/policies"
	"api-gateway/pkg/routes"
	"api-gateway/pkg/utils"
	"fmt"

	fibercasbin "github.com/arsmn/fiber-casbin/v2"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/golang-jwt/jwt/v4"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3008",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	db := config.InitDB()
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize casbin adapter: %v", err))
	}

	enforcer, err := casbin.NewEnforcer("pkg/config/rbac_model.conf", adapter)
	if err != nil {
		panic(fmt.Sprintf("failed to create casbin enforcer: %v", err))
	}

	policies.SetupPolicies(enforcer)

	auth := fibercasbin.New(fibercasbin.Config{
		ModelFilePath: "pkg/config/rbac_model.conf",
		PolicyAdapter: adapter,
		Lookup: func(c *fiber.Ctx) string {
			headers := c.GetReqHeaders()
			tokenString := headers["Authorization"]
			if tokenString == "" {
				return ""
			}

			token, err := utils.ValidateToken(tokenString)
			if err != nil {
				return ""
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return ""
			}

			if token.Valid {
				return claims["email"].(string)
			}
			return ""
		},
	})

	// routes.SetupAdminRoutes(app, auth)
	// routes.SetupAdRoutes(app, auth)
	// routes.SetupApplicationRoutes(app, auth)
	routes.SetupEmployeeRoutes(app, auth)
	// routes.SetupEmployerRoutes(app, auth)
	// routes.SetupReviewRoutes(app, auth)

	app.Listen(":3007")
}
