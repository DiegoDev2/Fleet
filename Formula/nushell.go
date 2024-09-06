package main

// Nushell - Modern shell for the GitHub era
// Homepage: https://www.nushell.sh

import (
	"fmt"
	
	"os/exec"
)

func installNushell() {
	// Método 1: Descargar y extraer .tar.gz
	nushell_tar_url := "https://github.com/nushell/nushell/archive/refs/tags/0.97.1.tar.gz"
	nushell_cmd_tar := exec.Command("curl", "-L", nushell_tar_url, "-o", "package.tar.gz")
	err := nushell_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nushell_zip_url := "https://github.com/nushell/nushell/archive/refs/tags/0.97.1.zip"
	nushell_cmd_zip := exec.Command("curl", "-L", nushell_zip_url, "-o", "package.zip")
	err = nushell_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nushell_bin_url := "https://github.com/nushell/nushell/archive/refs/tags/0.97.1.bin"
	nushell_cmd_bin := exec.Command("curl", "-L", nushell_bin_url, "-o", "binary.bin")
	err = nushell_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nushell_src_url := "https://github.com/nushell/nushell/archive/refs/tags/0.97.1.src.tar.gz"
	nushell_cmd_src := exec.Command("curl", "-L", nushell_src_url, "-o", "source.tar.gz")
	err = nushell_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nushell_cmd_direct := exec.Command("./binary")
	err = nushell_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libgit2")
	exec.Command("latte", "install", "libgit2").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxcb")
	exec.Command("latte", "install", "libxcb").Run()
}
