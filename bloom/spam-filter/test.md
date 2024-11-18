
add item to spam list:
curl -X POST -d '{"item":"spam@example.com"}' -H "Content-Type: application/json" http://localhost:8080/add

check item:
curl -X POST -d '{"item":"spam@example.com"}' -H "Content-Type: application/json" http://localhost:8080/check

