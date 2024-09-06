package main

// Oslo - CLI tool for the OpenSLO spec
// Homepage: https://openslo.com/

import (
	"fmt"
	
	"os/exec"
)

func installOslo() {
	// Método 1: Descargar y extraer .tar.gz
	oslo_tar_url := "https://github.com/OpenSLO/oslo/archive/refs/tags/v0.12.0.tar.gz"
	oslo_cmd_tar := exec.Command("curl", "-L", oslo_tar_url, "-o", "package.tar.gz")
	err := oslo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	oslo_zip_url := "https://github.com/OpenSLO/oslo/archive/refs/tags/v0.12.0.zip"
	oslo_cmd_zip := exec.Command("curl", "-L", oslo_zip_url, "-o", "package.zip")
	err = oslo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	oslo_bin_url := "https://github.com/OpenSLO/oslo/archive/refs/tags/v0.12.0.bin"
	oslo_cmd_bin := exec.Command("curl", "-L", oslo_bin_url, "-o", "binary.bin")
	err = oslo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	oslo_src_url := "https://github.com/OpenSLO/oslo/archive/refs/tags/v0.12.0.src.tar.gz"
	oslo_cmd_src := exec.Command("curl", "-L", oslo_src_url, "-o", "source.tar.gz")
	err = oslo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	oslo_cmd_direct := exec.Command("./binary")
	err = oslo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
