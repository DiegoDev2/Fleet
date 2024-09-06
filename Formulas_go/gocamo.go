package main

// GoCamo - Secure image proxy server
// Homepage: https://github.com/cactus/go-camo

import (
	"fmt"
	
	"os/exec"
)

func installGoCamo() {
	// Método 1: Descargar y extraer .tar.gz
	gocamo_tar_url := "https://github.com/cactus/go-camo/archive/refs/tags/v2.6.0.tar.gz"
	gocamo_cmd_tar := exec.Command("curl", "-L", gocamo_tar_url, "-o", "package.tar.gz")
	err := gocamo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gocamo_zip_url := "https://github.com/cactus/go-camo/archive/refs/tags/v2.6.0.zip"
	gocamo_cmd_zip := exec.Command("curl", "-L", gocamo_zip_url, "-o", "package.zip")
	err = gocamo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gocamo_bin_url := "https://github.com/cactus/go-camo/archive/refs/tags/v2.6.0.bin"
	gocamo_cmd_bin := exec.Command("curl", "-L", gocamo_bin_url, "-o", "binary.bin")
	err = gocamo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gocamo_src_url := "https://github.com/cactus/go-camo/archive/refs/tags/v2.6.0.src.tar.gz"
	gocamo_cmd_src := exec.Command("curl", "-L", gocamo_src_url, "-o", "source.tar.gz")
	err = gocamo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gocamo_cmd_direct := exec.Command("./binary")
	err = gocamo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
