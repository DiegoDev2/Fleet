package main

// Falco - VCL parser and linter optimized for Fastly
// Homepage: https://github.com/ysugimoto/falco

import (
	"fmt"
	
	"os/exec"
)

func installFalco() {
	// Método 1: Descargar y extraer .tar.gz
	falco_tar_url := "https://github.com/ysugimoto/falco/archive/refs/tags/v1.9.1.tar.gz"
	falco_cmd_tar := exec.Command("curl", "-L", falco_tar_url, "-o", "package.tar.gz")
	err := falco_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	falco_zip_url := "https://github.com/ysugimoto/falco/archive/refs/tags/v1.9.1.zip"
	falco_cmd_zip := exec.Command("curl", "-L", falco_zip_url, "-o", "package.zip")
	err = falco_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	falco_bin_url := "https://github.com/ysugimoto/falco/archive/refs/tags/v1.9.1.bin"
	falco_cmd_bin := exec.Command("curl", "-L", falco_bin_url, "-o", "binary.bin")
	err = falco_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	falco_src_url := "https://github.com/ysugimoto/falco/archive/refs/tags/v1.9.1.src.tar.gz"
	falco_cmd_src := exec.Command("curl", "-L", falco_src_url, "-o", "source.tar.gz")
	err = falco_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	falco_cmd_direct := exec.Command("./binary")
	err = falco_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
