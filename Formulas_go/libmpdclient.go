package main

// Libmpdclient - Library for MPD in the C, C++, and Objective-C languages
// Homepage: https://www.musicpd.org/libs/libmpdclient/

import (
	"fmt"
	
	"os/exec"
)

func installLibmpdclient() {
	// Método 1: Descargar y extraer .tar.gz
	libmpdclient_tar_url := "https://www.musicpd.org/download/libmpdclient/2/libmpdclient-2.22.tar.xz"
	libmpdclient_cmd_tar := exec.Command("curl", "-L", libmpdclient_tar_url, "-o", "package.tar.gz")
	err := libmpdclient_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libmpdclient_zip_url := "https://www.musicpd.org/download/libmpdclient/2/libmpdclient-2.22.tar.xz"
	libmpdclient_cmd_zip := exec.Command("curl", "-L", libmpdclient_zip_url, "-o", "package.zip")
	err = libmpdclient_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libmpdclient_bin_url := "https://www.musicpd.org/download/libmpdclient/2/libmpdclient-2.22.tar.xz"
	libmpdclient_cmd_bin := exec.Command("curl", "-L", libmpdclient_bin_url, "-o", "binary.bin")
	err = libmpdclient_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libmpdclient_src_url := "https://www.musicpd.org/download/libmpdclient/2/libmpdclient-2.22.tar.xz"
	libmpdclient_cmd_src := exec.Command("curl", "-L", libmpdclient_src_url, "-o", "source.tar.gz")
	err = libmpdclient_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libmpdclient_cmd_direct := exec.Command("./binary")
	err = libmpdclient_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
}
