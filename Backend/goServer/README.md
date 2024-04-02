
```shell
docker build -t goserver:latest .
```

```shell
docker tag goserver:latest fieelina/goserver:latest
```

```shell
docker push fieelina/goserver:latest
```

```shell
kubectl get svc
NAME             TYPE        CLUSTER-IP        EXTERNAL-IP   PORT(S)          AGE
kubernetes       ClusterIP   192.168.194.129   <none>        443/TCP          44m
python-service   NodePort    192.168.194.229   <none>        5000:30752/TCP   42m
go-service       NodePort    192.168.194.230   <none>        8080:32502/TCP   17m
```


```shell
kubectl port-forward svc/python-service 5000:5000
```

```shell
kubectl cluster-info
```

```shell
kubectl get ingress
```

```shell
kubectl describe ingress example-ingress
```

查看集群中Ingress控制器是否运行
```shell
kubectl get pods --all-namespaces | grep ingress
```

### 安装NGINX Ingress Controller
NGINX Ingress Controller可以通过多种方式安装，包括使用Helm或直接使用YAML文件。这里，我们将通过YAML文件来进行安装。


### 获取安装YAML文件
官方NGINX Ingress Controller提供了一个用于安装的YAML文件。你可以从官方GitHub仓库下载最新版本的部署YAML文件。为了确保操作的简便性，这里提供一个命令来直接应用最新版的NGINX Ingress Controller部署文件：

```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/cloud/deploy.yaml
```
注意：这个命令使用的是针对云提供商环境（如GCP、AWS、Azure等）的通用部署文件。如果你的Kubernetes集群部署在特定环境下（例如裸机），可能需要调整部署文件或选择不同的安装方法。请参考NGINX Ingress Controller的官方文档以获取更适合你环境的安装步骤。