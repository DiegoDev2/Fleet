package main

// Osqp - Operator splitting QP solver
// Homepage: https://osqp.org/

import (
	"fmt"
	
	"os/exec"
)

func installOsqp() {
	// Método 1: Descargar y extraer .tar.gz
	osqp_tar_url := "https://github.com/osqp/osqp/archive/refs/tags/v0.6.3.tar.gz"
	osqp_cmd_tar := exec.Command("curl", "-L", osqp_tar_url, "-o", "package.tar.gz")
	err := osqp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	osqp_zip_url := "https://github.com/osqp/osqp/archive/refs/tags/v0.6.3.zip"
	osqp_cmd_zip := exec.Command("curl", "-L", osqp_zip_url, "-o", "package.zip")
	err = osqp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	osqp_bin_url := "https://github.com/osqp/osqp/archive/refs/tags/v0.6.3.bin"
	osqp_cmd_bin := exec.Command("curl", "-L", osqp_bin_url, "-o", "binary.bin")
	err = osqp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	osqp_src_url := "https://github.com/osqp/osqp/archive/refs/tags/v0.6.3.src.tar.gz"
	osqp_cmd_src := exec.Command("curl", "-L", osqp_src_url, "-o", "source.tar.gz")
	err = osqp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	osqp_cmd_direct := exec.Command("./binary")
	err = osqp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
