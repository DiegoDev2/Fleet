package main

// Goproxy - Global proxy for Go modules
// Homepage: https://github.com/goproxyio/goproxy

import (
	"fmt"
	
	"os/exec"
)

func installGoproxy() {
	// Método 1: Descargar y extraer .tar.gz
	goproxy_tar_url := "https://github.com/goproxyio/goproxy/archive/refs/tags/v2.0.7.tar.gz"
	goproxy_cmd_tar := exec.Command("curl", "-L", goproxy_tar_url, "-o", "package.tar.gz")
	err := goproxy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	goproxy_zip_url := "https://github.com/goproxyio/goproxy/archive/refs/tags/v2.0.7.zip"
	goproxy_cmd_zip := exec.Command("curl", "-L", goproxy_zip_url, "-o", "package.zip")
	err = goproxy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	goproxy_bin_url := "https://github.com/goproxyio/goproxy/archive/refs/tags/v2.0.7.bin"
	goproxy_cmd_bin := exec.Command("curl", "-L", goproxy_bin_url, "-o", "binary.bin")
	err = goproxy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	goproxy_src_url := "https://github.com/goproxyio/goproxy/archive/refs/tags/v2.0.7.src.tar.gz"
	goproxy_cmd_src := exec.Command("curl", "-L", goproxy_src_url, "-o", "source.tar.gz")
	err = goproxy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	goproxy_cmd_direct := exec.Command("./binary")
	err = goproxy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
