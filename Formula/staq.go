package main

// Staq - Full-stack quantum processing toolkit
// Homepage: https://github.com/softwareQinc/staq

import (
	"fmt"
	
	"os/exec"
)

func installStaq() {
	// Método 1: Descargar y extraer .tar.gz
	staq_tar_url := "https://github.com/softwareQinc/staq/archive/refs/tags/v3.5.tar.gz"
	staq_cmd_tar := exec.Command("curl", "-L", staq_tar_url, "-o", "package.tar.gz")
	err := staq_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	staq_zip_url := "https://github.com/softwareQinc/staq/archive/refs/tags/v3.5.zip"
	staq_cmd_zip := exec.Command("curl", "-L", staq_zip_url, "-o", "package.zip")
	err = staq_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	staq_bin_url := "https://github.com/softwareQinc/staq/archive/refs/tags/v3.5.bin"
	staq_cmd_bin := exec.Command("curl", "-L", staq_bin_url, "-o", "binary.bin")
	err = staq_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	staq_src_url := "https://github.com/softwareQinc/staq/archive/refs/tags/v3.5.src.tar.gz"
	staq_cmd_src := exec.Command("curl", "-L", staq_src_url, "-o", "source.tar.gz")
	err = staq_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	staq_cmd_direct := exec.Command("./binary")
	err = staq_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
