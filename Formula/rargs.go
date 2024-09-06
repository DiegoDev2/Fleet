package main

// Rargs - Util like xargs + awk with pattern matching support
// Homepage: https://github.com/lotabout/rargs

import (
	"fmt"
	
	"os/exec"
)

func installRargs() {
	// Método 1: Descargar y extraer .tar.gz
	rargs_tar_url := "https://github.com/lotabout/rargs/archive/refs/tags/v0.3.0.tar.gz"
	rargs_cmd_tar := exec.Command("curl", "-L", rargs_tar_url, "-o", "package.tar.gz")
	err := rargs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rargs_zip_url := "https://github.com/lotabout/rargs/archive/refs/tags/v0.3.0.zip"
	rargs_cmd_zip := exec.Command("curl", "-L", rargs_zip_url, "-o", "package.zip")
	err = rargs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rargs_bin_url := "https://github.com/lotabout/rargs/archive/refs/tags/v0.3.0.bin"
	rargs_cmd_bin := exec.Command("curl", "-L", rargs_bin_url, "-o", "binary.bin")
	err = rargs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rargs_src_url := "https://github.com/lotabout/rargs/archive/refs/tags/v0.3.0.src.tar.gz"
	rargs_cmd_src := exec.Command("curl", "-L", rargs_src_url, "-o", "source.tar.gz")
	err = rargs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rargs_cmd_direct := exec.Command("./binary")
	err = rargs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
