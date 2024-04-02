
```shell
docker build -t pythonserver:latest .
```

```shell
docker tag pythonserver:latest fieelina/pythonserver:latest
```

```shell
docker push fieelina/pythonserver:latest
```

```shell
kubectl apply -f python-service-deployment-and-service.yaml
```