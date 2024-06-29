# local stage

| DEV_TARGET | fb-emu auth | fb-emu suite | apisvr | uisvr          | rproxy  | apisvr origin from uisvr | apisvr CORS                                 |
| ---------- | ----------- | ------------ | ------ | -------------- | ------- | ------------------------ | ------------------------------------------- |
| all        | 9099        | 4000         | 8080   | 4173           | 8000    | (same origin)            | http://localhost:4173,http://localhost:5173 |
| rproxy     | 9099        | 4000         | 8080   | 4173           | (8000 ) | (same origin)            | http://localhost:4173,http://localhost:5173 |
| uisvr      | 9099        | 4000         | 8080   | (5173 or 4173) | -       | http://localhost:8080    | http://localhost:4173,http://localhost:5173 |
| apisvr     | 9099        | 4000         | (8080) | -              | -       | http://localhost:8080    | http://localhost:4173,http://localhost:5173 |

`(port)` is opened by debugging environment.

- fb-emu auth: Firebase Emulators Authentication
- fb-emu suite: Firebase Emulators Suite
