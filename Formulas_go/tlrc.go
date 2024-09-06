package main

// Tlrc - Official tldr client written in Rust
// Homepage: https://github.com/tldr-pages/tlrc

import (
	"fmt"
	
	"os/exec"
)

func installTlrc() {
	// Método 1: Descargar y extraer .tar.gz
	tlrc_tar_url := "https://github.com/tldr-pages/tlrc/archive/refs/tags/v1.9.3.tar.gz"
	tlrc_cmd_tar := exec.Command("curl", "-L", tlrc_tar_url, "-o", "package.tar.gz")
	err := tlrc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tlrc_zip_url := "https://github.com/tldr-pages/tlrc/archive/refs/tags/v1.9.3.zip"
	tlrc_cmd_zip := exec.Command("curl", "-L", tlrc_zip_url, "-o", "package.zip")
	err = tlrc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tlrc_bin_url := "https://github.com/tldr-pages/tlrc/archive/refs/tags/v1.9.3.bin"
	tlrc_cmd_bin := exec.Command("curl", "-L", tlrc_bin_url, "-o", "binary.bin")
	err = tlrc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tlrc_src_url := "https://github.com/tldr-pages/tlrc/archive/refs/tags/v1.9.3.src.tar.gz"
	tlrc_cmd_src := exec.Command("curl", "-L", tlrc_src_url, "-o", "source.tar.gz")
	err = tlrc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tlrc_cmd_direct := exec.Command("./binary")
	err = tlrc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
