package main

// Cassowary - Modern cross-platform HTTP load-testing tool written in Go
// Homepage: https://github.com/rogerwelin/cassowary

import (
	"fmt"
	
	"os/exec"
)

func installCassowary() {
	// Método 1: Descargar y extraer .tar.gz
	cassowary_tar_url := "https://github.com/rogerwelin/cassowary/archive/refs/tags/v0.17.0.tar.gz"
	cassowary_cmd_tar := exec.Command("curl", "-L", cassowary_tar_url, "-o", "package.tar.gz")
	err := cassowary_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cassowary_zip_url := "https://github.com/rogerwelin/cassowary/archive/refs/tags/v0.17.0.zip"
	cassowary_cmd_zip := exec.Command("curl", "-L", cassowary_zip_url, "-o", "package.zip")
	err = cassowary_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cassowary_bin_url := "https://github.com/rogerwelin/cassowary/archive/refs/tags/v0.17.0.bin"
	cassowary_cmd_bin := exec.Command("curl", "-L", cassowary_bin_url, "-o", "binary.bin")
	err = cassowary_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cassowary_src_url := "https://github.com/rogerwelin/cassowary/archive/refs/tags/v0.17.0.src.tar.gz"
	cassowary_cmd_src := exec.Command("curl", "-L", cassowary_src_url, "-o", "source.tar.gz")
	err = cassowary_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cassowary_cmd_direct := exec.Command("./binary")
	err = cassowary_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
