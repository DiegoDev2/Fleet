package main

// Oxipng - Multithreaded PNG optimizer written in Rust
// Homepage: https://github.com/shssoichiro/oxipng

import (
	"fmt"
	
	"os/exec"
)

func installOxipng() {
	// Método 1: Descargar y extraer .tar.gz
	oxipng_tar_url := "https://github.com/shssoichiro/oxipng/archive/refs/tags/v9.1.2.tar.gz"
	oxipng_cmd_tar := exec.Command("curl", "-L", oxipng_tar_url, "-o", "package.tar.gz")
	err := oxipng_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	oxipng_zip_url := "https://github.com/shssoichiro/oxipng/archive/refs/tags/v9.1.2.zip"
	oxipng_cmd_zip := exec.Command("curl", "-L", oxipng_zip_url, "-o", "package.zip")
	err = oxipng_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	oxipng_bin_url := "https://github.com/shssoichiro/oxipng/archive/refs/tags/v9.1.2.bin"
	oxipng_cmd_bin := exec.Command("curl", "-L", oxipng_bin_url, "-o", "binary.bin")
	err = oxipng_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	oxipng_src_url := "https://github.com/shssoichiro/oxipng/archive/refs/tags/v9.1.2.src.tar.gz"
	oxipng_cmd_src := exec.Command("curl", "-L", oxipng_src_url, "-o", "source.tar.gz")
	err = oxipng_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	oxipng_cmd_direct := exec.Command("./binary")
	err = oxipng_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
