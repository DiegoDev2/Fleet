package main

// Rgf - Regularized Greedy Forest library
// Homepage: https://github.com/RGF-team/rgf

import (
	"fmt"
	
	"os/exec"
)

func installRgf() {
	// Método 1: Descargar y extraer .tar.gz
	rgf_tar_url := "https://github.com/RGF-team/rgf/archive/refs/tags/3.12.0.tar.gz"
	rgf_cmd_tar := exec.Command("curl", "-L", rgf_tar_url, "-o", "package.tar.gz")
	err := rgf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rgf_zip_url := "https://github.com/RGF-team/rgf/archive/refs/tags/3.12.0.zip"
	rgf_cmd_zip := exec.Command("curl", "-L", rgf_zip_url, "-o", "package.zip")
	err = rgf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rgf_bin_url := "https://github.com/RGF-team/rgf/archive/refs/tags/3.12.0.bin"
	rgf_cmd_bin := exec.Command("curl", "-L", rgf_bin_url, "-o", "binary.bin")
	err = rgf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rgf_src_url := "https://github.com/RGF-team/rgf/archive/refs/tags/3.12.0.src.tar.gz"
	rgf_cmd_src := exec.Command("curl", "-L", rgf_src_url, "-o", "source.tar.gz")
	err = rgf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rgf_cmd_direct := exec.Command("./binary")
	err = rgf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
