package main

// Libgit2 - C library of Git core methods that is re-entrant and linkable
// Homepage: https://libgit2.github.com/

import (
	"fmt"
	
	"os/exec"
)

func installLibgit2() {
	// Método 1: Descargar y extraer .tar.gz
	libgit2_tar_url := "https://github.com/libgit2/libgit2/archive/refs/tags/v1.8.1.tar.gz"
	libgit2_cmd_tar := exec.Command("curl", "-L", libgit2_tar_url, "-o", "package.tar.gz")
	err := libgit2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libgit2_zip_url := "https://github.com/libgit2/libgit2/archive/refs/tags/v1.8.1.zip"
	libgit2_cmd_zip := exec.Command("curl", "-L", libgit2_zip_url, "-o", "package.zip")
	err = libgit2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libgit2_bin_url := "https://github.com/libgit2/libgit2/archive/refs/tags/v1.8.1.bin"
	libgit2_cmd_bin := exec.Command("curl", "-L", libgit2_bin_url, "-o", "binary.bin")
	err = libgit2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libgit2_src_url := "https://github.com/libgit2/libgit2/archive/refs/tags/v1.8.1.src.tar.gz"
	libgit2_cmd_src := exec.Command("curl", "-L", libgit2_src_url, "-o", "source.tar.gz")
	err = libgit2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libgit2_cmd_direct := exec.Command("./binary")
	err = libgit2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libssh2")
	exec.Command("latte", "install", "libssh2").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
