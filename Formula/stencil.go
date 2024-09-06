package main

// Stencil - Smart templating engine for service development
// Homepage: https://engineering.outreach.io/stencil/

import (
	"fmt"
	
	"os/exec"
)

func installStencil() {
	// Método 1: Descargar y extraer .tar.gz
	stencil_tar_url := "https://github.com/getoutreach/stencil/archive/refs/tags/v1.39.1.tar.gz"
	stencil_cmd_tar := exec.Command("curl", "-L", stencil_tar_url, "-o", "package.tar.gz")
	err := stencil_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stencil_zip_url := "https://github.com/getoutreach/stencil/archive/refs/tags/v1.39.1.zip"
	stencil_cmd_zip := exec.Command("curl", "-L", stencil_zip_url, "-o", "package.zip")
	err = stencil_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stencil_bin_url := "https://github.com/getoutreach/stencil/archive/refs/tags/v1.39.1.bin"
	stencil_cmd_bin := exec.Command("curl", "-L", stencil_bin_url, "-o", "binary.bin")
	err = stencil_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stencil_src_url := "https://github.com/getoutreach/stencil/archive/refs/tags/v1.39.1.src.tar.gz"
	stencil_cmd_src := exec.Command("curl", "-L", stencil_src_url, "-o", "source.tar.gz")
	err = stencil_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stencil_cmd_direct := exec.Command("./binary")
	err = stencil_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
