variable "server_address" {
	description = "Contrail server address"
}

variable "username" {
	description = "Username used for authentication via keystone"
}

variable "tenant_name" {
	description = "Tenant to be used"
}

variable "password" {
	description = "Password used for authentication via keystone"
}

provider "contrail" {
	server = "${var.server_address}"
	auth_url = "http://${var.server_address}:5000/v2.0/"
	username = "${var.username}"
	tenant_name = "${var.tenant_name}"
	password = "${var.password}"
}

resource "contrail_virtual_network" "spocknet" {
	name = "spocknet"
	parent_fq_name = "default-domain:demo"
}
