package main

// GitMultipush - Push a branch to multiple remotes in one command
// Homepage: https://github.com/gavinbeatty/git-multipush

import (
	"fmt"
	
	"os/exec"
)

func installGitMultipush() {
	// Método 1: Descargar y extraer .tar.gz
	gitmultipush_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/git-multipush/git-multipush-2.3.tar.bz2"
	gitmultipush_cmd_tar := exec.Command("curl", "-L", gitmultipush_tar_url, "-o", "package.tar.gz")
	err := gitmultipush_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitmultipush_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/git-multipush/git-multipush-2.3.tar.bz2"
	gitmultipush_cmd_zip := exec.Command("curl", "-L", gitmultipush_zip_url, "-o", "package.zip")
	err = gitmultipush_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitmultipush_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/git-multipush/git-multipush-2.3.tar.bz2"
	gitmultipush_cmd_bin := exec.Command("curl", "-L", gitmultipush_bin_url, "-o", "binary.bin")
	err = gitmultipush_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitmultipush_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/git-multipush/git-multipush-2.3.tar.bz2"
	gitmultipush_cmd_src := exec.Command("curl", "-L", gitmultipush_src_url, "-o", "source.tar.gz")
	err = gitmultipush_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitmultipush_cmd_direct := exec.Command("./binary")
	err = gitmultipush_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asciidoc")
exec.Command("latte", "install", "asciidoc")
}
