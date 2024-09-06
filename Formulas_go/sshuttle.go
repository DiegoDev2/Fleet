package main

// Sshuttle - Proxy server that works as a poor man's VPN
// Homepage: https://github.com/sshuttle/sshuttle

import (
	"fmt"
	
	"os/exec"
)

func installSshuttle() {
	// Método 1: Descargar y extraer .tar.gz
	sshuttle_tar_url := "https://files.pythonhosted.org/packages/94/6e/f9a1fb50cd034cac1ee4efd017a9873301f75103271205a8f1c411a9fb1e/sshuttle-1.1.2.tar.gz"
	sshuttle_cmd_tar := exec.Command("curl", "-L", sshuttle_tar_url, "-o", "package.tar.gz")
	err := sshuttle_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sshuttle_zip_url := "https://files.pythonhosted.org/packages/94/6e/f9a1fb50cd034cac1ee4efd017a9873301f75103271205a8f1c411a9fb1e/sshuttle-1.1.2.zip"
	sshuttle_cmd_zip := exec.Command("curl", "-L", sshuttle_zip_url, "-o", "package.zip")
	err = sshuttle_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sshuttle_bin_url := "https://files.pythonhosted.org/packages/94/6e/f9a1fb50cd034cac1ee4efd017a9873301f75103271205a8f1c411a9fb1e/sshuttle-1.1.2.bin"
	sshuttle_cmd_bin := exec.Command("curl", "-L", sshuttle_bin_url, "-o", "binary.bin")
	err = sshuttle_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sshuttle_src_url := "https://files.pythonhosted.org/packages/94/6e/f9a1fb50cd034cac1ee4efd017a9873301f75103271205a8f1c411a9fb1e/sshuttle-1.1.2.src.tar.gz"
	sshuttle_cmd_src := exec.Command("curl", "-L", sshuttle_src_url, "-o", "source.tar.gz")
	err = sshuttle_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sshuttle_cmd_direct := exec.Command("./binary")
	err = sshuttle_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
