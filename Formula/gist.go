package main

// Gist - Command-line utility for uploading Gists
// Homepage: https://github.com/defunkt/gist

import (
	"fmt"
	
	"os/exec"
)

func installGist() {
	// Método 1: Descargar y extraer .tar.gz
	gist_tar_url := "https://github.com/defunkt/gist/archive/refs/tags/v6.0.0.tar.gz"
	gist_cmd_tar := exec.Command("curl", "-L", gist_tar_url, "-o", "package.tar.gz")
	err := gist_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gist_zip_url := "https://github.com/defunkt/gist/archive/refs/tags/v6.0.0.zip"
	gist_cmd_zip := exec.Command("curl", "-L", gist_zip_url, "-o", "package.zip")
	err = gist_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gist_bin_url := "https://github.com/defunkt/gist/archive/refs/tags/v6.0.0.bin"
	gist_cmd_bin := exec.Command("curl", "-L", gist_bin_url, "-o", "binary.bin")
	err = gist_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gist_src_url := "https://github.com/defunkt/gist/archive/refs/tags/v6.0.0.src.tar.gz"
	gist_cmd_src := exec.Command("curl", "-L", gist_src_url, "-o", "source.tar.gz")
	err = gist_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gist_cmd_direct := exec.Command("./binary")
	err = gist_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
