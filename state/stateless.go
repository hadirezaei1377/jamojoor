package state

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

type User2 struct {
	Username string
	Password string
}

// loginHandler for stateless authentication
func loginHandler2(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Authenticate user (this is just a simulation)
		if username != "user" || password != "password" {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// Create JWT token
		expirationTime := time.Now().Add(5 * time.Minute)
		claims := &jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Subject:   username,
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			http.Error(w, "Could not create token", http.StatusInternalServerError)
			return
		}

		// Send token as response
		w.Write([]byte(tokenString))
	}
}

// ProtectedEndpoint for accessing protected resources
func protectedEndpoint(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Welcome %s!", claims.Subject)
}

func main2() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/protected", protectedEndpoint)
	http.ListenAndServe(":8080", nil)
}
