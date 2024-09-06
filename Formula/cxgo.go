package main

// Cxgo - Transpiling C to Go
// Homepage: https://github.com/gotranspile/cxgo

import (
	"fmt"
	
	"os/exec"
)

func installCxgo() {
	// Método 1: Descargar y extraer .tar.gz
	cxgo_tar_url := "https://github.com/gotranspile/cxgo/archive/refs/tags/v0.4.0.tar.gz"
	cxgo_cmd_tar := exec.Command("curl", "-L", cxgo_tar_url, "-o", "package.tar.gz")
	err := cxgo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cxgo_zip_url := "https://github.com/gotranspile/cxgo/archive/refs/tags/v0.4.0.zip"
	cxgo_cmd_zip := exec.Command("curl", "-L", cxgo_zip_url, "-o", "package.zip")
	err = cxgo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cxgo_bin_url := "https://github.com/gotranspile/cxgo/archive/refs/tags/v0.4.0.bin"
	cxgo_cmd_bin := exec.Command("curl", "-L", cxgo_bin_url, "-o", "binary.bin")
	err = cxgo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cxgo_src_url := "https://github.com/gotranspile/cxgo/archive/refs/tags/v0.4.0.src.tar.gz"
	cxgo_cmd_src := exec.Command("curl", "-L", cxgo_src_url, "-o", "source.tar.gz")
	err = cxgo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cxgo_cmd_direct := exec.Command("./binary")
	err = cxgo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
