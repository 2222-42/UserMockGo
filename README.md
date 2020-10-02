# UserMockGo


## POST User

### Create User

```
curl -XPOST -H "content-type:application/json" localhost:8080/users --data '{"email": "teste@test.com", "password":"testtesttest", "password_confirmation":"testtesttest"}'
```

### ActivateUser

```
curl 'localhost:8080/user/activate?email=test1@test.com&token=aaa'
```

### Reissue Token

failed: 
```
curl -XPOST -H "content-type:application/json" localhost:8080/user/reissue --data '{"email": "test3@test.com"}'
```
