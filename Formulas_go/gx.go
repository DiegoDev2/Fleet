package main

// Gx - Language-agnostic, universal package manager
// Homepage: https://github.com/whyrusleeping/gx

import (
	"fmt"
	
	"os/exec"
)

func installGx() {
	// Método 1: Descargar y extraer .tar.gz
	gx_tar_url := "https://github.com/whyrusleeping/gx/archive/refs/tags/v0.14.3.tar.gz"
	gx_cmd_tar := exec.Command("curl", "-L", gx_tar_url, "-o", "package.tar.gz")
	err := gx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gx_zip_url := "https://github.com/whyrusleeping/gx/archive/refs/tags/v0.14.3.zip"
	gx_cmd_zip := exec.Command("curl", "-L", gx_zip_url, "-o", "package.zip")
	err = gx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gx_bin_url := "https://github.com/whyrusleeping/gx/archive/refs/tags/v0.14.3.bin"
	gx_cmd_bin := exec.Command("curl", "-L", gx_bin_url, "-o", "binary.bin")
	err = gx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gx_src_url := "https://github.com/whyrusleeping/gx/archive/refs/tags/v0.14.3.src.tar.gz"
	gx_cmd_src := exec.Command("curl", "-L", gx_src_url, "-o", "source.tar.gz")
	err = gx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gx_cmd_direct := exec.Command("./binary")
	err = gx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
