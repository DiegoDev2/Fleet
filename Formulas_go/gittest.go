package main

// GitTest - Run tests on each distinct tree in a revision list
// Homepage: https://github.com/spotify/git-test

import (
	"fmt"
	
	"os/exec"
)

func installGitTest() {
	// Método 1: Descargar y extraer .tar.gz
	gittest_tar_url := "https://github.com/spotify/git-test/archive/refs/tags/v1.0.4.tar.gz"
	gittest_cmd_tar := exec.Command("curl", "-L", gittest_tar_url, "-o", "package.tar.gz")
	err := gittest_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gittest_zip_url := "https://github.com/spotify/git-test/archive/refs/tags/v1.0.4.zip"
	gittest_cmd_zip := exec.Command("curl", "-L", gittest_zip_url, "-o", "package.zip")
	err = gittest_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gittest_bin_url := "https://github.com/spotify/git-test/archive/refs/tags/v1.0.4.bin"
	gittest_cmd_bin := exec.Command("curl", "-L", gittest_bin_url, "-o", "binary.bin")
	err = gittest_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gittest_src_url := "https://github.com/spotify/git-test/archive/refs/tags/v1.0.4.src.tar.gz"
	gittest_cmd_src := exec.Command("curl", "-L", gittest_src_url, "-o", "source.tar.gz")
	err = gittest_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gittest_cmd_direct := exec.Command("./binary")
	err = gittest_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
