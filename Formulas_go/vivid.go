package main

// Vivid - Generator for LS_COLORS with support for multiple color themes
// Homepage: https://github.com/sharkdp/vivid

import (
	"fmt"
	
	"os/exec"
)

func installVivid() {
	// Método 1: Descargar y extraer .tar.gz
	vivid_tar_url := "https://github.com/sharkdp/vivid/archive/refs/tags/v0.10.1.tar.gz"
	vivid_cmd_tar := exec.Command("curl", "-L", vivid_tar_url, "-o", "package.tar.gz")
	err := vivid_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vivid_zip_url := "https://github.com/sharkdp/vivid/archive/refs/tags/v0.10.1.zip"
	vivid_cmd_zip := exec.Command("curl", "-L", vivid_zip_url, "-o", "package.zip")
	err = vivid_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vivid_bin_url := "https://github.com/sharkdp/vivid/archive/refs/tags/v0.10.1.bin"
	vivid_cmd_bin := exec.Command("curl", "-L", vivid_bin_url, "-o", "binary.bin")
	err = vivid_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vivid_src_url := "https://github.com/sharkdp/vivid/archive/refs/tags/v0.10.1.src.tar.gz"
	vivid_cmd_src := exec.Command("curl", "-L", vivid_src_url, "-o", "source.tar.gz")
	err = vivid_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vivid_cmd_direct := exec.Command("./binary")
	err = vivid_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
