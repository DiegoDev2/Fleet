package main

// Dura - Backs up your work automatically via Git commits
// Homepage: https://github.com/tkellogg/dura

import (
	"fmt"
	
	"os/exec"
)

func installDura() {
	// Método 1: Descargar y extraer .tar.gz
	dura_tar_url := "https://github.com/tkellogg/dura/archive/refs/tags/v0.2.0.tar.gz"
	dura_cmd_tar := exec.Command("curl", "-L", dura_tar_url, "-o", "package.tar.gz")
	err := dura_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dura_zip_url := "https://github.com/tkellogg/dura/archive/refs/tags/v0.2.0.zip"
	dura_cmd_zip := exec.Command("curl", "-L", dura_zip_url, "-o", "package.zip")
	err = dura_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dura_bin_url := "https://github.com/tkellogg/dura/archive/refs/tags/v0.2.0.bin"
	dura_cmd_bin := exec.Command("curl", "-L", dura_bin_url, "-o", "binary.bin")
	err = dura_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dura_src_url := "https://github.com/tkellogg/dura/archive/refs/tags/v0.2.0.src.tar.gz"
	dura_cmd_src := exec.Command("curl", "-L", dura_src_url, "-o", "source.tar.gz")
	err = dura_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dura_cmd_direct := exec.Command("./binary")
	err = dura_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
