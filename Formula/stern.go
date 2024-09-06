package main

// Stern - Tail multiple Kubernetes pods & their containers
// Homepage: https://github.com/stern/stern

import (
	"fmt"
	
	"os/exec"
)

func installStern() {
	// Método 1: Descargar y extraer .tar.gz
	stern_tar_url := "https://github.com/stern/stern/archive/refs/tags/v1.30.0.tar.gz"
	stern_cmd_tar := exec.Command("curl", "-L", stern_tar_url, "-o", "package.tar.gz")
	err := stern_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stern_zip_url := "https://github.com/stern/stern/archive/refs/tags/v1.30.0.zip"
	stern_cmd_zip := exec.Command("curl", "-L", stern_zip_url, "-o", "package.zip")
	err = stern_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stern_bin_url := "https://github.com/stern/stern/archive/refs/tags/v1.30.0.bin"
	stern_cmd_bin := exec.Command("curl", "-L", stern_bin_url, "-o", "binary.bin")
	err = stern_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stern_src_url := "https://github.com/stern/stern/archive/refs/tags/v1.30.0.src.tar.gz"
	stern_cmd_src := exec.Command("curl", "-L", stern_src_url, "-o", "source.tar.gz")
	err = stern_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stern_cmd_direct := exec.Command("./binary")
	err = stern_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
