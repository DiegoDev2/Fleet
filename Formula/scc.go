package main

// Scc - Fast and accurate code counter with complexity and COCOMO estimates
// Homepage: https://github.com/boyter/scc/

import (
	"fmt"
	
	"os/exec"
)

func installScc() {
	// Método 1: Descargar y extraer .tar.gz
	scc_tar_url := "https://github.com/boyter/scc/archive/refs/tags/v3.3.5.tar.gz"
	scc_cmd_tar := exec.Command("curl", "-L", scc_tar_url, "-o", "package.tar.gz")
	err := scc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scc_zip_url := "https://github.com/boyter/scc/archive/refs/tags/v3.3.5.zip"
	scc_cmd_zip := exec.Command("curl", "-L", scc_zip_url, "-o", "package.zip")
	err = scc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scc_bin_url := "https://github.com/boyter/scc/archive/refs/tags/v3.3.5.bin"
	scc_cmd_bin := exec.Command("curl", "-L", scc_bin_url, "-o", "binary.bin")
	err = scc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scc_src_url := "https://github.com/boyter/scc/archive/refs/tags/v3.3.5.src.tar.gz"
	scc_cmd_src := exec.Command("curl", "-L", scc_src_url, "-o", "source.tar.gz")
	err = scc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scc_cmd_direct := exec.Command("./binary")
	err = scc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
