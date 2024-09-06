package main

// Redpen - Proofreading tool to help writers of technical documentation
// Homepage: https://redpen.cc/

import (
	"fmt"
	
	"os/exec"
)

func installRedpen() {
	// Método 1: Descargar y extraer .tar.gz
	redpen_tar_url := "https://github.com/redpen-cc/redpen/releases/download/redpen-1.10.4/redpen-1.10.4.tar.gz"
	redpen_cmd_tar := exec.Command("curl", "-L", redpen_tar_url, "-o", "package.tar.gz")
	err := redpen_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	redpen_zip_url := "https://github.com/redpen-cc/redpen/releases/download/redpen-1.10.4/redpen-1.10.4.zip"
	redpen_cmd_zip := exec.Command("curl", "-L", redpen_zip_url, "-o", "package.zip")
	err = redpen_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	redpen_bin_url := "https://github.com/redpen-cc/redpen/releases/download/redpen-1.10.4/redpen-1.10.4.bin"
	redpen_cmd_bin := exec.Command("curl", "-L", redpen_bin_url, "-o", "binary.bin")
	err = redpen_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	redpen_src_url := "https://github.com/redpen-cc/redpen/releases/download/redpen-1.10.4/redpen-1.10.4.src.tar.gz"
	redpen_cmd_src := exec.Command("curl", "-L", redpen_src_url, "-o", "source.tar.gz")
	err = redpen_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	redpen_cmd_direct := exec.Command("./binary")
	err = redpen_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@11")
exec.Command("latte", "install", "openjdk@11")
}
