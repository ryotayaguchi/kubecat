#!/bin/bash
set -eu

function f_build(){
    cd ${BASE_DIR}
    eval $(minikube -p minikube docker-env)
    docker build -t ${NAME} .
}

function f_undeploy(){
    if [[ `minikube kubectl -- get service/${NAME}` ]]; then
        minikube kubectl -- delete -f minikube/service.yml
    fi
    if [[ `minikube kubectl -- get deployment/${NAME}` ]]; then
        minikube kubectl -- delete -f minikube/deployment.yml
    fi
    if [[ `minikube kubectl -- get configmap/${NAME}` ]]; then
        minikube kubectl -- delete -f minikube/configmap.yml
    fi
}

function f_deploy(){
    f_undeploy
    f_build
    minikube kubectl -- apply -f minikube/configmap.yml
    sed "s/DeploymentHash/$(date)/g" minikube/deployment.yml | minikube kubectl -- apply -f -
    minikube kubectl -- rollout status deployment/${NAME}
    minikube kubectl -- apply -f minikube/service.yml
}

function f_restart(){
    minikube kubectl -- rollout restart deployment/${NAME}
}

function f_test(){
    IP=$(minikube ip)
    PORT=$(minikube kubectl -- get svc/${NAME} -o jsonpath='{.spec.ports[*].nodePort}')
    echo "cURL /uptime/: $(curl -s ${IP}:${PORT}/uptime/)"
    echo "cURL /liveness/: $(curl -s ${IP}:${PORT}/liveness/)"
    echo "cURL /readiness/: $(curl -s ${IP}:${PORT}/readiness/)"
}


NAME=kubecat
BASE_DIR=$(cd `dirname $0`; pwd)

[[ $# != 1 ]] && exit 1
case ${1} in
    "build")    f_build;;
    "undeploy") f_undeploy;;
    "deploy")   f_deploy;;
    "restart")  f_restart;;
    "test")     f_test;;
    *)          ;;
esac

