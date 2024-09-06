package main

// GitAbsorb - Automatic git commit --fixup
// Homepage: https://github.com/tummychow/git-absorb

import (
	"fmt"
	
	"os/exec"
)

func installGitAbsorb() {
	// Método 1: Descargar y extraer .tar.gz
	gitabsorb_tar_url := "https://github.com/tummychow/git-absorb/archive/refs/tags/0.6.15.tar.gz"
	gitabsorb_cmd_tar := exec.Command("curl", "-L", gitabsorb_tar_url, "-o", "package.tar.gz")
	err := gitabsorb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitabsorb_zip_url := "https://github.com/tummychow/git-absorb/archive/refs/tags/0.6.15.zip"
	gitabsorb_cmd_zip := exec.Command("curl", "-L", gitabsorb_zip_url, "-o", "package.zip")
	err = gitabsorb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitabsorb_bin_url := "https://github.com/tummychow/git-absorb/archive/refs/tags/0.6.15.bin"
	gitabsorb_cmd_bin := exec.Command("curl", "-L", gitabsorb_bin_url, "-o", "binary.bin")
	err = gitabsorb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitabsorb_src_url := "https://github.com/tummychow/git-absorb/archive/refs/tags/0.6.15.src.tar.gz"
	gitabsorb_cmd_src := exec.Command("curl", "-L", gitabsorb_src_url, "-o", "source.tar.gz")
	err = gitabsorb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitabsorb_cmd_direct := exec.Command("./binary")
	err = gitabsorb_cmd_direct.Run()
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
