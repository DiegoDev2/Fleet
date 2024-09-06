package main

// GdriveAT2 - Google Drive CLI Client
// Homepage: https://github.com/prasmussen/gdrive

import (
	"fmt"
	
	"os/exec"
)

func installGdriveAT2() {
	// Método 1: Descargar y extraer .tar.gz
	gdriveat2_tar_url := "https://github.com/prasmussen/gdrive/archive/refs/tags/2.1.1.tar.gz"
	gdriveat2_cmd_tar := exec.Command("curl", "-L", gdriveat2_tar_url, "-o", "package.tar.gz")
	err := gdriveat2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gdriveat2_zip_url := "https://github.com/prasmussen/gdrive/archive/refs/tags/2.1.1.zip"
	gdriveat2_cmd_zip := exec.Command("curl", "-L", gdriveat2_zip_url, "-o", "package.zip")
	err = gdriveat2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gdriveat2_bin_url := "https://github.com/prasmussen/gdrive/archive/refs/tags/2.1.1.bin"
	gdriveat2_cmd_bin := exec.Command("curl", "-L", gdriveat2_bin_url, "-o", "binary.bin")
	err = gdriveat2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gdriveat2_src_url := "https://github.com/prasmussen/gdrive/archive/refs/tags/2.1.1.src.tar.gz"
	gdriveat2_cmd_src := exec.Command("curl", "-L", gdriveat2_src_url, "-o", "source.tar.gz")
	err = gdriveat2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gdriveat2_cmd_direct := exec.Command("./binary")
	err = gdriveat2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
