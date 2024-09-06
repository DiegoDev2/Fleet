package main

// Golines - Golang formatter that fixes long lines
// Homepage: https://github.com/segmentio/golines

import (
	"fmt"
	
	"os/exec"
)

func installGolines() {
	// Método 1: Descargar y extraer .tar.gz
	golines_tar_url := "https://github.com/segmentio/golines/archive/refs/tags/v0.12.2.tar.gz"
	golines_cmd_tar := exec.Command("curl", "-L", golines_tar_url, "-o", "package.tar.gz")
	err := golines_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	golines_zip_url := "https://github.com/segmentio/golines/archive/refs/tags/v0.12.2.zip"
	golines_cmd_zip := exec.Command("curl", "-L", golines_zip_url, "-o", "package.zip")
	err = golines_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	golines_bin_url := "https://github.com/segmentio/golines/archive/refs/tags/v0.12.2.bin"
	golines_cmd_bin := exec.Command("curl", "-L", golines_bin_url, "-o", "binary.bin")
	err = golines_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	golines_src_url := "https://github.com/segmentio/golines/archive/refs/tags/v0.12.2.src.tar.gz"
	golines_cmd_src := exec.Command("curl", "-L", golines_src_url, "-o", "source.tar.gz")
	err = golines_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	golines_cmd_direct := exec.Command("./binary")
	err = golines_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
