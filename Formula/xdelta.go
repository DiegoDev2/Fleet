package main

// Xdelta - Binary diff, differential compression tools
// Homepage: http://xdelta.org

import (
	"fmt"
	
	"os/exec"
)

func installXdelta() {
	// Método 1: Descargar y extraer .tar.gz
	xdelta_tar_url := "https://github.com/jmacd/xdelta/archive/refs/tags/v3.1.0.tar.gz"
	xdelta_cmd_tar := exec.Command("curl", "-L", xdelta_tar_url, "-o", "package.tar.gz")
	err := xdelta_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xdelta_zip_url := "https://github.com/jmacd/xdelta/archive/refs/tags/v3.1.0.zip"
	xdelta_cmd_zip := exec.Command("curl", "-L", xdelta_zip_url, "-o", "package.zip")
	err = xdelta_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xdelta_bin_url := "https://github.com/jmacd/xdelta/archive/refs/tags/v3.1.0.bin"
	xdelta_cmd_bin := exec.Command("curl", "-L", xdelta_bin_url, "-o", "binary.bin")
	err = xdelta_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xdelta_src_url := "https://github.com/jmacd/xdelta/archive/refs/tags/v3.1.0.src.tar.gz"
	xdelta_cmd_src := exec.Command("curl", "-L", xdelta_src_url, "-o", "source.tar.gz")
	err = xdelta_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xdelta_cmd_direct := exec.Command("./binary")
	err = xdelta_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
}
