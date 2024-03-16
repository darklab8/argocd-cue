package sample

import (
  corev1 "k8s.io/api/core/v1"
)

#pod: [Name=_]: corev1.#Pod

#pod: [Name=_]: {
	apiVersion: "v1"
	kind:       string | *"Pod"
	metadata: {
		name: Name
		labels: app: "nginx"
	}
	spec: containers: [{
		image:           "nginx:1.25.4-alpine3.18"
		imagePullPolicy: "IfNotPresent"
		name:            Name
		ports: [{
			containerPort: 80
			protocol:      "TCP"
		}]
	}]
}


#pod: smth: {}
#pod: smth2: {}


objects: [
    #pod.smth,
    #pod.smth2
]