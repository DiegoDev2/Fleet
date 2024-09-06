package main

// Gitfs - Version controlled file system
// Homepage: https://www.presslabs.com/gitfs

import (
	"fmt"
	
	"os/exec"
)

func installGitfs() {
	// Método 1: Descargar y extraer .tar.gz
	gitfs_tar_url := "https://github.com/presslabs/gitfs/archive/refs/tags/0.5.2.tar.gz"
	gitfs_cmd_tar := exec.Command("curl", "-L", gitfs_tar_url, "-o", "package.tar.gz")
	err := gitfs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitfs_zip_url := "https://github.com/presslabs/gitfs/archive/refs/tags/0.5.2.zip"
	gitfs_cmd_zip := exec.Command("curl", "-L", gitfs_zip_url, "-o", "package.zip")
	err = gitfs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitfs_bin_url := "https://github.com/presslabs/gitfs/archive/refs/tags/0.5.2.bin"
	gitfs_cmd_bin := exec.Command("curl", "-L", gitfs_bin_url, "-o", "binary.bin")
	err = gitfs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitfs_src_url := "https://github.com/presslabs/gitfs/archive/refs/tags/0.5.2.src.tar.gz"
	gitfs_cmd_src := exec.Command("curl", "-L", gitfs_src_url, "-o", "source.tar.gz")
	err = gitfs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitfs_cmd_direct := exec.Command("./binary")
	err = gitfs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libffi")
	exec.Command("latte", "install", "libffi").Run()
	fmt.Println("Instalando dependencia: libfuse")
	exec.Command("latte", "install", "libfuse").Run()
	fmt.Println("Instalando dependencia: libgit2")
	exec.Command("latte", "install", "libgit2").Run()
	fmt.Println("Instalando dependencia: python@3.9")
	exec.Command("latte", "install", "python@3.9").Run()
}
