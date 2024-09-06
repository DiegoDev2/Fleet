package main

// Eza - Modern, maintained replacement for ls
// Homepage: https://github.com/eza-community/eza

import (
	"fmt"
	
	"os/exec"
)

func installEza() {
	// Método 1: Descargar y extraer .tar.gz
	eza_tar_url := "https://github.com/eza-community/eza/archive/refs/tags/v0.19.2.tar.gz"
	eza_cmd_tar := exec.Command("curl", "-L", eza_tar_url, "-o", "package.tar.gz")
	err := eza_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	eza_zip_url := "https://github.com/eza-community/eza/archive/refs/tags/v0.19.2.zip"
	eza_cmd_zip := exec.Command("curl", "-L", eza_zip_url, "-o", "package.zip")
	err = eza_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	eza_bin_url := "https://github.com/eza-community/eza/archive/refs/tags/v0.19.2.bin"
	eza_cmd_bin := exec.Command("curl", "-L", eza_bin_url, "-o", "binary.bin")
	err = eza_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	eza_src_url := "https://github.com/eza-community/eza/archive/refs/tags/v0.19.2.src.tar.gz"
	eza_cmd_src := exec.Command("curl", "-L", eza_src_url, "-o", "source.tar.gz")
	err = eza_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	eza_cmd_direct := exec.Command("./binary")
	err = eza_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pandoc")
exec.Command("latte", "install", "pandoc")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: libgit2")
exec.Command("latte", "install", "libgit2")
}
