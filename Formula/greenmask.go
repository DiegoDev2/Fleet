package main

// Greenmask - PostgreSQL dump and obfuscation tool
// Homepage: https://greenmask.io

import (
	"fmt"
	
	"os/exec"
)

func installGreenmask() {
	// Método 1: Descargar y extraer .tar.gz
	greenmask_tar_url := "https://github.com/GreenmaskIO/greenmask/archive/refs/tags/v0.1.14.tar.gz"
	greenmask_cmd_tar := exec.Command("curl", "-L", greenmask_tar_url, "-o", "package.tar.gz")
	err := greenmask_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	greenmask_zip_url := "https://github.com/GreenmaskIO/greenmask/archive/refs/tags/v0.1.14.zip"
	greenmask_cmd_zip := exec.Command("curl", "-L", greenmask_zip_url, "-o", "package.zip")
	err = greenmask_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	greenmask_bin_url := "https://github.com/GreenmaskIO/greenmask/archive/refs/tags/v0.1.14.bin"
	greenmask_cmd_bin := exec.Command("curl", "-L", greenmask_bin_url, "-o", "binary.bin")
	err = greenmask_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	greenmask_src_url := "https://github.com/GreenmaskIO/greenmask/archive/refs/tags/v0.1.14.src.tar.gz"
	greenmask_cmd_src := exec.Command("curl", "-L", greenmask_src_url, "-o", "source.tar.gz")
	err = greenmask_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	greenmask_cmd_direct := exec.Command("./binary")
	err = greenmask_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
