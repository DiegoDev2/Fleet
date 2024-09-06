package main

// Lychee - Fast, async, resource-friendly link checker
// Homepage: https://lychee.cli.rs/

import (
	"fmt"
	
	"os/exec"
)

func installLychee() {
	// Método 1: Descargar y extraer .tar.gz
	lychee_tar_url := "https://github.com/lycheeverse/lychee/archive/refs/tags/v0.15.1.tar.gz"
	lychee_cmd_tar := exec.Command("curl", "-L", lychee_tar_url, "-o", "package.tar.gz")
	err := lychee_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lychee_zip_url := "https://github.com/lycheeverse/lychee/archive/refs/tags/v0.15.1.zip"
	lychee_cmd_zip := exec.Command("curl", "-L", lychee_zip_url, "-o", "package.zip")
	err = lychee_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lychee_bin_url := "https://github.com/lycheeverse/lychee/archive/refs/tags/v0.15.1.bin"
	lychee_cmd_bin := exec.Command("curl", "-L", lychee_bin_url, "-o", "binary.bin")
	err = lychee_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lychee_src_url := "https://github.com/lycheeverse/lychee/archive/refs/tags/v0.15.1.src.tar.gz"
	lychee_cmd_src := exec.Command("curl", "-L", lychee_src_url, "-o", "source.tar.gz")
	err = lychee_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lychee_cmd_direct := exec.Command("./binary")
	err = lychee_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
