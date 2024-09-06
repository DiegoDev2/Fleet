package main

// Vaulted - Allows the secure storage and execution of environments
// Homepage: https://github.com/miquella/vaulted

import (
	"fmt"
	
	"os/exec"
)

func installVaulted() {
	// Método 1: Descargar y extraer .tar.gz
	vaulted_tar_url := "https://github.com/miquella/vaulted/archive/refs/tags/v3.0.0.tar.gz"
	vaulted_cmd_tar := exec.Command("curl", "-L", vaulted_tar_url, "-o", "package.tar.gz")
	err := vaulted_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vaulted_zip_url := "https://github.com/miquella/vaulted/archive/refs/tags/v3.0.0.zip"
	vaulted_cmd_zip := exec.Command("curl", "-L", vaulted_zip_url, "-o", "package.zip")
	err = vaulted_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vaulted_bin_url := "https://github.com/miquella/vaulted/archive/refs/tags/v3.0.0.bin"
	vaulted_cmd_bin := exec.Command("curl", "-L", vaulted_bin_url, "-o", "binary.bin")
	err = vaulted_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vaulted_src_url := "https://github.com/miquella/vaulted/archive/refs/tags/v3.0.0.src.tar.gz"
	vaulted_cmd_src := exec.Command("curl", "-L", vaulted_src_url, "-o", "source.tar.gz")
	err = vaulted_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vaulted_cmd_direct := exec.Command("./binary")
	err = vaulted_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
