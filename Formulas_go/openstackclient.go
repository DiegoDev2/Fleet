package main

// Openstackclient - Command-line client for OpenStack
// Homepage: https://openstack.org

import (
	"fmt"
	
	"os/exec"
)

func installOpenstackclient() {
	// Método 1: Descargar y extraer .tar.gz
	openstackclient_tar_url := "https://files.pythonhosted.org/packages/92/fc/c3c17c98ba811a364094ff35057475dcc1ceaeb5edf4f9884542644982d0/python-openstackclient-7.0.0.tar.gz"
	openstackclient_cmd_tar := exec.Command("curl", "-L", openstackclient_tar_url, "-o", "package.tar.gz")
	err := openstackclient_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openstackclient_zip_url := "https://files.pythonhosted.org/packages/92/fc/c3c17c98ba811a364094ff35057475dcc1ceaeb5edf4f9884542644982d0/python-openstackclient-7.0.0.zip"
	openstackclient_cmd_zip := exec.Command("curl", "-L", openstackclient_zip_url, "-o", "package.zip")
	err = openstackclient_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openstackclient_bin_url := "https://files.pythonhosted.org/packages/92/fc/c3c17c98ba811a364094ff35057475dcc1ceaeb5edf4f9884542644982d0/python-openstackclient-7.0.0.bin"
	openstackclient_cmd_bin := exec.Command("curl", "-L", openstackclient_bin_url, "-o", "binary.bin")
	err = openstackclient_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openstackclient_src_url := "https://files.pythonhosted.org/packages/92/fc/c3c17c98ba811a364094ff35057475dcc1ceaeb5edf4f9884542644982d0/python-openstackclient-7.0.0.src.tar.gz"
	openstackclient_cmd_src := exec.Command("curl", "-L", openstackclient_src_url, "-o", "source.tar.gz")
	err = openstackclient_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openstackclient_cmd_direct := exec.Command("./binary")
	err = openstackclient_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: cryptography")
exec.Command("latte", "install", "cryptography")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
