package main

// Gitoxide - Idiomatic, lean, fast & safe pure Rust implementation of Git
// Homepage: https://github.com/Byron/gitoxide

import (
	"fmt"
	
	"os/exec"
)

func installGitoxide() {
	// Método 1: Descargar y extraer .tar.gz
	gitoxide_tar_url := "https://github.com/Byron/gitoxide/archive/refs/tags/v0.37.0.tar.gz"
	gitoxide_cmd_tar := exec.Command("curl", "-L", gitoxide_tar_url, "-o", "package.tar.gz")
	err := gitoxide_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitoxide_zip_url := "https://github.com/Byron/gitoxide/archive/refs/tags/v0.37.0.zip"
	gitoxide_cmd_zip := exec.Command("curl", "-L", gitoxide_zip_url, "-o", "package.zip")
	err = gitoxide_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitoxide_bin_url := "https://github.com/Byron/gitoxide/archive/refs/tags/v0.37.0.bin"
	gitoxide_cmd_bin := exec.Command("curl", "-L", gitoxide_bin_url, "-o", "binary.bin")
	err = gitoxide_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitoxide_src_url := "https://github.com/Byron/gitoxide/archive/refs/tags/v0.37.0.src.tar.gz"
	gitoxide_cmd_src := exec.Command("curl", "-L", gitoxide_src_url, "-o", "source.tar.gz")
	err = gitoxide_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitoxide_cmd_direct := exec.Command("./binary")
	err = gitoxide_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
