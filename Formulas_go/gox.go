package main

// Gox - Go cross compile tool
// Homepage: https://github.com/mitchellh/gox

import (
	"fmt"
	
	"os/exec"
)

func installGox() {
	// Método 1: Descargar y extraer .tar.gz
	gox_tar_url := "https://github.com/mitchellh/gox/archive/refs/tags/v1.0.1.tar.gz"
	gox_cmd_tar := exec.Command("curl", "-L", gox_tar_url, "-o", "package.tar.gz")
	err := gox_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gox_zip_url := "https://github.com/mitchellh/gox/archive/refs/tags/v1.0.1.zip"
	gox_cmd_zip := exec.Command("curl", "-L", gox_zip_url, "-o", "package.zip")
	err = gox_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gox_bin_url := "https://github.com/mitchellh/gox/archive/refs/tags/v1.0.1.bin"
	gox_cmd_bin := exec.Command("curl", "-L", gox_bin_url, "-o", "binary.bin")
	err = gox_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gox_src_url := "https://github.com/mitchellh/gox/archive/refs/tags/v1.0.1.src.tar.gz"
	gox_cmd_src := exec.Command("curl", "-L", gox_src_url, "-o", "source.tar.gz")
	err = gox_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gox_cmd_direct := exec.Command("./binary")
	err = gox_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
