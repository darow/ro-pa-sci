```shell
go run ./cmd/
```
[тестовый фронт](http://localhost:8000/swagger/index.html#)

[swagger документация http](http://localhost:8000/swagger/index.html#)

### websocket команды

создать приглашение
```shell
  {"action":"createInvite", "body":{"to":1}}
```

ответить на приглашение
decision 
0 - отклонить
1 - принять
```shell
  {"action":"decideInvite", "body":{"inviteID":1, "decision":0}}
```

