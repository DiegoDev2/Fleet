package main

// Flawz - Terminal UI for browsing security vulnerabilities (CVEs)
// Homepage: https://github.com/orhun/flawz

import (
	"fmt"
	
	"os/exec"
)

func installFlawz() {
	// Método 1: Descargar y extraer .tar.gz
	flawz_tar_url := "https://github.com/orhun/flawz/archive/refs/tags/v0.2.2.tar.gz"
	flawz_cmd_tar := exec.Command("curl", "-L", flawz_tar_url, "-o", "package.tar.gz")
	err := flawz_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	flawz_zip_url := "https://github.com/orhun/flawz/archive/refs/tags/v0.2.2.zip"
	flawz_cmd_zip := exec.Command("curl", "-L", flawz_zip_url, "-o", "package.zip")
	err = flawz_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	flawz_bin_url := "https://github.com/orhun/flawz/archive/refs/tags/v0.2.2.bin"
	flawz_cmd_bin := exec.Command("curl", "-L", flawz_bin_url, "-o", "binary.bin")
	err = flawz_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	flawz_src_url := "https://github.com/orhun/flawz/archive/refs/tags/v0.2.2.src.tar.gz"
	flawz_cmd_src := exec.Command("curl", "-L", flawz_src_url, "-o", "source.tar.gz")
	err = flawz_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	flawz_cmd_direct := exec.Command("./binary")
	err = flawz_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
