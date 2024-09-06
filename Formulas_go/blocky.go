package main

// Blocky - Fast and lightweight DNS proxy as ad-blocker for local network
// Homepage: https://0xerr0r.github.io/blocky

import (
	"fmt"
	
	"os/exec"
)

func installBlocky() {
	// Método 1: Descargar y extraer .tar.gz
	blocky_tar_url := "https://github.com/0xerr0r/blocky/archive/refs/tags/v0.24.tar.gz"
	blocky_cmd_tar := exec.Command("curl", "-L", blocky_tar_url, "-o", "package.tar.gz")
	err := blocky_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	blocky_zip_url := "https://github.com/0xerr0r/blocky/archive/refs/tags/v0.24.zip"
	blocky_cmd_zip := exec.Command("curl", "-L", blocky_zip_url, "-o", "package.zip")
	err = blocky_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	blocky_bin_url := "https://github.com/0xerr0r/blocky/archive/refs/tags/v0.24.bin"
	blocky_cmd_bin := exec.Command("curl", "-L", blocky_bin_url, "-o", "binary.bin")
	err = blocky_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	blocky_src_url := "https://github.com/0xerr0r/blocky/archive/refs/tags/v0.24.src.tar.gz"
	blocky_cmd_src := exec.Command("curl", "-L", blocky_src_url, "-o", "source.tar.gz")
	err = blocky_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	blocky_cmd_direct := exec.Command("./binary")
	err = blocky_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
