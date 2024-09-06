package main

// Licensor - Write licenses to stdout
// Homepage: https://github.com/raftario/licensor

import (
	"fmt"
	
	"os/exec"
)

func installLicensor() {
	// Método 1: Descargar y extraer .tar.gz
	licensor_tar_url := "https://github.com/raftario/licensor/archive/refs/tags/v2.1.0.tar.gz"
	licensor_cmd_tar := exec.Command("curl", "-L", licensor_tar_url, "-o", "package.tar.gz")
	err := licensor_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	licensor_zip_url := "https://github.com/raftario/licensor/archive/refs/tags/v2.1.0.zip"
	licensor_cmd_zip := exec.Command("curl", "-L", licensor_zip_url, "-o", "package.zip")
	err = licensor_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	licensor_bin_url := "https://github.com/raftario/licensor/archive/refs/tags/v2.1.0.bin"
	licensor_cmd_bin := exec.Command("curl", "-L", licensor_bin_url, "-o", "binary.bin")
	err = licensor_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	licensor_src_url := "https://github.com/raftario/licensor/archive/refs/tags/v2.1.0.src.tar.gz"
	licensor_cmd_src := exec.Command("curl", "-L", licensor_src_url, "-o", "source.tar.gz")
	err = licensor_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	licensor_cmd_direct := exec.Command("./binary")
	err = licensor_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
