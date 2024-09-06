package main

// Cconv - Iconv based simplified-traditional Chinese conversion tool
// Homepage: https://github.com/xiaoyjy/cconv

import (
	"fmt"
	
	"os/exec"
)

func installCconv() {
	// Método 1: Descargar y extraer .tar.gz
	cconv_tar_url := "https://github.com/xiaoyjy/cconv/archive/refs/tags/v0.6.3.tar.gz"
	cconv_cmd_tar := exec.Command("curl", "-L", cconv_tar_url, "-o", "package.tar.gz")
	err := cconv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cconv_zip_url := "https://github.com/xiaoyjy/cconv/archive/refs/tags/v0.6.3.zip"
	cconv_cmd_zip := exec.Command("curl", "-L", cconv_zip_url, "-o", "package.zip")
	err = cconv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cconv_bin_url := "https://github.com/xiaoyjy/cconv/archive/refs/tags/v0.6.3.bin"
	cconv_cmd_bin := exec.Command("curl", "-L", cconv_bin_url, "-o", "binary.bin")
	err = cconv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cconv_src_url := "https://github.com/xiaoyjy/cconv/archive/refs/tags/v0.6.3.src.tar.gz"
	cconv_cmd_src := exec.Command("curl", "-L", cconv_src_url, "-o", "source.tar.gz")
	err = cconv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cconv_cmd_direct := exec.Command("./binary")
	err = cconv_cmd_direct.Run()
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
}
