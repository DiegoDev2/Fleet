package main

// Squealer - Scans Git repositories or filesystems for secrets in commit histories
// Homepage: https://github.com/owenrumney/squealer

import (
	"fmt"
	
	"os/exec"
)

func installSquealer() {
	// Método 1: Descargar y extraer .tar.gz
	squealer_tar_url := "https://github.com/owenrumney/squealer/archive/refs/tags/v1.2.4.tar.gz"
	squealer_cmd_tar := exec.Command("curl", "-L", squealer_tar_url, "-o", "package.tar.gz")
	err := squealer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	squealer_zip_url := "https://github.com/owenrumney/squealer/archive/refs/tags/v1.2.4.zip"
	squealer_cmd_zip := exec.Command("curl", "-L", squealer_zip_url, "-o", "package.zip")
	err = squealer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	squealer_bin_url := "https://github.com/owenrumney/squealer/archive/refs/tags/v1.2.4.bin"
	squealer_cmd_bin := exec.Command("curl", "-L", squealer_bin_url, "-o", "binary.bin")
	err = squealer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	squealer_src_url := "https://github.com/owenrumney/squealer/archive/refs/tags/v1.2.4.src.tar.gz"
	squealer_cmd_src := exec.Command("curl", "-L", squealer_src_url, "-o", "source.tar.gz")
	err = squealer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	squealer_cmd_direct := exec.Command("./binary")
	err = squealer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
