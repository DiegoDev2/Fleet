package main

// Stgit - Manage Git commits as a stack of patches
// Homepage: https://stacked-git.github.io

import (
	"fmt"
	
	"os/exec"
)

func installStgit() {
	// Método 1: Descargar y extraer .tar.gz
	stgit_tar_url := "https://github.com/stacked-git/stgit/releases/download/v2.4.11/stgit-2.4.11.tar.gz"
	stgit_cmd_tar := exec.Command("curl", "-L", stgit_tar_url, "-o", "package.tar.gz")
	err := stgit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stgit_zip_url := "https://github.com/stacked-git/stgit/releases/download/v2.4.11/stgit-2.4.11.zip"
	stgit_cmd_zip := exec.Command("curl", "-L", stgit_zip_url, "-o", "package.zip")
	err = stgit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stgit_bin_url := "https://github.com/stacked-git/stgit/releases/download/v2.4.11/stgit-2.4.11.bin"
	stgit_cmd_bin := exec.Command("curl", "-L", stgit_bin_url, "-o", "binary.bin")
	err = stgit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stgit_src_url := "https://github.com/stacked-git/stgit/releases/download/v2.4.11/stgit-2.4.11.src.tar.gz"
	stgit_cmd_src := exec.Command("curl", "-L", stgit_src_url, "-o", "source.tar.gz")
	err = stgit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stgit_cmd_direct := exec.Command("./binary")
	err = stgit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asciidoc")
exec.Command("latte", "install", "asciidoc")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: xmlto")
exec.Command("latte", "install", "xmlto")
	fmt.Println("Instalando dependencia: git")
exec.Command("latte", "install", "git")
}
