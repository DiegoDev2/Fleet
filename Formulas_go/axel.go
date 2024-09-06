package main

// Axel - Light UNIX download accelerator
// Homepage: https://github.com/axel-download-accelerator/axel

import (
	"fmt"
	
	"os/exec"
)

func installAxel() {
	// Método 1: Descargar y extraer .tar.gz
	axel_tar_url := "https://github.com/axel-download-accelerator/axel/releases/download/v2.17.14/axel-2.17.14.tar.xz"
	axel_cmd_tar := exec.Command("curl", "-L", axel_tar_url, "-o", "package.tar.gz")
	err := axel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	axel_zip_url := "https://github.com/axel-download-accelerator/axel/releases/download/v2.17.14/axel-2.17.14.tar.xz"
	axel_cmd_zip := exec.Command("curl", "-L", axel_zip_url, "-o", "package.zip")
	err = axel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	axel_bin_url := "https://github.com/axel-download-accelerator/axel/releases/download/v2.17.14/axel-2.17.14.tar.xz"
	axel_cmd_bin := exec.Command("curl", "-L", axel_bin_url, "-o", "binary.bin")
	err = axel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	axel_src_url := "https://github.com/axel-download-accelerator/axel/releases/download/v2.17.14/axel-2.17.14.tar.xz"
	axel_cmd_src := exec.Command("curl", "-L", axel_src_url, "-o", "source.tar.gz")
	err = axel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	axel_cmd_direct := exec.Command("./binary")
	err = axel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: autoconf-archive")
exec.Command("latte", "install", "autoconf-archive")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: gawk")
exec.Command("latte", "install", "gawk")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: txt2man")
exec.Command("latte", "install", "txt2man")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
