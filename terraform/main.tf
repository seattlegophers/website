provider "azurerm" {
  subscription_id = "${var.subscription_id}"
  client_id = "${var.client_id}"
  client_secret = "${var.client_secret}"
  tenant_id = "${var.tenant_id}"
}


resource "azurerm_resource_group" "seattleGo-ResourceGroup" {
    name     = "myResourceGroup"
    location = "${var.location}"

    tags = {
        environment = "${var.prefix}-seattleGo"
    }
}

resource "azurerm_virtual_network" "seattleGo-Vnet" {
    name                = "seattleGoVnet"
    address_space       = ["10.0.0.0/16"]
    location            = "${var.location}"
    resource_group_name = "${azurerm_resource_group.seattleGo-ResourceGroup.name}"

    tags = {
        environment = "${var.prefix}-seattleGo"
    }
}


resource "azurerm_subnet" "seattleGo-PublicSubnet" {
    name                 = "seattleGo-PublicSubnet"
    resource_group_name  = "${azurerm_resource_group.seattleGo-ResourceGroup.name}"
    virtual_network_name = "${azurerm_virtual_network.seattleGo-Vnet.name}"
    address_prefix       = "10.0.2.0/24"
}

resource "azurerm_public_ip" "seattleGo-PublicIP" {
    name                         = "seattleGo-PublicIP"
    location                     = "${var.location}"
    resource_group_name          = "${azurerm_resource_group.seattleGo-ResourceGroup.name}"
    allocation_method            = "Static"
    #resolves to seattlego.westus.cloudapp.azure.com
    domain_name_label = "seattlego"

    tags = {
        environment = "${var.prefix}-seattleGo"
    }
}

resource "azurerm_network_security_group" "seattleGo-SG" {
    name                = "myNetworkSecurityGroup"
    location            = "${var.location}"
    resource_group_name = "${azurerm_resource_group.seattleGo-ResourceGroup.name}"



    tags = {
        environment = "${var.prefix}-seattleGo"
    }
}
resource "azurerm_network_security_rule" "seattleGo-SG_SSH_Inbound" {

        name                       = "SSH_inbound"
        priority                   = 1002
        direction                  = "Inbound"
        access                     = "Allow"
        protocol                   = "Tcp"
        source_port_range          = "*"
        destination_port_range     = "22"
        source_address_prefix      = "*"
        destination_address_prefix = "*"
        resource_group_name        = "${azurerm_resource_group.seattleGo-ResourceGroup.name}"
        network_security_group_name = "${azurerm_network_security_group.seattleGo-SG.name}"
    }

resource "azurerm_network_security_rule" "seattleGo-SG_http_outbound" {

        name                       = "http_outbound"
        priority                   = 1001
        direction                  = "Outbound"
        access                     = "Allow"
        protocol                   = "Tcp"
        source_port_range          = "*"
        destination_port_range     = "80"
        source_address_prefix      = "*"
        destination_address_prefix = "*"
        resource_group_name        = "${azurerm_resource_group.seattleGo-ResourceGroup.name}"
        network_security_group_name = "${azurerm_network_security_group.seattleGo-SG.name}"
}
resource "azurerm_network_security_rule" "seattleGo-SG_http_inbound" {

        name                       = "http_inbound"
        priority                   = 1001
        direction                  = "Inbound"
        access                     = "Allow"
        protocol                   = "Tcp"
        source_port_range          = "*"
        destination_port_range     = "80"
        source_address_prefix      = "*"
        destination_address_prefix = "*"
        resource_group_name        = "${azurerm_resource_group.seattleGo-ResourceGroup.name}"
        network_security_group_name = "${azurerm_network_security_group.seattleGo-SG.name}"
}
resource "azurerm_network_security_rule" "seattleGo-SG_app_inbound" {

        name                       = "http_inbound"
        priority                   = 1003
        direction                  = "Inbound"
        access                     = "Allow"
        protocol                   = "Tcp"
        source_port_range          = "*"
        destination_port_range     = "3000"
        source_address_prefix      = "*"
        destination_address_prefix = "*"
        resource_group_name        = "${azurerm_resource_group.seattleGo-ResourceGroup.name}"
        network_security_group_name = "${azurerm_network_security_group.seattleGo-SG.name}"
}

resource "azurerm_network_interface" "seattleGo-NIC" {
    name                = "myNIC"
    location            = "${var.location}"
    resource_group_name = "${azurerm_resource_group.seattleGo-ResourceGroup.name}"
    network_security_group_id = "${azurerm_network_security_group.seattleGo-SG.id}"

    ip_configuration {
        name                          = "myNicConfiguration"
        subnet_id                     = "${azurerm_subnet.seattleGo-PublicSubnet.id}"
        private_ip_address_allocation = "Dynamic"
        public_ip_address_id          = "${azurerm_public_ip.seattleGo-PublicIP.id}"
    }

    tags = {
        environment = "${var.prefix}-seattleGo"
    }
}

resource "azurerm_virtual_machine" "seattleGo-BackEnd" {
    name                  = "VM"
    location              = "${var.location}"
    resource_group_name   = "${azurerm_resource_group.seattleGo-ResourceGroup.name}"
    network_interface_ids = ["${azurerm_network_interface.seattleGo-NIC.id}"]
    vm_size               = "Standard_DS1_v2"

    storage_os_disk {
        name              = "myOsDisk"
        caching           = "ReadWrite"
        create_option     = "FromImage"
        managed_disk_type = "Premium_LRS"
    }

    storage_image_reference {
        publisher = "Canonical"
        offer     = "UbuntuServer"
        sku       = "${lookup(var.sku, var.location)}"
        version   = "latest"
    }

    os_profile {
        computer_name  = "backEnd"
        admin_username = "ubuntuSecure"
    }

    os_profile_linux_config {
        disable_password_authentication = true
        ssh_keys {
            path     = "/home/ubuntuSecure/.ssh/authorized_keys"
            key_data = "${file("~/.ssh/id_rsa.pub")}"
        }
    }

    tags = {
        environment = "${var.prefix}-seattleGo"
    }
}
output "public_IP" {
  value = "${azurerm_public_ip.seattleGo-PublicIP.ip_address}"
}
