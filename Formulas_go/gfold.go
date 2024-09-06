package main

// Gfold - Help keep track of your Git repositories, written in Rust
// Homepage: https://github.com/nickgerace/gfold

import (
	"fmt"
	
	"os/exec"
)

func installGfold() {
	// Método 1: Descargar y extraer .tar.gz
	gfold_tar_url := "https://github.com/nickgerace/gfold/archive/refs/tags/4.5.0.tar.gz"
	gfold_cmd_tar := exec.Command("curl", "-L", gfold_tar_url, "-o", "package.tar.gz")
	err := gfold_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gfold_zip_url := "https://github.com/nickgerace/gfold/archive/refs/tags/4.5.0.zip"
	gfold_cmd_zip := exec.Command("curl", "-L", gfold_zip_url, "-o", "package.zip")
	err = gfold_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gfold_bin_url := "https://github.com/nickgerace/gfold/archive/refs/tags/4.5.0.bin"
	gfold_cmd_bin := exec.Command("curl", "-L", gfold_bin_url, "-o", "binary.bin")
	err = gfold_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gfold_src_url := "https://github.com/nickgerace/gfold/archive/refs/tags/4.5.0.src.tar.gz"
	gfold_cmd_src := exec.Command("curl", "-L", gfold_src_url, "-o", "source.tar.gz")
	err = gfold_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gfold_cmd_direct := exec.Command("./binary")
	err = gfold_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: libgit2@1.7")
exec.Command("latte", "install", "libgit2@1.7")
}
