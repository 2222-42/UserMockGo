# UserMockGo


## POST User

### Create User

```
curl -XPOST -H "content-type:application/json" localhost:8080/users --data '{"email": "teste@test.com", "password":"test", "password_confirmation":"test"}'
```

### ActivateUser

```
curl -XPOST -H "content-type:application/json" localhost:8080/user/activate --data '{"email": "test1@test.com", "token":"aaa"}'
```
