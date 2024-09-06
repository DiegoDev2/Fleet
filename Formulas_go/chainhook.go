package main

// Chainhook - Reorg-aware indexing engine for the Stacks & Bitcoin blockchains
// Homepage: https://github.com/hirosystems/chainhook

import (
	"fmt"
	
	"os/exec"
)

func installChainhook() {
	// Método 1: Descargar y extraer .tar.gz
	chainhook_tar_url := "https://github.com/hirosystems/chainhook/archive/refs/tags/v1.8.0.tar.gz"
	chainhook_cmd_tar := exec.Command("curl", "-L", chainhook_tar_url, "-o", "package.tar.gz")
	err := chainhook_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chainhook_zip_url := "https://github.com/hirosystems/chainhook/archive/refs/tags/v1.8.0.zip"
	chainhook_cmd_zip := exec.Command("curl", "-L", chainhook_zip_url, "-o", "package.zip")
	err = chainhook_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chainhook_bin_url := "https://github.com/hirosystems/chainhook/archive/refs/tags/v1.8.0.bin"
	chainhook_cmd_bin := exec.Command("curl", "-L", chainhook_bin_url, "-o", "binary.bin")
	err = chainhook_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chainhook_src_url := "https://github.com/hirosystems/chainhook/archive/refs/tags/v1.8.0.src.tar.gz"
	chainhook_cmd_src := exec.Command("curl", "-L", chainhook_src_url, "-o", "source.tar.gz")
	err = chainhook_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chainhook_cmd_direct := exec.Command("./binary")
	err = chainhook_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
