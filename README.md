#####
## call the login service  
TOKEN=$(curl -s -v -X GET http://localhost:8081/login)


## call the service passing the TOKEN  
curl -H 'Accept: application/json' -H "Token: ${TOKEN}" http://localhost:8080/oci
#####

