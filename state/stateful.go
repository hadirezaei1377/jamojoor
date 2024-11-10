package state

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID       int
	Username string
	Password string
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	http.HandleFunc("/login", loginHandler)
	http.ListenAndServe(":8080", nil)
}

// loginHandler for stateful authentication
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		var user User
		err := db.QueryRow("SELECT id, username FROM users WHERE username = ? AND password = ?", username, password).Scan(&user.ID, &user.Username)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// Create session (in practice, store this in session storage)
		sessionID := createSession(user.ID) // Function to create a session in the database
		http.SetCookie(w, &http.Cookie{Name: "session_id", Value: sessionID, Path: "/"})
		fmt.Fprintf(w, "User logged in: %s", user.Username)
	}
}

// createSession simulates creating a session in the database
func createSession(userID int) string {
	// Here, we would generate a session ID and store it in the session table
	return fmt.Sprintf("session_%d", userID)
}
