package main

// Immortal - OS agnostic (*nix) cross-platform supervisor
// Homepage: https://immortal.run/

import (
	"fmt"
	
	"os/exec"
)

func installImmortal() {
	// Método 1: Descargar y extraer .tar.gz
	immortal_tar_url := "https://github.com/immortal/immortal/archive/refs/tags/v0.24.6.tar.gz"
	immortal_cmd_tar := exec.Command("curl", "-L", immortal_tar_url, "-o", "package.tar.gz")
	err := immortal_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	immortal_zip_url := "https://github.com/immortal/immortal/archive/refs/tags/v0.24.6.zip"
	immortal_cmd_zip := exec.Command("curl", "-L", immortal_zip_url, "-o", "package.zip")
	err = immortal_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	immortal_bin_url := "https://github.com/immortal/immortal/archive/refs/tags/v0.24.6.bin"
	immortal_cmd_bin := exec.Command("curl", "-L", immortal_bin_url, "-o", "binary.bin")
	err = immortal_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	immortal_src_url := "https://github.com/immortal/immortal/archive/refs/tags/v0.24.6.src.tar.gz"
	immortal_cmd_src := exec.Command("curl", "-L", immortal_src_url, "-o", "source.tar.gz")
	err = immortal_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	immortal_cmd_direct := exec.Command("./binary")
	err = immortal_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
