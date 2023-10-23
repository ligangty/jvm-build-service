#!/bin/sh

kubectl delete --ignore-not-found deployments.apps hacbs-jvm-operator -n jvm-build-service
# we don't restart the cache and local storage by default
# for most cases in development this is not necessary, and just slows things
# down by needing things to be re-cached/rebuilt

function cleanAllArtifacts() {
     kubectl delete --ignore-not-found namespaces test-jvm-namespace
}

kubectl delete --ignore-not-found deployments.apps jvm-build-workspace-artifact-cache
if [ "$1" = "--clean" ]; then
    cleanAllArtifacts
fi

DIR=`dirname $0`
kubectl apply -f $DIR/namespace.yaml
kubectl config set-context --current --namespace=test-jvm-namespace
kubectl delete --ignore-not-found secret jvm-build-image-secrets jvm-build-git-secrets jvm-build-maven-repo-secrets

if [ -n "$QUAY_ORG" ] && [ -n "$QUAY_TOKEN" ]; then
    kubectl delete --ignore-not-found secret  -n image-controller quaytoken
    kubectl create secret generic -n image-controller quaytoken --from-literal "quaytoken=$QUAY_TOKEN" --from-literal "organization=$QUAY_ORG"
fi
kubectl create secret generic jvm-build-image-secrets --from-file=.dockerconfigjson=$HOME/.docker/config.json --type=kubernetes.io/dockerconfigjson
kubectl create secret generic jvm-build-git-secrets --from-literal .git-credentials="
https://$GITHUB_E2E_ORGANIZATION:$GITHUB_TOKEN@github.com
https://test:test@gitlab.com
"
if [ -n "$MAVEN_PASSWORD" ]; then
    kubectl create secret generic jvm-build-maven-repo-secrets --from-literal mavenpassword="$MAVEN_PASSWORD"
fi

JVM_BUILD_SERVICE_IMAGE=quay.io/$QUAY_USERNAME/hacbs-jvm-controller \
JVM_BUILD_SERVICE_CACHE_IMAGE=quay.io/$QUAY_USERNAME/hacbs-jvm-cache \
JVM_BUILD_SERVICE_REQPROCESSOR_IMAGE=quay.io/$QUAY_USERNAME/hacbs-jvm-build-request-processor:dev \
$DIR/patch-yaml.sh
kubectl apply -k $DIR/overlays/development
