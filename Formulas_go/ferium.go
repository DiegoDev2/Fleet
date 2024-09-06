package main

// Ferium - Fast and multi-source CLI program for managing Minecraft mods and modpacks
// Homepage: https://github.com/gorilla-devs/ferium

import (
	"fmt"
	
	"os/exec"
)

func installFerium() {
	// Método 1: Descargar y extraer .tar.gz
	ferium_tar_url := "https://github.com/gorilla-devs/ferium/archive/refs/tags/v4.7.0.tar.gz"
	ferium_cmd_tar := exec.Command("curl", "-L", ferium_tar_url, "-o", "package.tar.gz")
	err := ferium_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ferium_zip_url := "https://github.com/gorilla-devs/ferium/archive/refs/tags/v4.7.0.zip"
	ferium_cmd_zip := exec.Command("curl", "-L", ferium_zip_url, "-o", "package.zip")
	err = ferium_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ferium_bin_url := "https://github.com/gorilla-devs/ferium/archive/refs/tags/v4.7.0.bin"
	ferium_cmd_bin := exec.Command("curl", "-L", ferium_bin_url, "-o", "binary.bin")
	err = ferium_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ferium_src_url := "https://github.com/gorilla-devs/ferium/archive/refs/tags/v4.7.0.src.tar.gz"
	ferium_cmd_src := exec.Command("curl", "-L", ferium_src_url, "-o", "source.tar.gz")
	err = ferium_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ferium_cmd_direct := exec.Command("./binary")
	err = ferium_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
