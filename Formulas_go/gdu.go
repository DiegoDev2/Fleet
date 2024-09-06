package main

// Gdu - Disk usage analyzer with console interface written in Go
// Homepage: https://github.com/dundee/gdu

import (
	"fmt"
	
	"os/exec"
)

func installGdu() {
	// Método 1: Descargar y extraer .tar.gz
	gdu_tar_url := "https://github.com/dundee/gdu/archive/refs/tags/v5.29.0.tar.gz"
	gdu_cmd_tar := exec.Command("curl", "-L", gdu_tar_url, "-o", "package.tar.gz")
	err := gdu_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gdu_zip_url := "https://github.com/dundee/gdu/archive/refs/tags/v5.29.0.zip"
	gdu_cmd_zip := exec.Command("curl", "-L", gdu_zip_url, "-o", "package.zip")
	err = gdu_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gdu_bin_url := "https://github.com/dundee/gdu/archive/refs/tags/v5.29.0.bin"
	gdu_cmd_bin := exec.Command("curl", "-L", gdu_bin_url, "-o", "binary.bin")
	err = gdu_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gdu_src_url := "https://github.com/dundee/gdu/archive/refs/tags/v5.29.0.src.tar.gz"
	gdu_cmd_src := exec.Command("curl", "-L", gdu_src_url, "-o", "source.tar.gz")
	err = gdu_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gdu_cmd_direct := exec.Command("./binary")
	err = gdu_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
