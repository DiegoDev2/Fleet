package main

// Libtool - Generic library support script
// Homepage: https://www.gnu.org/software/libtool/

import (
	"fmt"
	
	"os/exec"
)

func installLibtool() {
	// Método 1: Descargar y extraer .tar.gz
	libtool_tar_url := "https://ftp.gnu.org/gnu/libtool/libtool-2.4.7.tar.xz"
	libtool_cmd_tar := exec.Command("curl", "-L", libtool_tar_url, "-o", "package.tar.gz")
	err := libtool_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libtool_zip_url := "https://ftp.gnu.org/gnu/libtool/libtool-2.4.7.tar.xz"
	libtool_cmd_zip := exec.Command("curl", "-L", libtool_zip_url, "-o", "package.zip")
	err = libtool_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libtool_bin_url := "https://ftp.gnu.org/gnu/libtool/libtool-2.4.7.tar.xz"
	libtool_cmd_bin := exec.Command("curl", "-L", libtool_bin_url, "-o", "binary.bin")
	err = libtool_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libtool_src_url := "https://ftp.gnu.org/gnu/libtool/libtool-2.4.7.tar.xz"
	libtool_cmd_src := exec.Command("curl", "-L", libtool_src_url, "-o", "source.tar.gz")
	err = libtool_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libtool_cmd_direct := exec.Command("./binary")
	err = libtool_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: m4")
	exec.Command("latte", "install", "m4").Run()
}
