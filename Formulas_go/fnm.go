package main

// Fnm - Fast and simple Node.js version manager
// Homepage: https://github.com/Schniz/fnm

import (
	"fmt"
	
	"os/exec"
)

func installFnm() {
	// Método 1: Descargar y extraer .tar.gz
	fnm_tar_url := "https://github.com/Schniz/fnm/archive/refs/tags/v1.37.1.tar.gz"
	fnm_cmd_tar := exec.Command("curl", "-L", fnm_tar_url, "-o", "package.tar.gz")
	err := fnm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fnm_zip_url := "https://github.com/Schniz/fnm/archive/refs/tags/v1.37.1.zip"
	fnm_cmd_zip := exec.Command("curl", "-L", fnm_zip_url, "-o", "package.zip")
	err = fnm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fnm_bin_url := "https://github.com/Schniz/fnm/archive/refs/tags/v1.37.1.bin"
	fnm_cmd_bin := exec.Command("curl", "-L", fnm_bin_url, "-o", "binary.bin")
	err = fnm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fnm_src_url := "https://github.com/Schniz/fnm/archive/refs/tags/v1.37.1.src.tar.gz"
	fnm_cmd_src := exec.Command("curl", "-L", fnm_src_url, "-o", "source.tar.gz")
	err = fnm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fnm_cmd_direct := exec.Command("./binary")
	err = fnm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
