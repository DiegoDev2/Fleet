package main

// C2rust - Migrate C code to Rust
// Homepage: https://c2rust.com/

import (
	"fmt"
	
	"os/exec"
)

func installC2rust() {
	// Método 1: Descargar y extraer .tar.gz
	c2rust_tar_url := "https://github.com/immunant/c2rust/archive/refs/tags/v0.19.0.tar.gz"
	c2rust_cmd_tar := exec.Command("curl", "-L", c2rust_tar_url, "-o", "package.tar.gz")
	err := c2rust_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	c2rust_zip_url := "https://github.com/immunant/c2rust/archive/refs/tags/v0.19.0.zip"
	c2rust_cmd_zip := exec.Command("curl", "-L", c2rust_zip_url, "-o", "package.zip")
	err = c2rust_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	c2rust_bin_url := "https://github.com/immunant/c2rust/archive/refs/tags/v0.19.0.bin"
	c2rust_cmd_bin := exec.Command("curl", "-L", c2rust_bin_url, "-o", "binary.bin")
	err = c2rust_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	c2rust_src_url := "https://github.com/immunant/c2rust/archive/refs/tags/v0.19.0.src.tar.gz"
	c2rust_cmd_src := exec.Command("curl", "-L", c2rust_src_url, "-o", "source.tar.gz")
	err = c2rust_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	c2rust_cmd_direct := exec.Command("./binary")
	err = c2rust_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
}
