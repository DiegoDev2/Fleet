package main

// Git - Distributed revision control system
// Homepage: https://git-scm.com

import (
	"fmt"
	
	"os/exec"
)

func installGit() {
	// Método 1: Descargar y extraer .tar.gz
	git_tar_url := "https://mirrors.edge.kernel.org/pub/software/scm/git/git-2.46.0.tar.xz"
	git_cmd_tar := exec.Command("curl", "-L", git_tar_url, "-o", "package.tar.gz")
	err := git_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	git_zip_url := "https://mirrors.edge.kernel.org/pub/software/scm/git/git-2.46.0.tar.xz"
	git_cmd_zip := exec.Command("curl", "-L", git_zip_url, "-o", "package.zip")
	err = git_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	git_bin_url := "https://mirrors.edge.kernel.org/pub/software/scm/git/git-2.46.0.tar.xz"
	git_cmd_bin := exec.Command("curl", "-L", git_bin_url, "-o", "binary.bin")
	err = git_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	git_src_url := "https://mirrors.edge.kernel.org/pub/software/scm/git/git-2.46.0.tar.xz"
	git_cmd_src := exec.Command("curl", "-L", git_src_url, "-o", "source.tar.gz")
	err = git_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	git_cmd_direct := exec.Command("./binary")
	err = git_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
	fmt.Println("Instalando dependencia: linux-headers@5.15")
exec.Command("latte", "install", "linux-headers@5.15")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
