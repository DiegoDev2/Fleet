package main

// Gel - Modern gem manager
// Homepage: https://gel.dev

import (
	"fmt"
	
	"os/exec"
)

func installGel() {
	// Método 1: Descargar y extraer .tar.gz
	gel_tar_url := "https://github.com/gel-rb/gel/archive/refs/tags/v0.3.0.tar.gz"
	gel_cmd_tar := exec.Command("curl", "-L", gel_tar_url, "-o", "package.tar.gz")
	err := gel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gel_zip_url := "https://github.com/gel-rb/gel/archive/refs/tags/v0.3.0.zip"
	gel_cmd_zip := exec.Command("curl", "-L", gel_zip_url, "-o", "package.zip")
	err = gel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gel_bin_url := "https://github.com/gel-rb/gel/archive/refs/tags/v0.3.0.bin"
	gel_cmd_bin := exec.Command("curl", "-L", gel_bin_url, "-o", "binary.bin")
	err = gel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gel_src_url := "https://github.com/gel-rb/gel/archive/refs/tags/v0.3.0.src.tar.gz"
	gel_cmd_src := exec.Command("curl", "-L", gel_src_url, "-o", "source.tar.gz")
	err = gel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gel_cmd_direct := exec.Command("./binary")
	err = gel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ronn")
	exec.Command("latte", "install", "ronn").Run()
}
