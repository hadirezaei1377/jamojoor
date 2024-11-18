go run main.go

to see profiles:
http://localhost:6060/debug/pprof/


download cpu profiles for 30 seconds:
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30


see profiles in cli:
go tool pprof <file_name>


see results in browser:
go tool pprof -http=:8080 <file_name>


