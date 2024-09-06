package main

// GitUrlSub - Recursively substitute remote URLs for multiple repos
// Homepage: https://gosuri.github.io/git-url-sub

import (
	"fmt"
	
	"os/exec"
)

func installGitUrlSub() {
	// Método 1: Descargar y extraer .tar.gz
	giturlsub_tar_url := "https://github.com/gosuri/git-url-sub/archive/refs/tags/1.0.1.tar.gz"
	giturlsub_cmd_tar := exec.Command("curl", "-L", giturlsub_tar_url, "-o", "package.tar.gz")
	err := giturlsub_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	giturlsub_zip_url := "https://github.com/gosuri/git-url-sub/archive/refs/tags/1.0.1.zip"
	giturlsub_cmd_zip := exec.Command("curl", "-L", giturlsub_zip_url, "-o", "package.zip")
	err = giturlsub_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	giturlsub_bin_url := "https://github.com/gosuri/git-url-sub/archive/refs/tags/1.0.1.bin"
	giturlsub_cmd_bin := exec.Command("curl", "-L", giturlsub_bin_url, "-o", "binary.bin")
	err = giturlsub_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	giturlsub_src_url := "https://github.com/gosuri/git-url-sub/archive/refs/tags/1.0.1.src.tar.gz"
	giturlsub_cmd_src := exec.Command("curl", "-L", giturlsub_src_url, "-o", "source.tar.gz")
	err = giturlsub_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	giturlsub_cmd_direct := exec.Command("./binary")
	err = giturlsub_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
