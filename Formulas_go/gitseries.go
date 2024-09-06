package main

// GitSeries - Track changes to a patch series over time
// Homepage: https://github.com/git-series/git-series

import (
	"fmt"
	
	"os/exec"
)

func installGitSeries() {
	// Método 1: Descargar y extraer .tar.gz
	gitseries_tar_url := "https://github.com/git-series/git-series/archive/refs/tags/0.9.1.tar.gz"
	gitseries_cmd_tar := exec.Command("curl", "-L", gitseries_tar_url, "-o", "package.tar.gz")
	err := gitseries_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitseries_zip_url := "https://github.com/git-series/git-series/archive/refs/tags/0.9.1.zip"
	gitseries_cmd_zip := exec.Command("curl", "-L", gitseries_zip_url, "-o", "package.zip")
	err = gitseries_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitseries_bin_url := "https://github.com/git-series/git-series/archive/refs/tags/0.9.1.bin"
	gitseries_cmd_bin := exec.Command("curl", "-L", gitseries_bin_url, "-o", "binary.bin")
	err = gitseries_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitseries_src_url := "https://github.com/git-series/git-series/archive/refs/tags/0.9.1.src.tar.gz"
	gitseries_cmd_src := exec.Command("curl", "-L", gitseries_src_url, "-o", "source.tar.gz")
	err = gitseries_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitseries_cmd_direct := exec.Command("./binary")
	err = gitseries_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: libgit2")
exec.Command("latte", "install", "libgit2")
	fmt.Println("Instalando dependencia: libssh2")
exec.Command("latte", "install", "libssh2")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
