package main

// Libvorbis - Vorbis General Audio Compression Codec
// Homepage: https://xiph.org/vorbis/

import (
	"fmt"
	
	"os/exec"
)

func installLibvorbis() {
	// Método 1: Descargar y extraer .tar.gz
	libvorbis_tar_url := "https://downloads.xiph.org/releases/vorbis/libvorbis-1.3.7.tar.xz"
	libvorbis_cmd_tar := exec.Command("curl", "-L", libvorbis_tar_url, "-o", "package.tar.gz")
	err := libvorbis_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libvorbis_zip_url := "https://downloads.xiph.org/releases/vorbis/libvorbis-1.3.7.tar.xz"
	libvorbis_cmd_zip := exec.Command("curl", "-L", libvorbis_zip_url, "-o", "package.zip")
	err = libvorbis_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libvorbis_bin_url := "https://downloads.xiph.org/releases/vorbis/libvorbis-1.3.7.tar.xz"
	libvorbis_cmd_bin := exec.Command("curl", "-L", libvorbis_bin_url, "-o", "binary.bin")
	err = libvorbis_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libvorbis_src_url := "https://downloads.xiph.org/releases/vorbis/libvorbis-1.3.7.tar.xz"
	libvorbis_cmd_src := exec.Command("curl", "-L", libvorbis_src_url, "-o", "source.tar.gz")
	err = libvorbis_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libvorbis_cmd_direct := exec.Command("./binary")
	err = libvorbis_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libogg")
	exec.Command("latte", "install", "libogg").Run()
}
