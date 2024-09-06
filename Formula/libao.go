package main

// Libao - Cross-platform Audio Library
// Homepage: https://www.xiph.org/ao/

import (
	"fmt"
	
	"os/exec"
)

func installLibao() {
	// Método 1: Descargar y extraer .tar.gz
	libao_tar_url := "https://github.com/xiph/libao/archive/refs/tags/1.2.2.tar.gz"
	libao_cmd_tar := exec.Command("curl", "-L", libao_tar_url, "-o", "package.tar.gz")
	err := libao_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libao_zip_url := "https://github.com/xiph/libao/archive/refs/tags/1.2.2.zip"
	libao_cmd_zip := exec.Command("curl", "-L", libao_zip_url, "-o", "package.zip")
	err = libao_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libao_bin_url := "https://github.com/xiph/libao/archive/refs/tags/1.2.2.bin"
	libao_cmd_bin := exec.Command("curl", "-L", libao_bin_url, "-o", "binary.bin")
	err = libao_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libao_src_url := "https://github.com/xiph/libao/archive/refs/tags/1.2.2.src.tar.gz"
	libao_cmd_src := exec.Command("curl", "-L", libao_src_url, "-o", "source.tar.gz")
	err = libao_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libao_cmd_direct := exec.Command("./binary")
	err = libao_cmd_direct.Run()
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
}
