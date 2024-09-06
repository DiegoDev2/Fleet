package main

// CargoDeny - Cargo plugin for linting your dependencies
// Homepage: https://github.com/EmbarkStudios/cargo-deny

import (
	"fmt"
	
	"os/exec"
)

func installCargoDeny() {
	// Método 1: Descargar y extraer .tar.gz
	cargodeny_tar_url := "https://github.com/EmbarkStudios/cargo-deny/archive/refs/tags/0.16.1.tar.gz"
	cargodeny_cmd_tar := exec.Command("curl", "-L", cargodeny_tar_url, "-o", "package.tar.gz")
	err := cargodeny_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cargodeny_zip_url := "https://github.com/EmbarkStudios/cargo-deny/archive/refs/tags/0.16.1.zip"
	cargodeny_cmd_zip := exec.Command("curl", "-L", cargodeny_zip_url, "-o", "package.zip")
	err = cargodeny_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cargodeny_bin_url := "https://github.com/EmbarkStudios/cargo-deny/archive/refs/tags/0.16.1.bin"
	cargodeny_cmd_bin := exec.Command("curl", "-L", cargodeny_bin_url, "-o", "binary.bin")
	err = cargodeny_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cargodeny_src_url := "https://github.com/EmbarkStudios/cargo-deny/archive/refs/tags/0.16.1.src.tar.gz"
	cargodeny_cmd_src := exec.Command("curl", "-L", cargodeny_src_url, "-o", "source.tar.gz")
	err = cargodeny_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cargodeny_cmd_direct := exec.Command("./binary")
	err = cargodeny_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: rustup")
exec.Command("latte", "install", "rustup")
}
