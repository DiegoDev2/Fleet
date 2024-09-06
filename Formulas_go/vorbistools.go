package main

// VorbisTools - Ogg Vorbis CODEC tools
// Homepage: https://github.com/xiph/vorbis-tools

import (
	"fmt"
	
	"os/exec"
)

func installVorbisTools() {
	// Método 1: Descargar y extraer .tar.gz
	vorbistools_tar_url := "https://downloads.xiph.org/releases/vorbis/vorbis-tools-1.4.2.tar.gz"
	vorbistools_cmd_tar := exec.Command("curl", "-L", vorbistools_tar_url, "-o", "package.tar.gz")
	err := vorbistools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vorbistools_zip_url := "https://downloads.xiph.org/releases/vorbis/vorbis-tools-1.4.2.zip"
	vorbistools_cmd_zip := exec.Command("curl", "-L", vorbistools_zip_url, "-o", "package.zip")
	err = vorbistools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vorbistools_bin_url := "https://downloads.xiph.org/releases/vorbis/vorbis-tools-1.4.2.bin"
	vorbistools_cmd_bin := exec.Command("curl", "-L", vorbistools_bin_url, "-o", "binary.bin")
	err = vorbistools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vorbistools_src_url := "https://downloads.xiph.org/releases/vorbis/vorbis-tools-1.4.2.src.tar.gz"
	vorbistools_cmd_src := exec.Command("curl", "-L", vorbistools_src_url, "-o", "source.tar.gz")
	err = vorbistools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vorbistools_cmd_direct := exec.Command("./binary")
	err = vorbistools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: flac")
exec.Command("latte", "install", "flac")
	fmt.Println("Instalando dependencia: libao")
exec.Command("latte", "install", "libao")
	fmt.Println("Instalando dependencia: libogg")
exec.Command("latte", "install", "libogg")
	fmt.Println("Instalando dependencia: libvorbis")
exec.Command("latte", "install", "libvorbis")
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
