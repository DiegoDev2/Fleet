package main

// Azion - CLI for the Azion service
// Homepage: https://github.com/aziontech/azion

import (
	"fmt"
	
	"os/exec"
)

func installAzion() {
	// Método 1: Descargar y extraer .tar.gz
	azion_tar_url := "https://github.com/aziontech/azion/archive/refs/tags/1.37.0.tar.gz"
	azion_cmd_tar := exec.Command("curl", "-L", azion_tar_url, "-o", "package.tar.gz")
	err := azion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	azion_zip_url := "https://github.com/aziontech/azion/archive/refs/tags/1.37.0.zip"
	azion_cmd_zip := exec.Command("curl", "-L", azion_zip_url, "-o", "package.zip")
	err = azion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	azion_bin_url := "https://github.com/aziontech/azion/archive/refs/tags/1.37.0.bin"
	azion_cmd_bin := exec.Command("curl", "-L", azion_bin_url, "-o", "binary.bin")
	err = azion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	azion_src_url := "https://github.com/aziontech/azion/archive/refs/tags/1.37.0.src.tar.gz"
	azion_cmd_src := exec.Command("curl", "-L", azion_src_url, "-o", "source.tar.gz")
	err = azion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	azion_cmd_direct := exec.Command("./binary")
	err = azion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
