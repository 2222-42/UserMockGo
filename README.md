# UserMockGo


## POST User

### Create User

```
curl -XPOST -H "content-type:application/json" localhost:8080/users --data '{"email": "satoru.hanada@matsuri-tech.com", "password":"testtesttest", "password_confirmation":"testtesttest"}'
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

```
curl -XPOST -H "content-type:application/json" localhost:8080/user/reissue --data '{"email": "satoru.hanada@matsuri-tech.com", "password":"testtesttest", "password_confirmation":"testtesttest"}'
```

### Login

```
curl -XPOST -H "content-type:application/json" localhost:8080/user/login --data '{"email": "test3@test.com", "password":"testtesttest"}'
```

### MFA

```
curl -XPOST -H "content-type:application/json" localhost:8080/user/login --data '{"email": "test3@test.com", "password":"testtesttest"}'
```

### GetSelfInfo

no header
```
curl 'localhost:8080/users?id=1'
```

invalid token
```
curl -H 'X-Access-Token:test' 'localhost:8080/users?id=1'
```

