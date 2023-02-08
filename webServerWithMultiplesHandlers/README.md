## INTRO
TODO

## Explain
TODO

## Testing
You can make http request to test the POST endpoint like that:

```bash
$ curl -X POST -H 'Content-Type: application/json' -d '{"test":"hola"}' localhost:8080/create
$ curl -X POST -H 'Content-Type: application/json' -d '{"ZONE":"hola", "application":"test", "environment":"pre", "component":"dani", "pv":"mi_pv"}' localhost:8080/create
```
