package api

import (
	"github.com/Victoria-engine/api-v2/pkg/api/auth"
	authTransport "github.com/Victoria-engine/api-v2/pkg/api/auth/transport"
	"github.com/Victoria-engine/api-v2/pkg/utl/jwtauth"
	"github.com/Victoria-engine/api-v2/pkg/utl/seed"
	"time"
	//am "github.com/Victoria-engine/api-v2/pkg/utl/middleware/authmiddleware"
	"github.com/Victoria-engine/api-v2/pkg/utl/server"
	"os"
)

// Start : Initializes the server structure
func Start(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName, port string) error {
	s := server.New(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName)

	//authMiddleware = am.SetMiddlewareAuthentication()
	//TODO: Add JSON middleware for response

	expTime := time.Hour * 168
	jwt, err := jwtauth.New("HS256", os.Getenv("JWT_SECRET"), expTime, 64)
	if err != nil {
		return err
	}

	authTransport.NewHTTP(auth.Initialize(s.DB, jwt), s.Router)

	// Seed some dummy data for testing
	seed.Load(s.DB)

	s.Run(port)

	return nil
}
