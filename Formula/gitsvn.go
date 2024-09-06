package main

// GitSvn - Bidirectional operation between a Subversion repository and Git
// Homepage: https://git-scm.com

import (
	"fmt"
	
	"os/exec"
)

func installGitSvn() {
	// Método 1: Descargar y extraer .tar.gz
	gitsvn_tar_url := "https://mirrors.edge.kernel.org/pub/software/scm/git/git-2.46.0.tar.xz"
	gitsvn_cmd_tar := exec.Command("curl", "-L", gitsvn_tar_url, "-o", "package.tar.gz")
	err := gitsvn_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitsvn_zip_url := "https://mirrors.edge.kernel.org/pub/software/scm/git/git-2.46.0.tar.xz"
	gitsvn_cmd_zip := exec.Command("curl", "-L", gitsvn_zip_url, "-o", "package.zip")
	err = gitsvn_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitsvn_bin_url := "https://mirrors.edge.kernel.org/pub/software/scm/git/git-2.46.0.tar.xz"
	gitsvn_cmd_bin := exec.Command("curl", "-L", gitsvn_bin_url, "-o", "binary.bin")
	err = gitsvn_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitsvn_src_url := "https://mirrors.edge.kernel.org/pub/software/scm/git/git-2.46.0.tar.xz"
	gitsvn_cmd_src := exec.Command("curl", "-L", gitsvn_src_url, "-o", "source.tar.gz")
	err = gitsvn_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitsvn_cmd_direct := exec.Command("./binary")
	err = gitsvn_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: git")
	exec.Command("latte", "install", "git").Run()
	fmt.Println("Instalando dependencia: subversion")
	exec.Command("latte", "install", "subversion").Run()
}
