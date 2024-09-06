package main

// Poke - Extensible editor for structured binary data
// Homepage: https://jemarch.net/poke

import (
	"fmt"
	
	"os/exec"
)

func installPoke() {
	// Método 1: Descargar y extraer .tar.gz
	poke_tar_url := "https://ftp.gnu.org/gnu/poke/poke-4.2.tar.gz"
	poke_cmd_tar := exec.Command("curl", "-L", poke_tar_url, "-o", "package.tar.gz")
	err := poke_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	poke_zip_url := "https://ftp.gnu.org/gnu/poke/poke-4.2.zip"
	poke_cmd_zip := exec.Command("curl", "-L", poke_zip_url, "-o", "package.zip")
	err = poke_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	poke_bin_url := "https://ftp.gnu.org/gnu/poke/poke-4.2.bin"
	poke_cmd_bin := exec.Command("curl", "-L", poke_bin_url, "-o", "binary.bin")
	err = poke_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	poke_src_url := "https://ftp.gnu.org/gnu/poke/poke-4.2.src.tar.gz"
	poke_cmd_src := exec.Command("curl", "-L", poke_src_url, "-o", "source.tar.gz")
	err = poke_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	poke_cmd_direct := exec.Command("./binary")
	err = poke_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: help2man")
	exec.Command("latte", "install", "help2man").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: bdw-gc")
	exec.Command("latte", "install", "bdw-gc").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
