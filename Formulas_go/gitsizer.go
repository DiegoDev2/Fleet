package main

// GitSizer - Compute various size metrics for a Git repository
// Homepage: https://github.com/github/git-sizer

import (
	"fmt"
	
	"os/exec"
)

func installGitSizer() {
	// Método 1: Descargar y extraer .tar.gz
	gitsizer_tar_url := "https://github.com/github/git-sizer/archive/refs/tags/v1.5.0.tar.gz"
	gitsizer_cmd_tar := exec.Command("curl", "-L", gitsizer_tar_url, "-o", "package.tar.gz")
	err := gitsizer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitsizer_zip_url := "https://github.com/github/git-sizer/archive/refs/tags/v1.5.0.zip"
	gitsizer_cmd_zip := exec.Command("curl", "-L", gitsizer_zip_url, "-o", "package.zip")
	err = gitsizer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitsizer_bin_url := "https://github.com/github/git-sizer/archive/refs/tags/v1.5.0.bin"
	gitsizer_cmd_bin := exec.Command("curl", "-L", gitsizer_bin_url, "-o", "binary.bin")
	err = gitsizer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitsizer_src_url := "https://github.com/github/git-sizer/archive/refs/tags/v1.5.0.src.tar.gz"
	gitsizer_cmd_src := exec.Command("curl", "-L", gitsizer_src_url, "-o", "source.tar.gz")
	err = gitsizer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitsizer_cmd_direct := exec.Command("./binary")
	err = gitsizer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
