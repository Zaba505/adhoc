#
# Knative Serving
#

# Install the required custom resources
kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.7.2/serving-crds.yaml

# Install the core components of Knative Serving
kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.7.2/serving-core.yaml

# Install the Knative Kourier controller
kubectl apply -f https://github.com/knative/net-kourier/releases/download/knative-v1.7.0/kourier.yaml

# Configure Knative Serving to use Kourier by default
kubectl patch configmap/config-network \
  --namespace knative-serving \
  --type merge \
  --patch '{"data":{"ingress-class":"kourier.ingress.networking.knative.dev"}}'

# Knative provides a Kubernetes Job called default-domain
# that configures Knative Serving to use sslip.io as the default DNS suffix.
kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.7.2/serving-default-domain.yaml

kubectl patch configmap/config-domain \
  -n knative-serving \
  --type merge \
  -p '{"data":{"127.0.0.1.sslip.io":""}}'

#
# Knative Serving
#

# Install the required custom resource definitions (CRDs)
kubectl apply -f https://github.com/knative/eventing/releases/download/knative-v1.7.3/eventing-crds.yaml

# Install the core components of Eventing
kubectl apply -f https://github.com/knative/eventing/releases/download/knative-v1.7.3/eventing-core.yaml
