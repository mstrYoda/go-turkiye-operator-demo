# Operator SDK - Creating Custom Resource & Controllers

- operator-sdk new hello-operator
- operator-sdk add api --api-version=goturkiye.com/v1 --kind=Hello
- operator-sdk add controller --api-version=goturkiye.com/v1 --kind=Hello
- operator-sdk generate crds && operator-sdk generate k8s

# Deploy & Run Operator

- kubectl apply -f deploy/crds/*_crd.yaml
- kubectl apply -f deploy/crds/*_cr.yaml
- operator-sdk run local
- operator-sdk build image-name:tag
