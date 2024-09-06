package main

// Bat - Clone of cat(1) with syntax highlighting and Git integration
// Homepage: https://github.com/sharkdp/bat

import (
	"fmt"
	
	"os/exec"
)

func installBat() {
	// Método 1: Descargar y extraer .tar.gz
	bat_tar_url := "https://github.com/sharkdp/bat/archive/refs/tags/v0.24.0.tar.gz"
	bat_cmd_tar := exec.Command("curl", "-L", bat_tar_url, "-o", "package.tar.gz")
	err := bat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bat_zip_url := "https://github.com/sharkdp/bat/archive/refs/tags/v0.24.0.zip"
	bat_cmd_zip := exec.Command("curl", "-L", bat_zip_url, "-o", "package.zip")
	err = bat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bat_bin_url := "https://github.com/sharkdp/bat/archive/refs/tags/v0.24.0.bin"
	bat_cmd_bin := exec.Command("curl", "-L", bat_bin_url, "-o", "binary.bin")
	err = bat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bat_src_url := "https://github.com/sharkdp/bat/archive/refs/tags/v0.24.0.src.tar.gz"
	bat_cmd_src := exec.Command("curl", "-L", bat_src_url, "-o", "source.tar.gz")
	err = bat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bat_cmd_direct := exec.Command("./binary")
	err = bat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: libgit2@1.7")
exec.Command("latte", "install", "libgit2@1.7")
	fmt.Println("Instalando dependencia: oniguruma")
exec.Command("latte", "install", "oniguruma")
}
