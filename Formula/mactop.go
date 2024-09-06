package main

// Mactop - Apple Silicon Monitor Top written in Golang
// Homepage: https://github.com/context-labs/mactop

import (
	"fmt"
	
	"os/exec"
)

func installMactop() {
	// Método 1: Descargar y extraer .tar.gz
	mactop_tar_url := "https://github.com/context-labs/mactop/archive/refs/tags/v0.1.9.tar.gz"
	mactop_cmd_tar := exec.Command("curl", "-L", mactop_tar_url, "-o", "package.tar.gz")
	err := mactop_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mactop_zip_url := "https://github.com/context-labs/mactop/archive/refs/tags/v0.1.9.zip"
	mactop_cmd_zip := exec.Command("curl", "-L", mactop_zip_url, "-o", "package.zip")
	err = mactop_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mactop_bin_url := "https://github.com/context-labs/mactop/archive/refs/tags/v0.1.9.bin"
	mactop_cmd_bin := exec.Command("curl", "-L", mactop_bin_url, "-o", "binary.bin")
	err = mactop_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mactop_src_url := "https://github.com/context-labs/mactop/archive/refs/tags/v0.1.9.src.tar.gz"
	mactop_cmd_src := exec.Command("curl", "-L", mactop_src_url, "-o", "source.tar.gz")
	err = mactop_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mactop_cmd_direct := exec.Command("./binary")
	err = mactop_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
