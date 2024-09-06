package main

// Gat - Cat alternative written in Go
// Homepage: https://github.com/koki-develop/gat

import (
	"fmt"
	
	"os/exec"
)

func installGat() {
	// Método 1: Descargar y extraer .tar.gz
	gat_tar_url := "https://github.com/koki-develop/gat/archive/refs/tags/v0.18.0.tar.gz"
	gat_cmd_tar := exec.Command("curl", "-L", gat_tar_url, "-o", "package.tar.gz")
	err := gat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gat_zip_url := "https://github.com/koki-develop/gat/archive/refs/tags/v0.18.0.zip"
	gat_cmd_zip := exec.Command("curl", "-L", gat_zip_url, "-o", "package.zip")
	err = gat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gat_bin_url := "https://github.com/koki-develop/gat/archive/refs/tags/v0.18.0.bin"
	gat_cmd_bin := exec.Command("curl", "-L", gat_bin_url, "-o", "binary.bin")
	err = gat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gat_src_url := "https://github.com/koki-develop/gat/archive/refs/tags/v0.18.0.src.tar.gz"
	gat_cmd_src := exec.Command("curl", "-L", gat_src_url, "-o", "source.tar.gz")
	err = gat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gat_cmd_direct := exec.Command("./binary")
	err = gat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
