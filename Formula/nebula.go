package main

// Nebula - Scalable overlay networking tool for connecting computers anywhere
// Homepage: https://github.com/slackhq/nebula

import (
	"fmt"
	
	"os/exec"
)

func installNebula() {
	// Método 1: Descargar y extraer .tar.gz
	nebula_tar_url := "https://github.com/slackhq/nebula/archive/refs/tags/v1.9.3.tar.gz"
	nebula_cmd_tar := exec.Command("curl", "-L", nebula_tar_url, "-o", "package.tar.gz")
	err := nebula_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nebula_zip_url := "https://github.com/slackhq/nebula/archive/refs/tags/v1.9.3.zip"
	nebula_cmd_zip := exec.Command("curl", "-L", nebula_zip_url, "-o", "package.zip")
	err = nebula_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nebula_bin_url := "https://github.com/slackhq/nebula/archive/refs/tags/v1.9.3.bin"
	nebula_cmd_bin := exec.Command("curl", "-L", nebula_bin_url, "-o", "binary.bin")
	err = nebula_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nebula_src_url := "https://github.com/slackhq/nebula/archive/refs/tags/v1.9.3.src.tar.gz"
	nebula_cmd_src := exec.Command("curl", "-L", nebula_src_url, "-o", "source.tar.gz")
	err = nebula_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nebula_cmd_direct := exec.Command("./binary")
	err = nebula_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
