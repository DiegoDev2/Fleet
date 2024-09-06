package main

// Logdy - Web based real-time log viewer
// Homepage: https://logdy.dev

import (
	"fmt"
	
	"os/exec"
)

func installLogdy() {
	// Método 1: Descargar y extraer .tar.gz
	logdy_tar_url := "https://github.com/logdyhq/logdy-core/archive/refs/tags/v0.13.0.tar.gz"
	logdy_cmd_tar := exec.Command("curl", "-L", logdy_tar_url, "-o", "package.tar.gz")
	err := logdy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	logdy_zip_url := "https://github.com/logdyhq/logdy-core/archive/refs/tags/v0.13.0.zip"
	logdy_cmd_zip := exec.Command("curl", "-L", logdy_zip_url, "-o", "package.zip")
	err = logdy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	logdy_bin_url := "https://github.com/logdyhq/logdy-core/archive/refs/tags/v0.13.0.bin"
	logdy_cmd_bin := exec.Command("curl", "-L", logdy_bin_url, "-o", "binary.bin")
	err = logdy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	logdy_src_url := "https://github.com/logdyhq/logdy-core/archive/refs/tags/v0.13.0.src.tar.gz"
	logdy_cmd_src := exec.Command("curl", "-L", logdy_src_url, "-o", "source.tar.gz")
	err = logdy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	logdy_cmd_direct := exec.Command("./binary")
	err = logdy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
