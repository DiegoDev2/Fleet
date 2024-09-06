package main

// Shfmt - Autoformat shell script source code
// Homepage: https://github.com/mvdan/sh

import (
	"fmt"
	
	"os/exec"
)

func installShfmt() {
	// Método 1: Descargar y extraer .tar.gz
	shfmt_tar_url := "https://github.com/mvdan/sh/archive/refs/tags/v3.9.0.tar.gz"
	shfmt_cmd_tar := exec.Command("curl", "-L", shfmt_tar_url, "-o", "package.tar.gz")
	err := shfmt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	shfmt_zip_url := "https://github.com/mvdan/sh/archive/refs/tags/v3.9.0.zip"
	shfmt_cmd_zip := exec.Command("curl", "-L", shfmt_zip_url, "-o", "package.zip")
	err = shfmt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	shfmt_bin_url := "https://github.com/mvdan/sh/archive/refs/tags/v3.9.0.bin"
	shfmt_cmd_bin := exec.Command("curl", "-L", shfmt_bin_url, "-o", "binary.bin")
	err = shfmt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	shfmt_src_url := "https://github.com/mvdan/sh/archive/refs/tags/v3.9.0.src.tar.gz"
	shfmt_cmd_src := exec.Command("curl", "-L", shfmt_src_url, "-o", "source.tar.gz")
	err = shfmt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	shfmt_cmd_direct := exec.Command("./binary")
	err = shfmt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: scdoc")
	exec.Command("latte", "install", "scdoc").Run()
}
