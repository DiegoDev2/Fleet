package main

// Libmp3splt - Utility library to split mp3, ogg, and FLAC files
// Homepage: https://mp3splt.sourceforge.net/mp3splt_page/home.php

import (
	"fmt"
	
	"os/exec"
)

func installLibmp3splt() {
	// Método 1: Descargar y extraer .tar.gz
	libmp3splt_tar_url := "https://downloads.sourceforge.net/project/mp3splt/libmp3splt/0.9.2/libmp3splt-0.9.2.tar.gz"
	libmp3splt_cmd_tar := exec.Command("curl", "-L", libmp3splt_tar_url, "-o", "package.tar.gz")
	err := libmp3splt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libmp3splt_zip_url := "https://downloads.sourceforge.net/project/mp3splt/libmp3splt/0.9.2/libmp3splt-0.9.2.zip"
	libmp3splt_cmd_zip := exec.Command("curl", "-L", libmp3splt_zip_url, "-o", "package.zip")
	err = libmp3splt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libmp3splt_bin_url := "https://downloads.sourceforge.net/project/mp3splt/libmp3splt/0.9.2/libmp3splt-0.9.2.bin"
	libmp3splt_cmd_bin := exec.Command("curl", "-L", libmp3splt_bin_url, "-o", "binary.bin")
	err = libmp3splt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libmp3splt_src_url := "https://downloads.sourceforge.net/project/mp3splt/libmp3splt/0.9.2/libmp3splt-0.9.2.src.tar.gz"
	libmp3splt_cmd_src := exec.Command("curl", "-L", libmp3splt_src_url, "-o", "source.tar.gz")
	err = libmp3splt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libmp3splt_cmd_direct := exec.Command("./binary")
	err = libmp3splt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: flac")
	exec.Command("latte", "install", "flac").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: libid3tag")
	exec.Command("latte", "install", "libid3tag").Run()
	fmt.Println("Instalando dependencia: libogg")
	exec.Command("latte", "install", "libogg").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: libvorbis")
	exec.Command("latte", "install", "libvorbis").Run()
	fmt.Println("Instalando dependencia: mad")
	exec.Command("latte", "install", "mad").Run()
	fmt.Println("Instalando dependencia: pcre")
	exec.Command("latte", "install", "pcre").Run()
}
