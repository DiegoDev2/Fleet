package main

// FuncE - Easily run Envoy
// Homepage: https://func-e.io

import (
	"fmt"
	
	"os/exec"
)

func installFuncE() {
	// Método 1: Descargar y extraer .tar.gz
	funce_tar_url := "https://github.com/tetratelabs/func-e/archive/refs/tags/v1.1.4.tar.gz"
	funce_cmd_tar := exec.Command("curl", "-L", funce_tar_url, "-o", "package.tar.gz")
	err := funce_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	funce_zip_url := "https://github.com/tetratelabs/func-e/archive/refs/tags/v1.1.4.zip"
	funce_cmd_zip := exec.Command("curl", "-L", funce_zip_url, "-o", "package.zip")
	err = funce_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	funce_bin_url := "https://github.com/tetratelabs/func-e/archive/refs/tags/v1.1.4.bin"
	funce_cmd_bin := exec.Command("curl", "-L", funce_bin_url, "-o", "binary.bin")
	err = funce_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	funce_src_url := "https://github.com/tetratelabs/func-e/archive/refs/tags/v1.1.4.src.tar.gz"
	funce_cmd_src := exec.Command("curl", "-L", funce_src_url, "-o", "source.tar.gz")
	err = funce_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	funce_cmd_direct := exec.Command("./binary")
	err = funce_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
