package main

// Leaps - Collaborative web-based text editing service written in Golang
// Homepage: https://github.com/jeffail/leaps

import (
	"fmt"
	
	"os/exec"
)

func installLeaps() {
	// Método 1: Descargar y extraer .tar.gz
	leaps_tar_url := "https://github.com/Jeffail/leaps/archive/refs/tags/v0.9.1.tar.gz"
	leaps_cmd_tar := exec.Command("curl", "-L", leaps_tar_url, "-o", "package.tar.gz")
	err := leaps_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	leaps_zip_url := "https://github.com/Jeffail/leaps/archive/refs/tags/v0.9.1.zip"
	leaps_cmd_zip := exec.Command("curl", "-L", leaps_zip_url, "-o", "package.zip")
	err = leaps_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	leaps_bin_url := "https://github.com/Jeffail/leaps/archive/refs/tags/v0.9.1.bin"
	leaps_cmd_bin := exec.Command("curl", "-L", leaps_bin_url, "-o", "binary.bin")
	err = leaps_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	leaps_src_url := "https://github.com/Jeffail/leaps/archive/refs/tags/v0.9.1.src.tar.gz"
	leaps_cmd_src := exec.Command("curl", "-L", leaps_src_url, "-o", "source.tar.gz")
	err = leaps_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	leaps_cmd_direct := exec.Command("./binary")
	err = leaps_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
