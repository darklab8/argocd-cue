
package templates

import (
    appsv1 "k8s.io/api/apps/v1"
)

#deploy: appsv1.#Deployment

#deploy: {
  apiVersion: "apps/v1"
  kind: "Deployment"
  metadata: {
    name: "helm-sample-deploy"
    labels:
      app: "helm-sample"
  }
  spec:{
  revisionHistoryLimit: 5
    strategy: {
      type: "RollingUpdate"
      rollingUpdate: {
        maxSurge: 1
        maxUnavailable: 0
      }
    }
    selector:
      matchLabels:
        project: "helm-sample"
    template: {

      metadata:
        labels:
          project: "helm-sample"
      spec: {
        containers: [
          {
            name: "helm-sample"
            image: "nginx:1.25.4-alpine3.18"
            ports: [
              {
                containerPort: 80
            }
          ]
        }
      ]
      }
    }
  }
}

files: [
  {
    #deploy
  }
]