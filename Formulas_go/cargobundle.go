package main

// CargoBundle - Wrap rust executables in OS-specific app bundles
// Homepage: https://github.com/burtonageo/cargo-bundle

import (
	"fmt"
	
	"os/exec"
)

func installCargoBundle() {
	// Método 1: Descargar y extraer .tar.gz
	cargobundle_tar_url := "https://github.com/burtonageo/cargo-bundle/archive/refs/tags/v0.6.1.tar.gz"
	cargobundle_cmd_tar := exec.Command("curl", "-L", cargobundle_tar_url, "-o", "package.tar.gz")
	err := cargobundle_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cargobundle_zip_url := "https://github.com/burtonageo/cargo-bundle/archive/refs/tags/v0.6.1.zip"
	cargobundle_cmd_zip := exec.Command("curl", "-L", cargobundle_zip_url, "-o", "package.zip")
	err = cargobundle_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cargobundle_bin_url := "https://github.com/burtonageo/cargo-bundle/archive/refs/tags/v0.6.1.bin"
	cargobundle_cmd_bin := exec.Command("curl", "-L", cargobundle_bin_url, "-o", "binary.bin")
	err = cargobundle_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cargobundle_src_url := "https://github.com/burtonageo/cargo-bundle/archive/refs/tags/v0.6.1.src.tar.gz"
	cargobundle_cmd_src := exec.Command("curl", "-L", cargobundle_src_url, "-o", "source.tar.gz")
	err = cargobundle_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cargobundle_cmd_direct := exec.Command("./binary")
	err = cargobundle_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: rustup")
exec.Command("latte", "install", "rustup")
}
