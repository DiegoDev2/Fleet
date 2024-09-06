package main

// Liboil - C library of simple functions optimized for various CPUs
// Homepage: https://wiki.freedesktop.org/liboil/

import (
	"fmt"
	
	"os/exec"
)

func installLiboil() {
	// Método 1: Descargar y extraer .tar.gz
	liboil_tar_url := "https://liboil.freedesktop.org/download/liboil-0.3.17.tar.gz"
	liboil_cmd_tar := exec.Command("curl", "-L", liboil_tar_url, "-o", "package.tar.gz")
	err := liboil_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	liboil_zip_url := "https://liboil.freedesktop.org/download/liboil-0.3.17.zip"
	liboil_cmd_zip := exec.Command("curl", "-L", liboil_zip_url, "-o", "package.zip")
	err = liboil_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	liboil_bin_url := "https://liboil.freedesktop.org/download/liboil-0.3.17.bin"
	liboil_cmd_bin := exec.Command("curl", "-L", liboil_bin_url, "-o", "binary.bin")
	err = liboil_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	liboil_src_url := "https://liboil.freedesktop.org/download/liboil-0.3.17.src.tar.gz"
	liboil_cmd_src := exec.Command("curl", "-L", liboil_src_url, "-o", "source.tar.gz")
	err = liboil_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	liboil_cmd_direct := exec.Command("./binary")
	err = liboil_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: gtk-doc")
	exec.Command("latte", "install", "gtk-doc").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
