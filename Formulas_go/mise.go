package main

// Mise - Polyglot runtime manager (asdf rust clone)
// Homepage: https://mise.jdx.dev/

import (
	"fmt"
	
	"os/exec"
)

func installMise() {
	// Método 1: Descargar y extraer .tar.gz
	mise_tar_url := "https://github.com/jdx/mise/archive/refs/tags/v2024.9.0.tar.gz"
	mise_cmd_tar := exec.Command("curl", "-L", mise_tar_url, "-o", "package.tar.gz")
	err := mise_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mise_zip_url := "https://github.com/jdx/mise/archive/refs/tags/v2024.9.0.zip"
	mise_cmd_zip := exec.Command("curl", "-L", mise_zip_url, "-o", "package.zip")
	err = mise_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mise_bin_url := "https://github.com/jdx/mise/archive/refs/tags/v2024.9.0.bin"
	mise_cmd_bin := exec.Command("curl", "-L", mise_bin_url, "-o", "binary.bin")
	err = mise_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mise_src_url := "https://github.com/jdx/mise/archive/refs/tags/v2024.9.0.src.tar.gz"
	mise_cmd_src := exec.Command("curl", "-L", mise_src_url, "-o", "source.tar.gz")
	err = mise_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mise_cmd_direct := exec.Command("./binary")
	err = mise_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: usage")
exec.Command("latte", "install", "usage")
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
}
