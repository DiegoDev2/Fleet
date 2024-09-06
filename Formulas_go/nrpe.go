package main

// Nrpe - Nagios remote plugin executor
// Homepage: https://www.nagios.org/

import (
	"fmt"
	
	"os/exec"
)

func installNrpe() {
	// Método 1: Descargar y extraer .tar.gz
	nrpe_tar_url := "https://github.com/NagiosEnterprises/nrpe/releases/download/nrpe-4.1.1/nrpe-4.1.1.tar.gz"
	nrpe_cmd_tar := exec.Command("curl", "-L", nrpe_tar_url, "-o", "package.tar.gz")
	err := nrpe_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nrpe_zip_url := "https://github.com/NagiosEnterprises/nrpe/releases/download/nrpe-4.1.1/nrpe-4.1.1.zip"
	nrpe_cmd_zip := exec.Command("curl", "-L", nrpe_zip_url, "-o", "package.zip")
	err = nrpe_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nrpe_bin_url := "https://github.com/NagiosEnterprises/nrpe/releases/download/nrpe-4.1.1/nrpe-4.1.1.bin"
	nrpe_cmd_bin := exec.Command("curl", "-L", nrpe_bin_url, "-o", "binary.bin")
	err = nrpe_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nrpe_src_url := "https://github.com/NagiosEnterprises/nrpe/releases/download/nrpe-4.1.1/nrpe-4.1.1.src.tar.gz"
	nrpe_cmd_src := exec.Command("curl", "-L", nrpe_src_url, "-o", "source.tar.gz")
	err = nrpe_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nrpe_cmd_direct := exec.Command("./binary")
	err = nrpe_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: nagios-plugins")
exec.Command("latte", "install", "nagios-plugins")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
