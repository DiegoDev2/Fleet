package main

// Libmusicbrainz - MusicBrainz Client Library
// Homepage: https://musicbrainz.org/doc/libmusicbrainz

import (
	"fmt"
	
	"os/exec"
)

func installLibmusicbrainz() {
	// Método 1: Descargar y extraer .tar.gz
	libmusicbrainz_tar_url := "https://github.com/metabrainz/libmusicbrainz/releases/download/release-5.1.0/libmusicbrainz-5.1.0.tar.gz"
	libmusicbrainz_cmd_tar := exec.Command("curl", "-L", libmusicbrainz_tar_url, "-o", "package.tar.gz")
	err := libmusicbrainz_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libmusicbrainz_zip_url := "https://github.com/metabrainz/libmusicbrainz/releases/download/release-5.1.0/libmusicbrainz-5.1.0.zip"
	libmusicbrainz_cmd_zip := exec.Command("curl", "-L", libmusicbrainz_zip_url, "-o", "package.zip")
	err = libmusicbrainz_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libmusicbrainz_bin_url := "https://github.com/metabrainz/libmusicbrainz/releases/download/release-5.1.0/libmusicbrainz-5.1.0.bin"
	libmusicbrainz_cmd_bin := exec.Command("curl", "-L", libmusicbrainz_bin_url, "-o", "binary.bin")
	err = libmusicbrainz_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libmusicbrainz_src_url := "https://github.com/metabrainz/libmusicbrainz/releases/download/release-5.1.0/libmusicbrainz-5.1.0.src.tar.gz"
	libmusicbrainz_cmd_src := exec.Command("curl", "-L", libmusicbrainz_src_url, "-o", "source.tar.gz")
	err = libmusicbrainz_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libmusicbrainz_cmd_direct := exec.Command("./binary")
	err = libmusicbrainz_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: neon")
	exec.Command("latte", "install", "neon").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
