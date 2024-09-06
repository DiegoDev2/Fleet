package main

// Libgit2AT17 - C library of Git core methods that is re-entrant and linkable
// Homepage: https://libgit2.github.com/

import (
	"fmt"
	
	"os/exec"
)

func installLibgit2AT17() {
	// Método 1: Descargar y extraer .tar.gz
	libgit2at17_tar_url := "https://github.com/libgit2/libgit2/archive/refs/tags/v1.7.2.tar.gz"
	libgit2at17_cmd_tar := exec.Command("curl", "-L", libgit2at17_tar_url, "-o", "package.tar.gz")
	err := libgit2at17_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libgit2at17_zip_url := "https://github.com/libgit2/libgit2/archive/refs/tags/v1.7.2.zip"
	libgit2at17_cmd_zip := exec.Command("curl", "-L", libgit2at17_zip_url, "-o", "package.zip")
	err = libgit2at17_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libgit2at17_bin_url := "https://github.com/libgit2/libgit2/archive/refs/tags/v1.7.2.bin"
	libgit2at17_cmd_bin := exec.Command("curl", "-L", libgit2at17_bin_url, "-o", "binary.bin")
	err = libgit2at17_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libgit2at17_src_url := "https://github.com/libgit2/libgit2/archive/refs/tags/v1.7.2.src.tar.gz"
	libgit2at17_cmd_src := exec.Command("curl", "-L", libgit2at17_src_url, "-o", "source.tar.gz")
	err = libgit2at17_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libgit2at17_cmd_direct := exec.Command("./binary")
	err = libgit2at17_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libssh2")
exec.Command("latte", "install", "libssh2")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
