package main

// Muffet - Fast website link checker in Go
// Homepage: https://github.com/raviqqe/muffet

import (
	"fmt"
	
	"os/exec"
)

func installMuffet() {
	// Método 1: Descargar y extraer .tar.gz
	muffet_tar_url := "https://github.com/raviqqe/muffet/archive/refs/tags/v2.10.2.tar.gz"
	muffet_cmd_tar := exec.Command("curl", "-L", muffet_tar_url, "-o", "package.tar.gz")
	err := muffet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	muffet_zip_url := "https://github.com/raviqqe/muffet/archive/refs/tags/v2.10.2.zip"
	muffet_cmd_zip := exec.Command("curl", "-L", muffet_zip_url, "-o", "package.zip")
	err = muffet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	muffet_bin_url := "https://github.com/raviqqe/muffet/archive/refs/tags/v2.10.2.bin"
	muffet_cmd_bin := exec.Command("curl", "-L", muffet_bin_url, "-o", "binary.bin")
	err = muffet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	muffet_src_url := "https://github.com/raviqqe/muffet/archive/refs/tags/v2.10.2.src.tar.gz"
	muffet_cmd_src := exec.Command("curl", "-L", muffet_src_url, "-o", "source.tar.gz")
	err = muffet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	muffet_cmd_direct := exec.Command("./binary")
	err = muffet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
