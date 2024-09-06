package main

// MozGitTools - Tools for working with Git at Mozilla
// Homepage: https://github.com/mozilla/moz-git-tools

import (
	"fmt"
	
	"os/exec"
)

func installMozGitTools() {
	// Método 1: Descargar y extraer .tar.gz
	mozgittools_tar_url := "https://github.com/mozilla/moz-git-tools/archive/refs/tags/v0.1.tar.gz"
	mozgittools_cmd_tar := exec.Command("curl", "-L", mozgittools_tar_url, "-o", "package.tar.gz")
	err := mozgittools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mozgittools_zip_url := "https://github.com/mozilla/moz-git-tools/archive/refs/tags/v0.1.zip"
	mozgittools_cmd_zip := exec.Command("curl", "-L", mozgittools_zip_url, "-o", "package.zip")
	err = mozgittools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mozgittools_bin_url := "https://github.com/mozilla/moz-git-tools/archive/refs/tags/v0.1.bin"
	mozgittools_cmd_bin := exec.Command("curl", "-L", mozgittools_bin_url, "-o", "binary.bin")
	err = mozgittools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mozgittools_src_url := "https://github.com/mozilla/moz-git-tools/archive/refs/tags/v0.1.src.tar.gz"
	mozgittools_cmd_src := exec.Command("curl", "-L", mozgittools_src_url, "-o", "source.tar.gz")
	err = mozgittools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mozgittools_cmd_direct := exec.Command("./binary")
	err = mozgittools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
