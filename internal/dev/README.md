## Prompt Testing UI

Guide to start Prompt UI for development purpose. To try out, you need
Kubernetes cluster accessible from your local machine and kubeconfig
file to communicate with API server of a cluster.

1. Copy local kubeconfig file to current directory (`prompt/internal/dev/`):

    ```
    cp <kube-config-file-path> kubeconfig.yaml
    ```

1. Set `KUBECTL_BIN` to your local kubectl bin path:

    ```
    export KUBECTL_BIN=$(which kubectl)
    ```

1. Set `AUDIT_LOG_FILE`:

    ```
    export AUDIT_LOG_FILE=$(pwd)/audit.log
    ```

1. Start a Prompt server:

    ```
    go run .
    ```

1. Open below URL in your browser:

    ```
    http://localhost:7009/v2/debug/ui/
    ```

1. Click on "kube-shell" button then you will be able to run all
   kubectl commands from UI.