## Sample Curl Commands


### User Service

1. POST /users/signup
```
curl -XPOST -d '{"email":"me@gmail.com","password":"ddlelle3rii","phone":"55555555555","user_type":0,"portal_access_status":1}' localhost:9090/api/v1/users/signup
```