package main

// Cdebug - Swiss army knife of container debugging
// Homepage: https://github.com/iximiuz/cdebug

import (
	"fmt"
	
	"os/exec"
)

func installCdebug() {
	// Método 1: Descargar y extraer .tar.gz
	cdebug_tar_url := "https://github.com/iximiuz/cdebug/archive/refs/tags/v0.0.18.tar.gz"
	cdebug_cmd_tar := exec.Command("curl", "-L", cdebug_tar_url, "-o", "package.tar.gz")
	err := cdebug_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cdebug_zip_url := "https://github.com/iximiuz/cdebug/archive/refs/tags/v0.0.18.zip"
	cdebug_cmd_zip := exec.Command("curl", "-L", cdebug_zip_url, "-o", "package.zip")
	err = cdebug_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cdebug_bin_url := "https://github.com/iximiuz/cdebug/archive/refs/tags/v0.0.18.bin"
	cdebug_cmd_bin := exec.Command("curl", "-L", cdebug_bin_url, "-o", "binary.bin")
	err = cdebug_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cdebug_src_url := "https://github.com/iximiuz/cdebug/archive/refs/tags/v0.0.18.src.tar.gz"
	cdebug_cmd_src := exec.Command("curl", "-L", cdebug_src_url, "-o", "source.tar.gz")
	err = cdebug_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cdebug_cmd_direct := exec.Command("./binary")
	err = cdebug_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
