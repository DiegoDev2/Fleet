package main

// Librcsc - RoboCup Soccer Simulator library
// Homepage: https://github.com/helios-base/librcsc

import (
	"fmt"
	
	"os/exec"
)

func installLibrcsc() {
	// Método 1: Descargar y extraer .tar.gz
	librcsc_tar_url := "https://github.com/helios-base/librcsc/archive/refs/tags/rc2024.tar.gz"
	librcsc_cmd_tar := exec.Command("curl", "-L", librcsc_tar_url, "-o", "package.tar.gz")
	err := librcsc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	librcsc_zip_url := "https://github.com/helios-base/librcsc/archive/refs/tags/rc2024.zip"
	librcsc_cmd_zip := exec.Command("curl", "-L", librcsc_zip_url, "-o", "package.zip")
	err = librcsc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	librcsc_bin_url := "https://github.com/helios-base/librcsc/archive/refs/tags/rc2024.bin"
	librcsc_cmd_bin := exec.Command("curl", "-L", librcsc_bin_url, "-o", "binary.bin")
	err = librcsc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	librcsc_src_url := "https://github.com/helios-base/librcsc/archive/refs/tags/rc2024.src.tar.gz"
	librcsc_cmd_src := exec.Command("curl", "-L", librcsc_src_url, "-o", "source.tar.gz")
	err = librcsc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	librcsc_cmd_direct := exec.Command("./binary")
	err = librcsc_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
}
