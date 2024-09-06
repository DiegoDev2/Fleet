package main

// Jupp - Professional screen editor for programmers
// Homepage: http://www.mirbsd.org/jupp.htm

import (
	"fmt"
	
	"os/exec"
)

func installJupp() {
	// Método 1: Descargar y extraer .tar.gz
	jupp_tar_url := "http://www.mirbsd.org/MirOS/dist/jupp/joe-3.1jupp41.tgz"
	jupp_cmd_tar := exec.Command("curl", "-L", jupp_tar_url, "-o", "package.tar.gz")
	err := jupp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jupp_zip_url := "http://www.mirbsd.org/MirOS/dist/jupp/joe-3.1jupp41.tgz"
	jupp_cmd_zip := exec.Command("curl", "-L", jupp_zip_url, "-o", "package.zip")
	err = jupp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jupp_bin_url := "http://www.mirbsd.org/MirOS/dist/jupp/joe-3.1jupp41.tgz"
	jupp_cmd_bin := exec.Command("curl", "-L", jupp_bin_url, "-o", "binary.bin")
	err = jupp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jupp_src_url := "http://www.mirbsd.org/MirOS/dist/jupp/joe-3.1jupp41.tgz"
	jupp_cmd_src := exec.Command("curl", "-L", jupp_src_url, "-o", "source.tar.gz")
	err = jupp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jupp_cmd_direct := exec.Command("./binary")
	err = jupp_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: gnu-sed")
exec.Command("latte", "install", "gnu-sed")
}
