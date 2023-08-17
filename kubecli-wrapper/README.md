# kubecli-wrapper

kubecli-wrapper is a simple command-line utility built in Go that provides a wrapper around `kubectl` commands to simplify interactions with Kubernetes clusters. It aims to streamline common tasks and enhance the user experience when working with Kubernetes.

## Features

- **List Pods:** Easily list pods in a specified namespace.
- *Add more features here...*

## Prerequisites

Before using `kubecli-wrapper`, ensure you have the following prerequisites:

- Go installed (version 1.13+)
- A running Kubernetes cluster
- `kubectl` command-line tool installed and configured

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/yourusername/kubecli-wrapper.git
   cd kubecli-wrapper
   ```

2. Fetch project dependencies:
    ```sh
    go get k8s.io/apimachinery/pkg/apis/meta/v1
    go get k8s.io/client-go/kubernetes
    go get k8s.io/client-go/tools/clientcmd
    ```

3.  Build the KubeCLI wrapper:
   ```sh
    go build
   ```

**Usage**
Run the kubecli-wrapper executable followed by the desired command and any relevant arguments.

* List pods in the default namespace:

   ```sh
    ./kubecli-wrapper list-pods

   ```
* List pods in a specific namespace (e.g., "kube-system"):
   ```sh
    ./kubecli-wrapper list-pods kube-system
   ```

**Contribution**
Contributions to kubecli-wrapper are welcome! Feel free to fork this repository, make improvements, and submit pull requests.

If you encounter any issues or have suggestions for new features, please open an issue on the issue tracker.

License
This project is licensed under the MIT License.



