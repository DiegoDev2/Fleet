package main

// GitCliff - Highly customizable changelog generator
// Homepage: https://github.com/orhun/git-cliff

import (
	"fmt"
	
	"os/exec"
)

func installGitCliff() {
	// Método 1: Descargar y extraer .tar.gz
	gitcliff_tar_url := "https://github.com/orhun/git-cliff/archive/refs/tags/v2.5.0.tar.gz"
	gitcliff_cmd_tar := exec.Command("curl", "-L", gitcliff_tar_url, "-o", "package.tar.gz")
	err := gitcliff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitcliff_zip_url := "https://github.com/orhun/git-cliff/archive/refs/tags/v2.5.0.zip"
	gitcliff_cmd_zip := exec.Command("curl", "-L", gitcliff_zip_url, "-o", "package.zip")
	err = gitcliff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitcliff_bin_url := "https://github.com/orhun/git-cliff/archive/refs/tags/v2.5.0.bin"
	gitcliff_cmd_bin := exec.Command("curl", "-L", gitcliff_bin_url, "-o", "binary.bin")
	err = gitcliff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitcliff_src_url := "https://github.com/orhun/git-cliff/archive/refs/tags/v2.5.0.src.tar.gz"
	gitcliff_cmd_src := exec.Command("curl", "-L", gitcliff_src_url, "-o", "source.tar.gz")
	err = gitcliff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitcliff_cmd_direct := exec.Command("./binary")
	err = gitcliff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: libgit2")
exec.Command("latte", "install", "libgit2")
}
