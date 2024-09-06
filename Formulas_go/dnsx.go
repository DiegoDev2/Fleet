package main

// Dnsx - DNS query and resolution tool
// Homepage: https://github.com/projectdiscovery/dnsx

import (
	"fmt"
	
	"os/exec"
)

func installDnsx() {
	// Método 1: Descargar y extraer .tar.gz
	dnsx_tar_url := "https://github.com/projectdiscovery/dnsx/archive/refs/tags/v1.2.1.tar.gz"
	dnsx_cmd_tar := exec.Command("curl", "-L", dnsx_tar_url, "-o", "package.tar.gz")
	err := dnsx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dnsx_zip_url := "https://github.com/projectdiscovery/dnsx/archive/refs/tags/v1.2.1.zip"
	dnsx_cmd_zip := exec.Command("curl", "-L", dnsx_zip_url, "-o", "package.zip")
	err = dnsx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dnsx_bin_url := "https://github.com/projectdiscovery/dnsx/archive/refs/tags/v1.2.1.bin"
	dnsx_cmd_bin := exec.Command("curl", "-L", dnsx_bin_url, "-o", "binary.bin")
	err = dnsx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dnsx_src_url := "https://github.com/projectdiscovery/dnsx/archive/refs/tags/v1.2.1.src.tar.gz"
	dnsx_cmd_src := exec.Command("curl", "-L", dnsx_src_url, "-o", "source.tar.gz")
	err = dnsx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dnsx_cmd_direct := exec.Command("./binary")
	err = dnsx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
