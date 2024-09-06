package main

// Yorkie - Document store for collaborative applications
// Homepage: https://yorkie.dev/

import (
	"fmt"
	
	"os/exec"
)

func installYorkie() {
	// Método 1: Descargar y extraer .tar.gz
	yorkie_tar_url := "https://github.com/yorkie-team/yorkie/archive/refs/tags/v0.5.0.tar.gz"
	yorkie_cmd_tar := exec.Command("curl", "-L", yorkie_tar_url, "-o", "package.tar.gz")
	err := yorkie_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yorkie_zip_url := "https://github.com/yorkie-team/yorkie/archive/refs/tags/v0.5.0.zip"
	yorkie_cmd_zip := exec.Command("curl", "-L", yorkie_zip_url, "-o", "package.zip")
	err = yorkie_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yorkie_bin_url := "https://github.com/yorkie-team/yorkie/archive/refs/tags/v0.5.0.bin"
	yorkie_cmd_bin := exec.Command("curl", "-L", yorkie_bin_url, "-o", "binary.bin")
	err = yorkie_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yorkie_src_url := "https://github.com/yorkie-team/yorkie/archive/refs/tags/v0.5.0.src.tar.gz"
	yorkie_cmd_src := exec.Command("curl", "-L", yorkie_src_url, "-o", "source.tar.gz")
	err = yorkie_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yorkie_cmd_direct := exec.Command("./binary")
	err = yorkie_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
