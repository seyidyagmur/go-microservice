I aim to learn Go by developing this application.

### Run project in your local

```bash
$ minikube start
$ helm repo add hazelcast https://hazelcast-charts.s3.amazonaws.com/ 
$ helm repo update 
$ helm install hz-hazelcast hazelcast/hazelcast --set cluster.memberCount=1 --set mancenter.enabled=false
```
Run below command in the project path to deploy sources to kubernetes environment.

```bash
$ skaffold dev
```

Since I am developing in a Windows environment, I need to run the code below and run the curl commands from within the relevant pod.
```bash
$ minikube ssh
```

### Contributor
These source codes are taken from the content on [Hüseyin Babal's](https://www.youtube.com/channel/UCJCuRftuTRXrdsLGhgS1Cnw) youtube channel.