package main

// Minisat - Minimalistic and high-performance SAT solver
// Homepage: https://github.com/stp/minisat

import (
	"fmt"
	
	"os/exec"
)

func installMinisat() {
	// Método 1: Descargar y extraer .tar.gz
	minisat_tar_url := "https://github.com/stp/minisat/archive/refs/tags/releases/2.2.1.tar.gz"
	minisat_cmd_tar := exec.Command("curl", "-L", minisat_tar_url, "-o", "package.tar.gz")
	err := minisat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	minisat_zip_url := "https://github.com/stp/minisat/archive/refs/tags/releases/2.2.1.zip"
	minisat_cmd_zip := exec.Command("curl", "-L", minisat_zip_url, "-o", "package.zip")
	err = minisat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	minisat_bin_url := "https://github.com/stp/minisat/archive/refs/tags/releases/2.2.1.bin"
	minisat_cmd_bin := exec.Command("curl", "-L", minisat_bin_url, "-o", "binary.bin")
	err = minisat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	minisat_src_url := "https://github.com/stp/minisat/archive/refs/tags/releases/2.2.1.src.tar.gz"
	minisat_cmd_src := exec.Command("curl", "-L", minisat_src_url, "-o", "source.tar.gz")
	err = minisat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	minisat_cmd_direct := exec.Command("./binary")
	err = minisat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
