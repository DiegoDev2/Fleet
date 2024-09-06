package main

// Massren - Easily rename multiple files using your text editor
// Homepage: https://github.com/laurent22/massren

import (
	"fmt"
	
	"os/exec"
)

func installMassren() {
	// Método 1: Descargar y extraer .tar.gz
	massren_tar_url := "https://github.com/laurent22/massren/archive/refs/tags/v1.5.6.tar.gz"
	massren_cmd_tar := exec.Command("curl", "-L", massren_tar_url, "-o", "package.tar.gz")
	err := massren_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	massren_zip_url := "https://github.com/laurent22/massren/archive/refs/tags/v1.5.6.zip"
	massren_cmd_zip := exec.Command("curl", "-L", massren_zip_url, "-o", "package.zip")
	err = massren_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	massren_bin_url := "https://github.com/laurent22/massren/archive/refs/tags/v1.5.6.bin"
	massren_cmd_bin := exec.Command("curl", "-L", massren_bin_url, "-o", "binary.bin")
	err = massren_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	massren_src_url := "https://github.com/laurent22/massren/archive/refs/tags/v1.5.6.src.tar.gz"
	massren_cmd_src := exec.Command("curl", "-L", massren_src_url, "-o", "source.tar.gz")
	err = massren_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	massren_cmd_direct := exec.Command("./binary")
	err = massren_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
