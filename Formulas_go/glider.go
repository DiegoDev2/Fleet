package main

// Glider - Forward proxy with multiple protocols support
// Homepage: https://github.com/nadoo/glider

import (
	"fmt"
	
	"os/exec"
)

func installGlider() {
	// Método 1: Descargar y extraer .tar.gz
	glider_tar_url := "https://github.com/nadoo/glider/archive/refs/tags/v0.16.4.tar.gz"
	glider_cmd_tar := exec.Command("curl", "-L", glider_tar_url, "-o", "package.tar.gz")
	err := glider_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	glider_zip_url := "https://github.com/nadoo/glider/archive/refs/tags/v0.16.4.zip"
	glider_cmd_zip := exec.Command("curl", "-L", glider_zip_url, "-o", "package.zip")
	err = glider_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	glider_bin_url := "https://github.com/nadoo/glider/archive/refs/tags/v0.16.4.bin"
	glider_cmd_bin := exec.Command("curl", "-L", glider_bin_url, "-o", "binary.bin")
	err = glider_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	glider_src_url := "https://github.com/nadoo/glider/archive/refs/tags/v0.16.4.src.tar.gz"
	glider_cmd_src := exec.Command("curl", "-L", glider_src_url, "-o", "source.tar.gz")
	err = glider_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	glider_cmd_direct := exec.Command("./binary")
	err = glider_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
