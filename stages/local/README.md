# local stage

num | directory | target name    | usage              | firebase-emulators | apisvr         | uisvr                    | rproxy          | apisvr origin from uisvr | apisvr CORS 
---:|-----------|----------------|--------------------|--------------------|----------------|--------------------------|-----------------|--------------------------|------------
1   | dev       | up             | Manual test        | (9099),(4000)      | (8080)         | 4173                     | 8000            | (empty)                  | - 
2   | dev       | up_uisvr       | rproxy development | (9099),(4000)      | 8080           | 4173                     | (8000 on host)  | (empty)                  | -
3   | dev       | up_apisvr      | uisvr development  | 9099,(4000)        | 8080           | (5173 or 4173 on host)   | -               | http://localhost:8080    | http://localhost:4173,http://localhost:5173,http://localhost:10000
4   | dev       | up_middlewares | apisvr development | 9099,(4000)        | (8080 on host) | -                        | -               | http://localhost:8080    | http://localhost:4173,http://localhost:5173,http://localhost:10000
5   | test      | up             | E2E test           | -                  | -              | -                        | 8001            | (empty)                  | - 

`(port number)` means port which is exported but is not necessary to work. It's just for debug.
`4001` for firebase-emulator suite doesn't work well because firebase-emulator suite works only with 4000.
