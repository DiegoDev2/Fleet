package main

// Docuum - Perform least recently used (LRU) eviction of Docker images
// Homepage: https://github.com/stepchowfun/docuum

import (
	"fmt"
	
	"os/exec"
)

func installDocuum() {
	// Método 1: Descargar y extraer .tar.gz
	docuum_tar_url := "https://github.com/stepchowfun/docuum/archive/refs/tags/v0.25.0.tar.gz"
	docuum_cmd_tar := exec.Command("curl", "-L", docuum_tar_url, "-o", "package.tar.gz")
	err := docuum_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	docuum_zip_url := "https://github.com/stepchowfun/docuum/archive/refs/tags/v0.25.0.zip"
	docuum_cmd_zip := exec.Command("curl", "-L", docuum_zip_url, "-o", "package.zip")
	err = docuum_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	docuum_bin_url := "https://github.com/stepchowfun/docuum/archive/refs/tags/v0.25.0.bin"
	docuum_cmd_bin := exec.Command("curl", "-L", docuum_bin_url, "-o", "binary.bin")
	err = docuum_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	docuum_src_url := "https://github.com/stepchowfun/docuum/archive/refs/tags/v0.25.0.src.tar.gz"
	docuum_cmd_src := exec.Command("curl", "-L", docuum_src_url, "-o", "source.tar.gz")
	err = docuum_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	docuum_cmd_direct := exec.Command("./binary")
	err = docuum_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
