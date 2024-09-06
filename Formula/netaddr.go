package main

// Netaddr - Network address manipulation library
// Homepage: https://netaddr.readthedocs.io/en/latest/

import (
	"fmt"
	
	"os/exec"
)

func installNetaddr() {
	// Método 1: Descargar y extraer .tar.gz
	netaddr_tar_url := "https://files.pythonhosted.org/packages/54/90/188b2a69654f27b221fba92fda7217778208532c962509e959a9cee5229d/netaddr-1.3.0.tar.gz"
	netaddr_cmd_tar := exec.Command("curl", "-L", netaddr_tar_url, "-o", "package.tar.gz")
	err := netaddr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	netaddr_zip_url := "https://files.pythonhosted.org/packages/54/90/188b2a69654f27b221fba92fda7217778208532c962509e959a9cee5229d/netaddr-1.3.0.zip"
	netaddr_cmd_zip := exec.Command("curl", "-L", netaddr_zip_url, "-o", "package.zip")
	err = netaddr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	netaddr_bin_url := "https://files.pythonhosted.org/packages/54/90/188b2a69654f27b221fba92fda7217778208532c962509e959a9cee5229d/netaddr-1.3.0.bin"
	netaddr_cmd_bin := exec.Command("curl", "-L", netaddr_bin_url, "-o", "binary.bin")
	err = netaddr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	netaddr_src_url := "https://files.pythonhosted.org/packages/54/90/188b2a69654f27b221fba92fda7217778208532c962509e959a9cee5229d/netaddr-1.3.0.src.tar.gz"
	netaddr_cmd_src := exec.Command("curl", "-L", netaddr_src_url, "-o", "source.tar.gz")
	err = netaddr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	netaddr_cmd_direct := exec.Command("./binary")
	err = netaddr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
