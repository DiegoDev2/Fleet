package main

// Changie - Automated changelog tool for preparing releases
// Homepage: https://changie.dev/

import (
	"fmt"
	
	"os/exec"
)

func installChangie() {
	// Método 1: Descargar y extraer .tar.gz
	changie_tar_url := "https://github.com/miniscruff/changie/archive/refs/tags/v1.19.1.tar.gz"
	changie_cmd_tar := exec.Command("curl", "-L", changie_tar_url, "-o", "package.tar.gz")
	err := changie_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	changie_zip_url := "https://github.com/miniscruff/changie/archive/refs/tags/v1.19.1.zip"
	changie_cmd_zip := exec.Command("curl", "-L", changie_zip_url, "-o", "package.zip")
	err = changie_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	changie_bin_url := "https://github.com/miniscruff/changie/archive/refs/tags/v1.19.1.bin"
	changie_cmd_bin := exec.Command("curl", "-L", changie_bin_url, "-o", "binary.bin")
	err = changie_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	changie_src_url := "https://github.com/miniscruff/changie/archive/refs/tags/v1.19.1.src.tar.gz"
	changie_cmd_src := exec.Command("curl", "-L", changie_src_url, "-o", "source.tar.gz")
	err = changie_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	changie_cmd_direct := exec.Command("./binary")
	err = changie_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
