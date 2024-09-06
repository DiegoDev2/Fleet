package main

// Ccheck - Check X509 certificate expiration from the command-line, with TAP output
// Homepage: https://github.com/nerdlem/ccheck

import (
	"fmt"
	
	"os/exec"
)

func installCcheck() {
	// Método 1: Descargar y extraer .tar.gz
	ccheck_tar_url := "https://github.com/nerdlem/ccheck/archive/refs/tags/v1.0.1.tar.gz"
	ccheck_cmd_tar := exec.Command("curl", "-L", ccheck_tar_url, "-o", "package.tar.gz")
	err := ccheck_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ccheck_zip_url := "https://github.com/nerdlem/ccheck/archive/refs/tags/v1.0.1.zip"
	ccheck_cmd_zip := exec.Command("curl", "-L", ccheck_zip_url, "-o", "package.zip")
	err = ccheck_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ccheck_bin_url := "https://github.com/nerdlem/ccheck/archive/refs/tags/v1.0.1.bin"
	ccheck_cmd_bin := exec.Command("curl", "-L", ccheck_bin_url, "-o", "binary.bin")
	err = ccheck_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ccheck_src_url := "https://github.com/nerdlem/ccheck/archive/refs/tags/v1.0.1.src.tar.gz"
	ccheck_cmd_src := exec.Command("curl", "-L", ccheck_src_url, "-o", "source.tar.gz")
	err = ccheck_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ccheck_cmd_direct := exec.Command("./binary")
	err = ccheck_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
