package main

// Zellij - Pluggable terminal workspace, with terminal multiplexer as the base feature
// Homepage: https://zellij.dev

import (
	"fmt"
	
	"os/exec"
)

func installZellij() {
	// Método 1: Descargar y extraer .tar.gz
	zellij_tar_url := "https://github.com/zellij-org/zellij/archive/refs/tags/v0.40.1.tar.gz"
	zellij_cmd_tar := exec.Command("curl", "-L", zellij_tar_url, "-o", "package.tar.gz")
	err := zellij_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zellij_zip_url := "https://github.com/zellij-org/zellij/archive/refs/tags/v0.40.1.zip"
	zellij_cmd_zip := exec.Command("curl", "-L", zellij_zip_url, "-o", "package.zip")
	err = zellij_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zellij_bin_url := "https://github.com/zellij-org/zellij/archive/refs/tags/v0.40.1.bin"
	zellij_cmd_bin := exec.Command("curl", "-L", zellij_bin_url, "-o", "binary.bin")
	err = zellij_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zellij_src_url := "https://github.com/zellij-org/zellij/archive/refs/tags/v0.40.1.src.tar.gz"
	zellij_cmd_src := exec.Command("curl", "-L", zellij_src_url, "-o", "source.tar.gz")
	err = zellij_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zellij_cmd_direct := exec.Command("./binary")
	err = zellij_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
