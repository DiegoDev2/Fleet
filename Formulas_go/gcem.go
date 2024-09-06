package main

// Gcem - C++ compile-time math library
// Homepage: https://gcem.readthedocs.io/en/latest/

import (
	"fmt"
	
	"os/exec"
)

func installGcem() {
	// Método 1: Descargar y extraer .tar.gz
	gcem_tar_url := "https://github.com/kthohr/gcem/archive/refs/tags/v1.18.0.tar.gz"
	gcem_cmd_tar := exec.Command("curl", "-L", gcem_tar_url, "-o", "package.tar.gz")
	err := gcem_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gcem_zip_url := "https://github.com/kthohr/gcem/archive/refs/tags/v1.18.0.zip"
	gcem_cmd_zip := exec.Command("curl", "-L", gcem_zip_url, "-o", "package.zip")
	err = gcem_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gcem_bin_url := "https://github.com/kthohr/gcem/archive/refs/tags/v1.18.0.bin"
	gcem_cmd_bin := exec.Command("curl", "-L", gcem_bin_url, "-o", "binary.bin")
	err = gcem_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gcem_src_url := "https://github.com/kthohr/gcem/archive/refs/tags/v1.18.0.src.tar.gz"
	gcem_cmd_src := exec.Command("curl", "-L", gcem_src_url, "-o", "source.tar.gz")
	err = gcem_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gcem_cmd_direct := exec.Command("./binary")
	err = gcem_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
