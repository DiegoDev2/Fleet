package main

// GitSvnAbandon - History-preserving svn-to-git migration
// Homepage: https://github.com/nothingmuch/git-svn-abandon

import (
	"fmt"
	
	"os/exec"
)

func installGitSvnAbandon() {
	// Método 1: Descargar y extraer .tar.gz
	gitsvnabandon_tar_url := "https://github.com/nothingmuch/git-svn-abandon/archive/refs/tags/0.0.1.tar.gz"
	gitsvnabandon_cmd_tar := exec.Command("curl", "-L", gitsvnabandon_tar_url, "-o", "package.tar.gz")
	err := gitsvnabandon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitsvnabandon_zip_url := "https://github.com/nothingmuch/git-svn-abandon/archive/refs/tags/0.0.1.zip"
	gitsvnabandon_cmd_zip := exec.Command("curl", "-L", gitsvnabandon_zip_url, "-o", "package.zip")
	err = gitsvnabandon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitsvnabandon_bin_url := "https://github.com/nothingmuch/git-svn-abandon/archive/refs/tags/0.0.1.bin"
	gitsvnabandon_cmd_bin := exec.Command("curl", "-L", gitsvnabandon_bin_url, "-o", "binary.bin")
	err = gitsvnabandon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitsvnabandon_src_url := "https://github.com/nothingmuch/git-svn-abandon/archive/refs/tags/0.0.1.src.tar.gz"
	gitsvnabandon_cmd_src := exec.Command("curl", "-L", gitsvnabandon_src_url, "-o", "source.tar.gz")
	err = gitsvnabandon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitsvnabandon_cmd_direct := exec.Command("./binary")
	err = gitsvnabandon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
