package main

// Uniutils - Manipulate and analyze Unicode text
// Homepage: https://billposer.org/Software/unidesc.html

import (
	"fmt"
	
	"os/exec"
)

func installUniutils() {
	// Método 1: Descargar y extraer .tar.gz
	uniutils_tar_url := "https://billposer.org/Software/Downloads/uniutils-2.28.tar.gz"
	uniutils_cmd_tar := exec.Command("curl", "-L", uniutils_tar_url, "-o", "package.tar.gz")
	err := uniutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	uniutils_zip_url := "https://billposer.org/Software/Downloads/uniutils-2.28.zip"
	uniutils_cmd_zip := exec.Command("curl", "-L", uniutils_zip_url, "-o", "package.zip")
	err = uniutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	uniutils_bin_url := "https://billposer.org/Software/Downloads/uniutils-2.28.bin"
	uniutils_cmd_bin := exec.Command("curl", "-L", uniutils_bin_url, "-o", "binary.bin")
	err = uniutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	uniutils_src_url := "https://billposer.org/Software/Downloads/uniutils-2.28.src.tar.gz"
	uniutils_cmd_src := exec.Command("curl", "-L", uniutils_src_url, "-o", "source.tar.gz")
	err = uniutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	uniutils_cmd_direct := exec.Command("./binary")
	err = uniutils_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
