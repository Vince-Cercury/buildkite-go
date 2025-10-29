# buildkite-go
Little exercice to build a go project with buildkite docker


## Kubernetes preparation

You need a Kubernetes cluster running

Edit the file k8s/agent-stack-config.yaml and add your Buildkite agent token.
Then run the helm upgrade

```
helm upgrade --install agent-stack-k8s oci://ghcr.io/buildkite/helm/agent-stack-k8s \
    --namespace buildkite \
    --create-namespace \
    --values agent-stack-config.yaml
```

This will run one pod with the Buildkite controller.

To use this K8S based agent, specify the correct queue in your pipeline:

```
    agents:
      queue: gcloud
```

where gcloud is the name of the queue with the agent token you have used above