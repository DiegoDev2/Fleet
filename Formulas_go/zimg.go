package main

// Zimg - Scaling, colorspace conversion, and dithering library
// Homepage: https://github.com/sekrit-twc/zimg

import (
	"fmt"
	
	"os/exec"
)

func installZimg() {
	// Método 1: Descargar y extraer .tar.gz
	zimg_tar_url := "https://github.com/sekrit-twc/zimg/archive/refs/tags/release-3.0.5.tar.gz"
	zimg_cmd_tar := exec.Command("curl", "-L", zimg_tar_url, "-o", "package.tar.gz")
	err := zimg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zimg_zip_url := "https://github.com/sekrit-twc/zimg/archive/refs/tags/release-3.0.5.zip"
	zimg_cmd_zip := exec.Command("curl", "-L", zimg_zip_url, "-o", "package.zip")
	err = zimg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zimg_bin_url := "https://github.com/sekrit-twc/zimg/archive/refs/tags/release-3.0.5.bin"
	zimg_cmd_bin := exec.Command("curl", "-L", zimg_bin_url, "-o", "binary.bin")
	err = zimg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zimg_src_url := "https://github.com/sekrit-twc/zimg/archive/refs/tags/release-3.0.5.src.tar.gz"
	zimg_cmd_src := exec.Command("curl", "-L", zimg_src_url, "-o", "source.tar.gz")
	err = zimg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zimg_cmd_direct := exec.Command("./binary")
	err = zimg_cmd_direct.Run()
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
}
