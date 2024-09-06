package main

// Libdiscid - C library for creating MusicBrainz and freedb disc IDs
// Homepage: https://musicbrainz.org/doc/libdiscid

import (
	"fmt"
	
	"os/exec"
)

func installLibdiscid() {
	// Método 1: Descargar y extraer .tar.gz
	libdiscid_tar_url := "https://ftp.musicbrainz.org/pub/musicbrainz/libdiscid/libdiscid-0.6.4.tar.gz"
	libdiscid_cmd_tar := exec.Command("curl", "-L", libdiscid_tar_url, "-o", "package.tar.gz")
	err := libdiscid_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libdiscid_zip_url := "https://ftp.musicbrainz.org/pub/musicbrainz/libdiscid/libdiscid-0.6.4.zip"
	libdiscid_cmd_zip := exec.Command("curl", "-L", libdiscid_zip_url, "-o", "package.zip")
	err = libdiscid_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libdiscid_bin_url := "https://ftp.musicbrainz.org/pub/musicbrainz/libdiscid/libdiscid-0.6.4.bin"
	libdiscid_cmd_bin := exec.Command("curl", "-L", libdiscid_bin_url, "-o", "binary.bin")
	err = libdiscid_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libdiscid_src_url := "https://ftp.musicbrainz.org/pub/musicbrainz/libdiscid/libdiscid-0.6.4.src.tar.gz"
	libdiscid_cmd_src := exec.Command("curl", "-L", libdiscid_src_url, "-o", "source.tar.gz")
	err = libdiscid_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libdiscid_cmd_direct := exec.Command("./binary")
	err = libdiscid_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
