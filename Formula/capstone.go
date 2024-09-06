package main

// Capstone - Multi-platform, multi-architecture disassembly framework
// Homepage: https://www.capstone-engine.org/

import (
	"fmt"
	
	"os/exec"
)

func installCapstone() {
	// Método 1: Descargar y extraer .tar.gz
	capstone_tar_url := "https://github.com/capstone-engine/capstone/archive/refs/tags/5.0.3.tar.gz"
	capstone_cmd_tar := exec.Command("curl", "-L", capstone_tar_url, "-o", "package.tar.gz")
	err := capstone_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	capstone_zip_url := "https://github.com/capstone-engine/capstone/archive/refs/tags/5.0.3.zip"
	capstone_cmd_zip := exec.Command("curl", "-L", capstone_zip_url, "-o", "package.zip")
	err = capstone_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	capstone_bin_url := "https://github.com/capstone-engine/capstone/archive/refs/tags/5.0.3.bin"
	capstone_cmd_bin := exec.Command("curl", "-L", capstone_bin_url, "-o", "binary.bin")
	err = capstone_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	capstone_src_url := "https://github.com/capstone-engine/capstone/archive/refs/tags/5.0.3.src.tar.gz"
	capstone_cmd_src := exec.Command("curl", "-L", capstone_src_url, "-o", "source.tar.gz")
	err = capstone_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	capstone_cmd_direct := exec.Command("./binary")
	err = capstone_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
