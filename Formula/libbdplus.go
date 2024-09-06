package main

// Libbdplus - Implements the BD+ System Specifications
// Homepage: https://www.videolan.org/developers/libbdplus.html

import (
	"fmt"
	
	"os/exec"
)

func installLibbdplus() {
	// Método 1: Descargar y extraer .tar.gz
	libbdplus_tar_url := "https://get.videolan.org/libbdplus/0.2.0/libbdplus-0.2.0.tar.bz2"
	libbdplus_cmd_tar := exec.Command("curl", "-L", libbdplus_tar_url, "-o", "package.tar.gz")
	err := libbdplus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libbdplus_zip_url := "https://get.videolan.org/libbdplus/0.2.0/libbdplus-0.2.0.tar.bz2"
	libbdplus_cmd_zip := exec.Command("curl", "-L", libbdplus_zip_url, "-o", "package.zip")
	err = libbdplus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libbdplus_bin_url := "https://get.videolan.org/libbdplus/0.2.0/libbdplus-0.2.0.tar.bz2"
	libbdplus_cmd_bin := exec.Command("curl", "-L", libbdplus_bin_url, "-o", "binary.bin")
	err = libbdplus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libbdplus_src_url := "https://get.videolan.org/libbdplus/0.2.0/libbdplus-0.2.0.tar.bz2"
	libbdplus_cmd_src := exec.Command("curl", "-L", libbdplus_src_url, "-o", "source.tar.gz")
	err = libbdplus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libbdplus_cmd_direct := exec.Command("./binary")
	err = libbdplus_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libgcrypt")
	exec.Command("latte", "install", "libgcrypt").Run()
}
