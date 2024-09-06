package main

// Bartib - Simple timetracker for the command-line
// Homepage: https://github.com/nikolassv/bartib

import (
	"fmt"
	
	"os/exec"
)

func installBartib() {
	// Método 1: Descargar y extraer .tar.gz
	bartib_tar_url := "https://github.com/nikolassv/bartib/archive/refs/tags/v1.1.0.tar.gz"
	bartib_cmd_tar := exec.Command("curl", "-L", bartib_tar_url, "-o", "package.tar.gz")
	err := bartib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bartib_zip_url := "https://github.com/nikolassv/bartib/archive/refs/tags/v1.1.0.zip"
	bartib_cmd_zip := exec.Command("curl", "-L", bartib_zip_url, "-o", "package.zip")
	err = bartib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bartib_bin_url := "https://github.com/nikolassv/bartib/archive/refs/tags/v1.1.0.bin"
	bartib_cmd_bin := exec.Command("curl", "-L", bartib_bin_url, "-o", "binary.bin")
	err = bartib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bartib_src_url := "https://github.com/nikolassv/bartib/archive/refs/tags/v1.1.0.src.tar.gz"
	bartib_cmd_src := exec.Command("curl", "-L", bartib_src_url, "-o", "source.tar.gz")
	err = bartib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bartib_cmd_direct := exec.Command("./binary")
	err = bartib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
