package main

// Tldr - Simplified and community-driven man pages
// Homepage: https://tldr.sh/

import (
	"fmt"
	
	"os/exec"
)

func installTldr() {
	// Método 1: Descargar y extraer .tar.gz
	tldr_tar_url := "https://github.com/tldr-pages/tldr-c-client/archive/refs/tags/v1.6.1.tar.gz"
	tldr_cmd_tar := exec.Command("curl", "-L", tldr_tar_url, "-o", "package.tar.gz")
	err := tldr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tldr_zip_url := "https://github.com/tldr-pages/tldr-c-client/archive/refs/tags/v1.6.1.zip"
	tldr_cmd_zip := exec.Command("curl", "-L", tldr_zip_url, "-o", "package.zip")
	err = tldr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tldr_bin_url := "https://github.com/tldr-pages/tldr-c-client/archive/refs/tags/v1.6.1.bin"
	tldr_cmd_bin := exec.Command("curl", "-L", tldr_bin_url, "-o", "binary.bin")
	err = tldr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tldr_src_url := "https://github.com/tldr-pages/tldr-c-client/archive/refs/tags/v1.6.1.src.tar.gz"
	tldr_cmd_src := exec.Command("curl", "-L", tldr_src_url, "-o", "source.tar.gz")
	err = tldr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tldr_cmd_direct := exec.Command("./binary")
	err = tldr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libzip")
exec.Command("latte", "install", "libzip")
}
