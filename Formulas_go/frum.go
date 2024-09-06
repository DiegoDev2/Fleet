package main

// Frum - Fast and modern Ruby version manager written in Rust
// Homepage: https://github.com/TaKO8Ki/frum/

import (
	"fmt"
	
	"os/exec"
)

func installFrum() {
	// Método 1: Descargar y extraer .tar.gz
	frum_tar_url := "https://github.com/TaKO8Ki/frum/archive/refs/tags/v0.1.2.tar.gz"
	frum_cmd_tar := exec.Command("curl", "-L", frum_tar_url, "-o", "package.tar.gz")
	err := frum_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	frum_zip_url := "https://github.com/TaKO8Ki/frum/archive/refs/tags/v0.1.2.zip"
	frum_cmd_zip := exec.Command("curl", "-L", frum_zip_url, "-o", "package.zip")
	err = frum_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	frum_bin_url := "https://github.com/TaKO8Ki/frum/archive/refs/tags/v0.1.2.bin"
	frum_cmd_bin := exec.Command("curl", "-L", frum_bin_url, "-o", "binary.bin")
	err = frum_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	frum_src_url := "https://github.com/TaKO8Ki/frum/archive/refs/tags/v0.1.2.src.tar.gz"
	frum_cmd_src := exec.Command("curl", "-L", frum_src_url, "-o", "source.tar.gz")
	err = frum_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	frum_cmd_direct := exec.Command("./binary")
	err = frum_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
