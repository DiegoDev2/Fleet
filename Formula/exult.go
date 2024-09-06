package main

// Exult - Recreation of Ultima 7
// Homepage: https://exult.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installExult() {
	// Método 1: Descargar y extraer .tar.gz
	exult_tar_url := "https://github.com/exult/exult/archive/refs/tags/v1.8.tar.gz"
	exult_cmd_tar := exec.Command("curl", "-L", exult_tar_url, "-o", "package.tar.gz")
	err := exult_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	exult_zip_url := "https://github.com/exult/exult/archive/refs/tags/v1.8.zip"
	exult_cmd_zip := exec.Command("curl", "-L", exult_zip_url, "-o", "package.zip")
	err = exult_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	exult_bin_url := "https://github.com/exult/exult/archive/refs/tags/v1.8.bin"
	exult_cmd_bin := exec.Command("curl", "-L", exult_bin_url, "-o", "binary.bin")
	err = exult_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	exult_src_url := "https://github.com/exult/exult/archive/refs/tags/v1.8.src.tar.gz"
	exult_cmd_src := exec.Command("curl", "-L", exult_src_url, "-o", "source.tar.gz")
	err = exult_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	exult_cmd_direct := exec.Command("./binary")
	err = exult_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libogg")
	exec.Command("latte", "install", "libogg").Run()
	fmt.Println("Instalando dependencia: libvorbis")
	exec.Command("latte", "install", "libvorbis").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
}
