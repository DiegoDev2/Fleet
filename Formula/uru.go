package main

// Uru - Use multiple rubies on multiple platforms
// Homepage: https://bitbucket.org/jonforums/uru

import (
	"fmt"
	
	"os/exec"
)

func installUru() {
	// Método 1: Descargar y extraer .tar.gz
	uru_tar_url := "https://bitbucket.org/jonforums/uru/get/v0.8.5.tar.gz"
	uru_cmd_tar := exec.Command("curl", "-L", uru_tar_url, "-o", "package.tar.gz")
	err := uru_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	uru_zip_url := "https://bitbucket.org/jonforums/uru/get/v0.8.5.zip"
	uru_cmd_zip := exec.Command("curl", "-L", uru_zip_url, "-o", "package.zip")
	err = uru_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	uru_bin_url := "https://bitbucket.org/jonforums/uru/get/v0.8.5.bin"
	uru_cmd_bin := exec.Command("curl", "-L", uru_bin_url, "-o", "binary.bin")
	err = uru_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	uru_src_url := "https://bitbucket.org/jonforums/uru/get/v0.8.5.src.tar.gz"
	uru_cmd_src := exec.Command("curl", "-L", uru_src_url, "-o", "source.tar.gz")
	err = uru_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	uru_cmd_direct := exec.Command("./binary")
	err = uru_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
