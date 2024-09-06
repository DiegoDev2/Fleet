package main

// Flac - Free lossless audio codec
// Homepage: https://xiph.org/flac/

import (
	"fmt"
	
	"os/exec"
)

func installFlac() {
	// Método 1: Descargar y extraer .tar.gz
	flac_tar_url := "https://downloads.xiph.org/releases/flac/flac-1.4.3.tar.xz"
	flac_cmd_tar := exec.Command("curl", "-L", flac_tar_url, "-o", "package.tar.gz")
	err := flac_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	flac_zip_url := "https://downloads.xiph.org/releases/flac/flac-1.4.3.tar.xz"
	flac_cmd_zip := exec.Command("curl", "-L", flac_zip_url, "-o", "package.zip")
	err = flac_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	flac_bin_url := "https://downloads.xiph.org/releases/flac/flac-1.4.3.tar.xz"
	flac_cmd_bin := exec.Command("curl", "-L", flac_bin_url, "-o", "binary.bin")
	err = flac_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	flac_src_url := "https://downloads.xiph.org/releases/flac/flac-1.4.3.tar.xz"
	flac_cmd_src := exec.Command("curl", "-L", flac_src_url, "-o", "source.tar.gz")
	err = flac_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	flac_cmd_direct := exec.Command("./binary")
	err = flac_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libogg")
exec.Command("latte", "install", "libogg")
}
