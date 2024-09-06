package main

// Mycorrhiza - Lightweight wiki engine with hierarchy support
// Homepage: https://mycorrhiza.wiki

import (
	"fmt"
	
	"os/exec"
)

func installMycorrhiza() {
	// Método 1: Descargar y extraer .tar.gz
	mycorrhiza_tar_url := "https://github.com/bouncepaw/mycorrhiza/archive/refs/tags/v1.15.0.tar.gz"
	mycorrhiza_cmd_tar := exec.Command("curl", "-L", mycorrhiza_tar_url, "-o", "package.tar.gz")
	err := mycorrhiza_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mycorrhiza_zip_url := "https://github.com/bouncepaw/mycorrhiza/archive/refs/tags/v1.15.0.zip"
	mycorrhiza_cmd_zip := exec.Command("curl", "-L", mycorrhiza_zip_url, "-o", "package.zip")
	err = mycorrhiza_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mycorrhiza_bin_url := "https://github.com/bouncepaw/mycorrhiza/archive/refs/tags/v1.15.0.bin"
	mycorrhiza_cmd_bin := exec.Command("curl", "-L", mycorrhiza_bin_url, "-o", "binary.bin")
	err = mycorrhiza_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mycorrhiza_src_url := "https://github.com/bouncepaw/mycorrhiza/archive/refs/tags/v1.15.0.src.tar.gz"
	mycorrhiza_cmd_src := exec.Command("curl", "-L", mycorrhiza_src_url, "-o", "source.tar.gz")
	err = mycorrhiza_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mycorrhiza_cmd_direct := exec.Command("./binary")
	err = mycorrhiza_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
