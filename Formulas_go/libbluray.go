package main

// Libbluray - Blu-Ray disc playback library for media players like VLC
// Homepage: https://www.videolan.org/developers/libbluray.html

import (
	"fmt"
	
	"os/exec"
)

func installLibbluray() {
	// Método 1: Descargar y extraer .tar.gz
	libbluray_tar_url := "https://download.videolan.org/videolan/libbluray/1.3.4/libbluray-1.3.4.tar.bz2"
	libbluray_cmd_tar := exec.Command("curl", "-L", libbluray_tar_url, "-o", "package.tar.gz")
	err := libbluray_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libbluray_zip_url := "https://download.videolan.org/videolan/libbluray/1.3.4/libbluray-1.3.4.tar.bz2"
	libbluray_cmd_zip := exec.Command("curl", "-L", libbluray_zip_url, "-o", "package.zip")
	err = libbluray_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libbluray_bin_url := "https://download.videolan.org/videolan/libbluray/1.3.4/libbluray-1.3.4.tar.bz2"
	libbluray_cmd_bin := exec.Command("curl", "-L", libbluray_bin_url, "-o", "binary.bin")
	err = libbluray_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libbluray_src_url := "https://download.videolan.org/videolan/libbluray/1.3.4/libbluray-1.3.4.tar.bz2"
	libbluray_cmd_src := exec.Command("curl", "-L", libbluray_src_url, "-o", "source.tar.gz")
	err = libbluray_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libbluray_cmd_direct := exec.Command("./binary")
	err = libbluray_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
}
