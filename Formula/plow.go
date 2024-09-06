package main

// Plow - High-performance and real-time metrics displaying HTTP benchmarking tool
// Homepage: https://github.com/six-ddc/plow

import (
	"fmt"
	
	"os/exec"
)

func installPlow() {
	// Método 1: Descargar y extraer .tar.gz
	plow_tar_url := "https://github.com/six-ddc/plow/archive/refs/tags/v1.3.1.tar.gz"
	plow_cmd_tar := exec.Command("curl", "-L", plow_tar_url, "-o", "package.tar.gz")
	err := plow_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	plow_zip_url := "https://github.com/six-ddc/plow/archive/refs/tags/v1.3.1.zip"
	plow_cmd_zip := exec.Command("curl", "-L", plow_zip_url, "-o", "package.zip")
	err = plow_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	plow_bin_url := "https://github.com/six-ddc/plow/archive/refs/tags/v1.3.1.bin"
	plow_cmd_bin := exec.Command("curl", "-L", plow_bin_url, "-o", "binary.bin")
	err = plow_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	plow_src_url := "https://github.com/six-ddc/plow/archive/refs/tags/v1.3.1.src.tar.gz"
	plow_cmd_src := exec.Command("curl", "-L", plow_src_url, "-o", "source.tar.gz")
	err = plow_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	plow_cmd_direct := exec.Command("./binary")
	err = plow_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
