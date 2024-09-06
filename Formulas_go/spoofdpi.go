package main

// Spoofdpi - Simple and fast anti-censorship tool written in Go
// Homepage: https://github.com/xvzc/SpoofDPI

import (
	"fmt"
	
	"os/exec"
)

func installSpoofdpi() {
	// Método 1: Descargar y extraer .tar.gz
	spoofdpi_tar_url := "https://github.com/xvzc/SpoofDPI/archive/refs/tags/v0.11.1.tar.gz"
	spoofdpi_cmd_tar := exec.Command("curl", "-L", spoofdpi_tar_url, "-o", "package.tar.gz")
	err := spoofdpi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spoofdpi_zip_url := "https://github.com/xvzc/SpoofDPI/archive/refs/tags/v0.11.1.zip"
	spoofdpi_cmd_zip := exec.Command("curl", "-L", spoofdpi_zip_url, "-o", "package.zip")
	err = spoofdpi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spoofdpi_bin_url := "https://github.com/xvzc/SpoofDPI/archive/refs/tags/v0.11.1.bin"
	spoofdpi_cmd_bin := exec.Command("curl", "-L", spoofdpi_bin_url, "-o", "binary.bin")
	err = spoofdpi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spoofdpi_src_url := "https://github.com/xvzc/SpoofDPI/archive/refs/tags/v0.11.1.src.tar.gz"
	spoofdpi_cmd_src := exec.Command("curl", "-L", spoofdpi_src_url, "-o", "source.tar.gz")
	err = spoofdpi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spoofdpi_cmd_direct := exec.Command("./binary")
	err = spoofdpi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
