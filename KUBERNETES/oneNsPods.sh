# some pods are not going to run 
NAMESPACE="devops-lab"
kubectl create ns $NAMESPACE
kubectl run nginx --image nginx --namespace $NAMESPACE
kubectl run fail-pod-1 --image nginx --namespace $NAMESPACE -- /bin/sss asd
kubectl run fail-pod-2 --image nginx --namespace $NAMESPACE -- seeleeep 1
kubectl run busybox --image nginx --namespace $NAMESPACE -- sleep infinity
kubectl run busybox-copy --image nginx --namespace $NAMESPACE -- test test test
kubectl run fail-pod-3-fail --image nginx --restart="Never" --namespace $NAMESPACE -- /bin/bash exit 1
kubectl run busybox-fail --image nginx --restart="Never" --namespace $NAMESPACE -- /bin/bash exit 1 

