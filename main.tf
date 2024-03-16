# difficult to put into terraform in a normal way. difficult.

variable "context" {
  type = string
}

locals {
  plugin_files = fileset("${path.module}/plugin", "*")
}

resource "null_resource" "argo" {
  triggers = {
    for filename in local.plugin_files: filename => sha1(file("${path.module}/plugin/${filename}")) 
  }

  provisioner "local-exec" {
    command     = "kubectl apply --context ${var.context} -k plugin"
    working_dir = path.module
  }
}
