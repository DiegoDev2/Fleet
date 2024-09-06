package main

// Tabiew - TUI to view and query delimited files (CSV/TSV/etc.)
// Homepage: https://github.com/shshemi/tabiew

import (
	"fmt"
	
	"os/exec"
)

func installTabiew() {
	// Método 1: Descargar y extraer .tar.gz
	tabiew_tar_url := "https://github.com/shshemi/tabiew/archive/refs/tags/v0.6.2.tar.gz"
	tabiew_cmd_tar := exec.Command("curl", "-L", tabiew_tar_url, "-o", "package.tar.gz")
	err := tabiew_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tabiew_zip_url := "https://github.com/shshemi/tabiew/archive/refs/tags/v0.6.2.zip"
	tabiew_cmd_zip := exec.Command("curl", "-L", tabiew_zip_url, "-o", "package.zip")
	err = tabiew_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tabiew_bin_url := "https://github.com/shshemi/tabiew/archive/refs/tags/v0.6.2.bin"
	tabiew_cmd_bin := exec.Command("curl", "-L", tabiew_bin_url, "-o", "binary.bin")
	err = tabiew_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tabiew_src_url := "https://github.com/shshemi/tabiew/archive/refs/tags/v0.6.2.src.tar.gz"
	tabiew_cmd_src := exec.Command("curl", "-L", tabiew_src_url, "-o", "source.tar.gz")
	err = tabiew_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tabiew_cmd_direct := exec.Command("./binary")
	err = tabiew_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
