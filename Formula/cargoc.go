package main

// CargoC - Helper program to build and install c-like libraries
// Homepage: https://github.com/lu-zero/cargo-c

import (
	"fmt"
	
	"os/exec"
)

func installCargoC() {
	// Método 1: Descargar y extraer .tar.gz
	cargoc_tar_url := "https://github.com/lu-zero/cargo-c/archive/refs/tags/v0.10.4.tar.gz"
	cargoc_cmd_tar := exec.Command("curl", "-L", cargoc_tar_url, "-o", "package.tar.gz")
	err := cargoc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cargoc_zip_url := "https://github.com/lu-zero/cargo-c/archive/refs/tags/v0.10.4.zip"
	cargoc_cmd_zip := exec.Command("curl", "-L", cargoc_zip_url, "-o", "package.zip")
	err = cargoc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cargoc_bin_url := "https://github.com/lu-zero/cargo-c/archive/refs/tags/v0.10.4.bin"
	cargoc_cmd_bin := exec.Command("curl", "-L", cargoc_bin_url, "-o", "binary.bin")
	err = cargoc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cargoc_src_url := "https://github.com/lu-zero/cargo-c/archive/refs/tags/v0.10.4.src.tar.gz"
	cargoc_cmd_src := exec.Command("curl", "-L", cargoc_src_url, "-o", "source.tar.gz")
	err = cargoc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cargoc_cmd_direct := exec.Command("./binary")
	err = cargoc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: curl")
	exec.Command("latte", "install", "curl").Run()
	fmt.Println("Instalando dependencia: libgit2")
	exec.Command("latte", "install", "libgit2").Run()
	fmt.Println("Instalando dependencia: libssh2")
	exec.Command("latte", "install", "libssh2").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
