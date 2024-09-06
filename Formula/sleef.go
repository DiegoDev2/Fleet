package main

// Sleef - SIMD library for evaluating elementary functions
// Homepage: https://sleef.org

import (
	"fmt"
	
	"os/exec"
)

func installSleef() {
	// Método 1: Descargar y extraer .tar.gz
	sleef_tar_url := "https://github.com/shibatch/sleef/archive/refs/tags/3.6.1.tar.gz"
	sleef_cmd_tar := exec.Command("curl", "-L", sleef_tar_url, "-o", "package.tar.gz")
	err := sleef_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sleef_zip_url := "https://github.com/shibatch/sleef/archive/refs/tags/3.6.1.zip"
	sleef_cmd_zip := exec.Command("curl", "-L", sleef_zip_url, "-o", "package.zip")
	err = sleef_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sleef_bin_url := "https://github.com/shibatch/sleef/archive/refs/tags/3.6.1.bin"
	sleef_cmd_bin := exec.Command("curl", "-L", sleef_bin_url, "-o", "binary.bin")
	err = sleef_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sleef_src_url := "https://github.com/shibatch/sleef/archive/refs/tags/3.6.1.src.tar.gz"
	sleef_cmd_src := exec.Command("curl", "-L", sleef_src_url, "-o", "source.tar.gz")
	err = sleef_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sleef_cmd_direct := exec.Command("./binary")
	err = sleef_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
