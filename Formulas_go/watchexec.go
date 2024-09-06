package main

// Watchexec - Execute commands when watched files change
// Homepage: https://watchexec.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installWatchexec() {
	// Método 1: Descargar y extraer .tar.gz
	watchexec_tar_url := "https://github.com/watchexec/watchexec/archive/refs/tags/v2.1.2.tar.gz"
	watchexec_cmd_tar := exec.Command("curl", "-L", watchexec_tar_url, "-o", "package.tar.gz")
	err := watchexec_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	watchexec_zip_url := "https://github.com/watchexec/watchexec/archive/refs/tags/v2.1.2.zip"
	watchexec_cmd_zip := exec.Command("curl", "-L", watchexec_zip_url, "-o", "package.zip")
	err = watchexec_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	watchexec_bin_url := "https://github.com/watchexec/watchexec/archive/refs/tags/v2.1.2.bin"
	watchexec_cmd_bin := exec.Command("curl", "-L", watchexec_bin_url, "-o", "binary.bin")
	err = watchexec_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	watchexec_src_url := "https://github.com/watchexec/watchexec/archive/refs/tags/v2.1.2.src.tar.gz"
	watchexec_cmd_src := exec.Command("curl", "-L", watchexec_src_url, "-o", "source.tar.gz")
	err = watchexec_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	watchexec_cmd_direct := exec.Command("./binary")
	err = watchexec_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
