package main

// H3 - Hexagonal hierarchical geospatial indexing system
// Homepage: https://uber.github.io/h3/

import (
	"fmt"
	
	"os/exec"
)

func installH3() {
	// Método 1: Descargar y extraer .tar.gz
	h3_tar_url := "https://github.com/uber/h3/archive/refs/tags/v4.1.0.tar.gz"
	h3_cmd_tar := exec.Command("curl", "-L", h3_tar_url, "-o", "package.tar.gz")
	err := h3_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	h3_zip_url := "https://github.com/uber/h3/archive/refs/tags/v4.1.0.zip"
	h3_cmd_zip := exec.Command("curl", "-L", h3_zip_url, "-o", "package.zip")
	err = h3_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	h3_bin_url := "https://github.com/uber/h3/archive/refs/tags/v4.1.0.bin"
	h3_cmd_bin := exec.Command("curl", "-L", h3_bin_url, "-o", "binary.bin")
	err = h3_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	h3_src_url := "https://github.com/uber/h3/archive/refs/tags/v4.1.0.src.tar.gz"
	h3_cmd_src := exec.Command("curl", "-L", h3_src_url, "-o", "source.tar.gz")
	err = h3_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	h3_cmd_direct := exec.Command("./binary")
	err = h3_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
