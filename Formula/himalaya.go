package main

// Himalaya - CLI email client written in Rust
// Homepage: https://pimalaya.org/himalaya/cli/latest/

import (
	"fmt"
	
	"os/exec"
)

func installHimalaya() {
	// Método 1: Descargar y extraer .tar.gz
	himalaya_tar_url := "https://github.com/soywod/himalaya/archive/refs/tags/v0.9.0.tar.gz"
	himalaya_cmd_tar := exec.Command("curl", "-L", himalaya_tar_url, "-o", "package.tar.gz")
	err := himalaya_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	himalaya_zip_url := "https://github.com/soywod/himalaya/archive/refs/tags/v0.9.0.zip"
	himalaya_cmd_zip := exec.Command("curl", "-L", himalaya_zip_url, "-o", "package.zip")
	err = himalaya_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	himalaya_bin_url := "https://github.com/soywod/himalaya/archive/refs/tags/v0.9.0.bin"
	himalaya_cmd_bin := exec.Command("curl", "-L", himalaya_bin_url, "-o", "binary.bin")
	err = himalaya_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	himalaya_src_url := "https://github.com/soywod/himalaya/archive/refs/tags/v0.9.0.src.tar.gz"
	himalaya_cmd_src := exec.Command("curl", "-L", himalaya_src_url, "-o", "source.tar.gz")
	err = himalaya_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	himalaya_cmd_direct := exec.Command("./binary")
	err = himalaya_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
