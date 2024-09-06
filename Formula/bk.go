package main

// Bk - Terminal EPUB Reader
// Homepage: https://github.com/aeosynth/bk

import (
	"fmt"
	
	"os/exec"
)

func installBk() {
	// Método 1: Descargar y extraer .tar.gz
	bk_tar_url := "https://github.com/aeosynth/bk/archive/refs/tags/v0.6.0.tar.gz"
	bk_cmd_tar := exec.Command("curl", "-L", bk_tar_url, "-o", "package.tar.gz")
	err := bk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bk_zip_url := "https://github.com/aeosynth/bk/archive/refs/tags/v0.6.0.zip"
	bk_cmd_zip := exec.Command("curl", "-L", bk_zip_url, "-o", "package.zip")
	err = bk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bk_bin_url := "https://github.com/aeosynth/bk/archive/refs/tags/v0.6.0.bin"
	bk_cmd_bin := exec.Command("curl", "-L", bk_bin_url, "-o", "binary.bin")
	err = bk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bk_src_url := "https://github.com/aeosynth/bk/archive/refs/tags/v0.6.0.src.tar.gz"
	bk_cmd_src := exec.Command("curl", "-L", bk_src_url, "-o", "source.tar.gz")
	err = bk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bk_cmd_direct := exec.Command("./binary")
	err = bk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
