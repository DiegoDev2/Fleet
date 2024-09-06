package main

// Termius - CLI tool for termius.com (aka serverauditor.com)
// Homepage: https://termius.com

import (
	"fmt"
	
	"os/exec"
)

func installTermius() {
	// Método 1: Descargar y extraer .tar.gz
	termius_tar_url := "https://github.com/termius/termius-cli/archive/refs/tags/v1.2.15.tar.gz"
	termius_cmd_tar := exec.Command("curl", "-L", termius_tar_url, "-o", "package.tar.gz")
	err := termius_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	termius_zip_url := "https://github.com/termius/termius-cli/archive/refs/tags/v1.2.15.zip"
	termius_cmd_zip := exec.Command("curl", "-L", termius_zip_url, "-o", "package.zip")
	err = termius_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	termius_bin_url := "https://github.com/termius/termius-cli/archive/refs/tags/v1.2.15.bin"
	termius_cmd_bin := exec.Command("curl", "-L", termius_bin_url, "-o", "binary.bin")
	err = termius_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	termius_src_url := "https://github.com/termius/termius-cli/archive/refs/tags/v1.2.15.src.tar.gz"
	termius_cmd_src := exec.Command("curl", "-L", termius_src_url, "-o", "source.tar.gz")
	err = termius_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	termius_cmd_direct := exec.Command("./binary")
	err = termius_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: python@3.10")
exec.Command("latte", "install", "python@3.10")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
