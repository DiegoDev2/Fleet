package main

// Libpgm - Implements the PGM reliable multicast protocol
// Homepage: https://github.com/steve-o/openpgm

import (
	"fmt"
	
	"os/exec"
)

func installLibpgm() {
	// Método 1: Descargar y extraer .tar.gz
	libpgm_tar_url := "https://github.com/steve-o/openpgm/archive/refs/tags/release-5-3-128.tar.gz"
	libpgm_cmd_tar := exec.Command("curl", "-L", libpgm_tar_url, "-o", "package.tar.gz")
	err := libpgm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libpgm_zip_url := "https://github.com/steve-o/openpgm/archive/refs/tags/release-5-3-128.zip"
	libpgm_cmd_zip := exec.Command("curl", "-L", libpgm_zip_url, "-o", "package.zip")
	err = libpgm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libpgm_bin_url := "https://github.com/steve-o/openpgm/archive/refs/tags/release-5-3-128.bin"
	libpgm_cmd_bin := exec.Command("curl", "-L", libpgm_bin_url, "-o", "binary.bin")
	err = libpgm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libpgm_src_url := "https://github.com/steve-o/openpgm/archive/refs/tags/release-5-3-128.src.tar.gz"
	libpgm_cmd_src := exec.Command("curl", "-L", libpgm_src_url, "-o", "source.tar.gz")
	err = libpgm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libpgm_cmd_direct := exec.Command("./binary")
	err = libpgm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
