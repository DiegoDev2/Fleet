package main

// GitBranchless - High-velocity, monorepo-scale workflow for Git
// Homepage: https://github.com/arxanas/git-branchless

import (
	"fmt"
	
	"os/exec"
)

func installGitBranchless() {
	// Método 1: Descargar y extraer .tar.gz
	gitbranchless_tar_url := "https://github.com/arxanas/git-branchless/archive/refs/tags/v0.9.0.tar.gz"
	gitbranchless_cmd_tar := exec.Command("curl", "-L", gitbranchless_tar_url, "-o", "package.tar.gz")
	err := gitbranchless_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitbranchless_zip_url := "https://github.com/arxanas/git-branchless/archive/refs/tags/v0.9.0.zip"
	gitbranchless_cmd_zip := exec.Command("curl", "-L", gitbranchless_zip_url, "-o", "package.zip")
	err = gitbranchless_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitbranchless_bin_url := "https://github.com/arxanas/git-branchless/archive/refs/tags/v0.9.0.bin"
	gitbranchless_cmd_bin := exec.Command("curl", "-L", gitbranchless_bin_url, "-o", "binary.bin")
	err = gitbranchless_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitbranchless_src_url := "https://github.com/arxanas/git-branchless/archive/refs/tags/v0.9.0.src.tar.gz"
	gitbranchless_cmd_src := exec.Command("curl", "-L", gitbranchless_src_url, "-o", "source.tar.gz")
	err = gitbranchless_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitbranchless_cmd_direct := exec.Command("./binary")
	err = gitbranchless_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: libgit2@1.7")
	exec.Command("latte", "install", "libgit2@1.7").Run()
}
