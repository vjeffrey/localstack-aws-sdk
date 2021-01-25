# experiment for running localstack against AWS through the SDK


```
pip3 install localstack
localstack start
```

```
go run main.go
```

remove line 31 of main.go, `cfg.EndpointResolver = ...`, to have the program pick up your own credentials (via aws credential chain)
