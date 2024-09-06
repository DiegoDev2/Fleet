package main

// GoStatik - Embed files into a Go executable
// Homepage: https://github.com/rakyll/statik

import (
	"fmt"
	
	"os/exec"
)

func installGoStatik() {
	// Método 1: Descargar y extraer .tar.gz
	gostatik_tar_url := "https://github.com/rakyll/statik/archive/refs/tags/v0.1.7.tar.gz"
	gostatik_cmd_tar := exec.Command("curl", "-L", gostatik_tar_url, "-o", "package.tar.gz")
	err := gostatik_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gostatik_zip_url := "https://github.com/rakyll/statik/archive/refs/tags/v0.1.7.zip"
	gostatik_cmd_zip := exec.Command("curl", "-L", gostatik_zip_url, "-o", "package.zip")
	err = gostatik_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gostatik_bin_url := "https://github.com/rakyll/statik/archive/refs/tags/v0.1.7.bin"
	gostatik_cmd_bin := exec.Command("curl", "-L", gostatik_bin_url, "-o", "binary.bin")
	err = gostatik_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gostatik_src_url := "https://github.com/rakyll/statik/archive/refs/tags/v0.1.7.src.tar.gz"
	gostatik_cmd_src := exec.Command("curl", "-L", gostatik_src_url, "-o", "source.tar.gz")
	err = gostatik_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gostatik_cmd_direct := exec.Command("./binary")
	err = gostatik_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
