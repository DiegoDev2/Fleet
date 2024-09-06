package main

// Envd - Reproducible development environment for AI/ML
// Homepage: https://envd.tensorchord.ai

import (
	"fmt"
	
	"os/exec"
)

func installEnvd() {
	// Método 1: Descargar y extraer .tar.gz
	envd_tar_url := "https://github.com/tensorchord/envd/archive/refs/tags/v0.3.45.tar.gz"
	envd_cmd_tar := exec.Command("curl", "-L", envd_tar_url, "-o", "package.tar.gz")
	err := envd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	envd_zip_url := "https://github.com/tensorchord/envd/archive/refs/tags/v0.3.45.zip"
	envd_cmd_zip := exec.Command("curl", "-L", envd_zip_url, "-o", "package.zip")
	err = envd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	envd_bin_url := "https://github.com/tensorchord/envd/archive/refs/tags/v0.3.45.bin"
	envd_cmd_bin := exec.Command("curl", "-L", envd_bin_url, "-o", "binary.bin")
	err = envd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	envd_src_url := "https://github.com/tensorchord/envd/archive/refs/tags/v0.3.45.src.tar.gz"
	envd_cmd_src := exec.Command("curl", "-L", envd_src_url, "-o", "source.tar.gz")
	err = envd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	envd_cmd_direct := exec.Command("./binary")
	err = envd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
