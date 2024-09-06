package main

// Nlopt - Free/open-source library for nonlinear optimization
// Homepage: https://nlopt.readthedocs.io/

import (
	"fmt"
	
	"os/exec"
)

func installNlopt() {
	// Método 1: Descargar y extraer .tar.gz
	nlopt_tar_url := "https://github.com/stevengj/nlopt/archive/refs/tags/v2.8.0.tar.gz"
	nlopt_cmd_tar := exec.Command("curl", "-L", nlopt_tar_url, "-o", "package.tar.gz")
	err := nlopt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nlopt_zip_url := "https://github.com/stevengj/nlopt/archive/refs/tags/v2.8.0.zip"
	nlopt_cmd_zip := exec.Command("curl", "-L", nlopt_zip_url, "-o", "package.zip")
	err = nlopt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nlopt_bin_url := "https://github.com/stevengj/nlopt/archive/refs/tags/v2.8.0.bin"
	nlopt_cmd_bin := exec.Command("curl", "-L", nlopt_bin_url, "-o", "binary.bin")
	err = nlopt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nlopt_src_url := "https://github.com/stevengj/nlopt/archive/refs/tags/v2.8.0.src.tar.gz"
	nlopt_cmd_src := exec.Command("curl", "-L", nlopt_src_url, "-o", "source.tar.gz")
	err = nlopt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nlopt_cmd_direct := exec.Command("./binary")
	err = nlopt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
