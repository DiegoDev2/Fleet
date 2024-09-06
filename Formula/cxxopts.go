package main

// Cxxopts - Lightweight C++ command-line option parser
// Homepage: https://github.com/jarro2783/cxxopts

import (
	"fmt"
	
	"os/exec"
)

func installCxxopts() {
	// Método 1: Descargar y extraer .tar.gz
	cxxopts_tar_url := "https://github.com/jarro2783/cxxopts/archive/refs/tags/v3.2.1.tar.gz"
	cxxopts_cmd_tar := exec.Command("curl", "-L", cxxopts_tar_url, "-o", "package.tar.gz")
	err := cxxopts_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cxxopts_zip_url := "https://github.com/jarro2783/cxxopts/archive/refs/tags/v3.2.1.zip"
	cxxopts_cmd_zip := exec.Command("curl", "-L", cxxopts_zip_url, "-o", "package.zip")
	err = cxxopts_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cxxopts_bin_url := "https://github.com/jarro2783/cxxopts/archive/refs/tags/v3.2.1.bin"
	cxxopts_cmd_bin := exec.Command("curl", "-L", cxxopts_bin_url, "-o", "binary.bin")
	err = cxxopts_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cxxopts_src_url := "https://github.com/jarro2783/cxxopts/archive/refs/tags/v3.2.1.src.tar.gz"
	cxxopts_cmd_src := exec.Command("curl", "-L", cxxopts_src_url, "-o", "source.tar.gz")
	err = cxxopts_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cxxopts_cmd_direct := exec.Command("./binary")
	err = cxxopts_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
