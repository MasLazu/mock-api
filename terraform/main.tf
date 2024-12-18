provider "digitalocean" {
  token = var.do_token
}

resource "digitalocean_droplet" "mock_server" {
  name   = "mock-server"
  region = var.region
  size   = var.size
  image  = var.image

  ssh_keys = [
    var.ssh_key_name
  ]

  tags = ["mock-server"]
}

output "droplet_ip" {
  value = digitalocean_droplet.mock_server.ipv4_address
}
