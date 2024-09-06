package main

// GitWhenMerged - Find where a commit was merged in git
// Homepage: https://github.com/mhagger/git-when-merged

import (
	"fmt"
	
	"os/exec"
)

func installGitWhenMerged() {
	// Método 1: Descargar y extraer .tar.gz
	gitwhenmerged_tar_url := "https://github.com/mhagger/git-when-merged/archive/refs/tags/v1.2.1.tar.gz"
	gitwhenmerged_cmd_tar := exec.Command("curl", "-L", gitwhenmerged_tar_url, "-o", "package.tar.gz")
	err := gitwhenmerged_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitwhenmerged_zip_url := "https://github.com/mhagger/git-when-merged/archive/refs/tags/v1.2.1.zip"
	gitwhenmerged_cmd_zip := exec.Command("curl", "-L", gitwhenmerged_zip_url, "-o", "package.zip")
	err = gitwhenmerged_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitwhenmerged_bin_url := "https://github.com/mhagger/git-when-merged/archive/refs/tags/v1.2.1.bin"
	gitwhenmerged_cmd_bin := exec.Command("curl", "-L", gitwhenmerged_bin_url, "-o", "binary.bin")
	err = gitwhenmerged_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitwhenmerged_src_url := "https://github.com/mhagger/git-when-merged/archive/refs/tags/v1.2.1.src.tar.gz"
	gitwhenmerged_cmd_src := exec.Command("curl", "-L", gitwhenmerged_src_url, "-o", "source.tar.gz")
	err = gitwhenmerged_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitwhenmerged_cmd_direct := exec.Command("./binary")
	err = gitwhenmerged_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
