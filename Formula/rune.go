package main

// Rune - Embeddable dynamic programming language for Rust
// Homepage: https://rune-rs.github.io

import (
	"fmt"
	
	"os/exec"
)

func installRune() {
	// Método 1: Descargar y extraer .tar.gz
	rune_tar_url := "https://github.com/rune-rs/rune/archive/refs/tags/0.13.4.tar.gz"
	rune_cmd_tar := exec.Command("curl", "-L", rune_tar_url, "-o", "package.tar.gz")
	err := rune_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rune_zip_url := "https://github.com/rune-rs/rune/archive/refs/tags/0.13.4.zip"
	rune_cmd_zip := exec.Command("curl", "-L", rune_zip_url, "-o", "package.zip")
	err = rune_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rune_bin_url := "https://github.com/rune-rs/rune/archive/refs/tags/0.13.4.bin"
	rune_cmd_bin := exec.Command("curl", "-L", rune_bin_url, "-o", "binary.bin")
	err = rune_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rune_src_url := "https://github.com/rune-rs/rune/archive/refs/tags/0.13.4.src.tar.gz"
	rune_cmd_src := exec.Command("curl", "-L", rune_src_url, "-o", "source.tar.gz")
	err = rune_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rune_cmd_direct := exec.Command("./binary")
	err = rune_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
