package main

// Augeas - Configuration editing tool and API
// Homepage: https://augeas.net/

import (
	"fmt"
	
	"os/exec"
)

func installAugeas() {
	// Método 1: Descargar y extraer .tar.gz
	augeas_tar_url := "https://github.com/hercules-team/augeas/releases/download/release-1.14.1/augeas-1.14.1.tar.gz"
	augeas_cmd_tar := exec.Command("curl", "-L", augeas_tar_url, "-o", "package.tar.gz")
	err := augeas_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	augeas_zip_url := "https://github.com/hercules-team/augeas/releases/download/release-1.14.1/augeas-1.14.1.zip"
	augeas_cmd_zip := exec.Command("curl", "-L", augeas_zip_url, "-o", "package.zip")
	err = augeas_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	augeas_bin_url := "https://github.com/hercules-team/augeas/releases/download/release-1.14.1/augeas-1.14.1.bin"
	augeas_cmd_bin := exec.Command("curl", "-L", augeas_bin_url, "-o", "binary.bin")
	err = augeas_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	augeas_src_url := "https://github.com/hercules-team/augeas/releases/download/release-1.14.1/augeas-1.14.1.src.tar.gz"
	augeas_cmd_src := exec.Command("curl", "-L", augeas_src_url, "-o", "source.tar.gz")
	err = augeas_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	augeas_cmd_direct := exec.Command("./binary")
	err = augeas_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: bison")
exec.Command("latte", "install", "bison")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
