module "cluster" {
  source = "./cluster"

  master_username = "${var.master_username}"
  master_password = "${var.master_password}"
  back_end_tag    = "${var.back_end_tag}"
  zone            = "${var.zone}"

  providers = {
    google = "google.cluster"
  }
}

# TODO: Add secrets to container
# resource "kubernetes_secret" "ksec" {}

resource "kubernetes_service" "kser" {
  metadata {
    name = "antifreeze-kser"
  }

  spec {
    selector {
      App = "${kubernetes_pod.kp.metadata.0.labels.App}"
    }

    port {
      port        = 8081
      target_port = 8081
    }

    type = "LoadBalancer"

    # Assign static IP to this service's load balancer
    load_balancer_ip = "${google_compute_address.addr.address}"
  }
}

resource "kubernetes_pod" "kp" {
  metadata {
    name = "antifreeze-kp"

    # Used to select this pod in kubernetes_service
    labels {
      App = "antifreeze"
    }
  }

  spec {
    container {
      image = "nilsgs/antifreeze"

      # TODO: Increase efficiency by specifying version
      # Ensures that the container is updated
      image_pull_policy = "Always"

      name = "antifreeze-kc"

      # List of ports to expose
      port {
        # This is the port for the server
        container_port = 8081
      }
    }
  }
}

# Configuration of static IP
resource "google_compute_address" "addr" {
  name = "antifreeze-addr"
}
