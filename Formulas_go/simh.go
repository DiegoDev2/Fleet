package main

// Simh - Portable, multi-system simulator
// Homepage: http://simh.trailing-edge.com/

import (
	"fmt"
	
	"os/exec"
)

func installSimh() {
	// Método 1: Descargar y extraer .tar.gz
	simh_tar_url := "https://github.com/simh/simh/archive/refs/tags/v3.12-2.tar.gz"
	simh_cmd_tar := exec.Command("curl", "-L", simh_tar_url, "-o", "package.tar.gz")
	err := simh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	simh_zip_url := "https://github.com/simh/simh/archive/refs/tags/v3.12-2.zip"
	simh_cmd_zip := exec.Command("curl", "-L", simh_zip_url, "-o", "package.zip")
	err = simh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	simh_bin_url := "https://github.com/simh/simh/archive/refs/tags/v3.12-2.bin"
	simh_cmd_bin := exec.Command("curl", "-L", simh_bin_url, "-o", "binary.bin")
	err = simh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	simh_src_url := "https://github.com/simh/simh/archive/refs/tags/v3.12-2.src.tar.gz"
	simh_cmd_src := exec.Command("curl", "-L", simh_src_url, "-o", "source.tar.gz")
	err = simh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	simh_cmd_direct := exec.Command("./binary")
	err = simh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
}
