package main

// CargoCrev - Code review system for the cargo package manager
// Homepage: https://web.crev.dev/rust-reviews/

import (
	"fmt"
	
	"os/exec"
)

func installCargoCrev() {
	// Método 1: Descargar y extraer .tar.gz
	cargocrev_tar_url := "https://github.com/crev-dev/cargo-crev/archive/refs/tags/v0.25.9.tar.gz"
	cargocrev_cmd_tar := exec.Command("curl", "-L", cargocrev_tar_url, "-o", "package.tar.gz")
	err := cargocrev_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cargocrev_zip_url := "https://github.com/crev-dev/cargo-crev/archive/refs/tags/v0.25.9.zip"
	cargocrev_cmd_zip := exec.Command("curl", "-L", cargocrev_zip_url, "-o", "package.zip")
	err = cargocrev_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cargocrev_bin_url := "https://github.com/crev-dev/cargo-crev/archive/refs/tags/v0.25.9.bin"
	cargocrev_cmd_bin := exec.Command("curl", "-L", cargocrev_bin_url, "-o", "binary.bin")
	err = cargocrev_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cargocrev_src_url := "https://github.com/crev-dev/cargo-crev/archive/refs/tags/v0.25.9.src.tar.gz"
	cargocrev_cmd_src := exec.Command("curl", "-L", cargocrev_src_url, "-o", "source.tar.gz")
	err = cargocrev_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cargocrev_cmd_direct := exec.Command("./binary")
	err = cargocrev_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: rustup")
	exec.Command("latte", "install", "rustup").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
