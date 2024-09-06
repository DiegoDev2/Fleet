package main

// Cloudlist - Tool for listing assets from multiple cloud providers
// Homepage: https://github.com/projectdiscovery/cloudlist

import (
	"fmt"
	
	"os/exec"
)

func installCloudlist() {
	// Método 1: Descargar y extraer .tar.gz
	cloudlist_tar_url := "https://github.com/projectdiscovery/cloudlist/archive/refs/tags/v1.1.0.tar.gz"
	cloudlist_cmd_tar := exec.Command("curl", "-L", cloudlist_tar_url, "-o", "package.tar.gz")
	err := cloudlist_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cloudlist_zip_url := "https://github.com/projectdiscovery/cloudlist/archive/refs/tags/v1.1.0.zip"
	cloudlist_cmd_zip := exec.Command("curl", "-L", cloudlist_zip_url, "-o", "package.zip")
	err = cloudlist_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cloudlist_bin_url := "https://github.com/projectdiscovery/cloudlist/archive/refs/tags/v1.1.0.bin"
	cloudlist_cmd_bin := exec.Command("curl", "-L", cloudlist_bin_url, "-o", "binary.bin")
	err = cloudlist_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cloudlist_src_url := "https://github.com/projectdiscovery/cloudlist/archive/refs/tags/v1.1.0.src.tar.gz"
	cloudlist_cmd_src := exec.Command("curl", "-L", cloudlist_src_url, "-o", "source.tar.gz")
	err = cloudlist_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cloudlist_cmd_direct := exec.Command("./binary")
	err = cloudlist_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
