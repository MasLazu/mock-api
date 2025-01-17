terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
  }
}

provider "digitalocean" {
  token = var.do_token
}

resource "digitalocean_droplet" "mock_server" {
  name     = "mock-server"
  region   = var.region
  size     = var.size
  image    = var.image
  ssh_keys = [var.ssh_fingerprint]

  tags = ["mock-server"]
}

output "droplet_ip" {
  value = digitalocean_droplet.mock_server.ipv4_address
}
