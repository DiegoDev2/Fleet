package main

// Libshout - Data and connectivity library for the icecast server
// Homepage: https://icecast.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibshout() {
	// Método 1: Descargar y extraer .tar.gz
	libshout_tar_url := "https://downloads.xiph.org/releases/libshout/libshout-2.4.6.tar.gz"
	libshout_cmd_tar := exec.Command("curl", "-L", libshout_tar_url, "-o", "package.tar.gz")
	err := libshout_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libshout_zip_url := "https://downloads.xiph.org/releases/libshout/libshout-2.4.6.zip"
	libshout_cmd_zip := exec.Command("curl", "-L", libshout_zip_url, "-o", "package.zip")
	err = libshout_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libshout_bin_url := "https://downloads.xiph.org/releases/libshout/libshout-2.4.6.bin"
	libshout_cmd_bin := exec.Command("curl", "-L", libshout_bin_url, "-o", "binary.bin")
	err = libshout_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libshout_src_url := "https://downloads.xiph.org/releases/libshout/libshout-2.4.6.src.tar.gz"
	libshout_cmd_src := exec.Command("curl", "-L", libshout_src_url, "-o", "source.tar.gz")
	err = libshout_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libshout_cmd_direct := exec.Command("./binary")
	err = libshout_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libogg")
	exec.Command("latte", "install", "libogg").Run()
	fmt.Println("Instalando dependencia: libvorbis")
	exec.Command("latte", "install", "libvorbis").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: speex")
	exec.Command("latte", "install", "speex").Run()
	fmt.Println("Instalando dependencia: theora")
	exec.Command("latte", "install", "theora").Run()
}
