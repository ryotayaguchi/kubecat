<img src="img/i_am_a_kat.png" width="250" alt="I am a cat">

# What is kubecat?
kubecat is not a cat obviously but might be a kube.

# How to bring kubecat to your environment

<img src="img/penguin_wrapped.png" width="250" alt="penguin wrapped">

## in local Linux environment
- see https://golang.org/doc/install to install golang

```
git clone git@github.com:ryotayaguchi/kubecat.git
cd kubecat
./local.sh deploy
./local.sh test
./local.sh undeploy
```

<img src="img/wheel_whale.png" width="250" alt="whell whale">

## in Docker environment
- see https://docs.docker.com/engine/install/ to install Docker Engine and start Docker service

```
git clone git@github.com:ryotayaguchi/kubecat.git
cd kubecat
./docker.sh deploy
./docker.sh test
./docker.sh undeploy
```

<img src="img/wheel_is_yummy.png" width="250" alt="wheel is yummy">

## in minikube environment
- see https://docs.docker.com/engine/install/ to install Docker Engine and start Docker service
- see https://minikube.sigs.k8s.io/docs/start/ to install and start minikube

```
git clone git@github.com:ryotayaguchi/kubecat.git
cd kubecat
./minikube.sh deploy
./minikube.sh test
./minikube.sh undeploy
```

# How it works?
It works somehow

## build
```bash
go build
```

## run
```bash
export ENABLE_LIVENESS=true
export ENABLE_READINESS=true
export TIME_UNTIL_LIVENESS_TRUE=5
export TIME_UNTIL_READINESS_TRUE=10
export TIME_FOR_PRESTOP=5
export TIME_UNTIL_AUTO_STOP=60
export PORT_ADMIN=8080
./kubecat &> /tmp/kubecat.log &
```

## run away

```bash
sleep 60
tail -1 /tmp/kubecat.log
INFO: I do NOT work for longer than 60 sec for saving my healthy. Ciao churu!
```

## are you alive?

question

```bash
curl localhost:8080/liveness/
```

answer
```
I'm dead. Please don't talk to me.
```

```
I'm alive. Please don't put me in a microwave.
```

## are you hungry?

question

```bash
curl localhost:8080/readiness/
```

answer

```
I'm unready. You lucky mouse. ignored at the moment.
```

```
I'm ready. You are allowed to bring my lunch here.
```

## are you working?

question

```bash
curl localhost:8080/uptime/
```

answer

```
I already worked seriously for 7 sec
```

## do you like CPU?

question

```bash
curl localhost:8080/cpu/
```

answer

```
ok. I ignited your expensive processors.
```

## do you like memory?

question

```bash
curl localhost:8080/memory/
```

answer in case you have enough amount of memory

```
ok. I scattered garbage in your memory space.
```

answer in case you have only few amount of memory

```
curl: (52) Empty reply from server
```

you can find OOM logs in `/tmp/kubecat.log`.

## release

request

```bash
curl localhost:8080/prestop/
```

response

```
ok. I'm ready for running away.
```

