package main

// Libspectrum - Support library for ZX Spectrum emulator
// Homepage: https://fuse-emulator.sourceforge.net/libspectrum.php

import (
	"fmt"
	
	"os/exec"
)

func installLibspectrum() {
	// Método 1: Descargar y extraer .tar.gz
	libspectrum_tar_url := "https://downloads.sourceforge.net/project/fuse-emulator/libspectrum/1.5.0/libspectrum-1.5.0.tar.gz"
	libspectrum_cmd_tar := exec.Command("curl", "-L", libspectrum_tar_url, "-o", "package.tar.gz")
	err := libspectrum_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libspectrum_zip_url := "https://downloads.sourceforge.net/project/fuse-emulator/libspectrum/1.5.0/libspectrum-1.5.0.zip"
	libspectrum_cmd_zip := exec.Command("curl", "-L", libspectrum_zip_url, "-o", "package.zip")
	err = libspectrum_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libspectrum_bin_url := "https://downloads.sourceforge.net/project/fuse-emulator/libspectrum/1.5.0/libspectrum-1.5.0.bin"
	libspectrum_cmd_bin := exec.Command("curl", "-L", libspectrum_bin_url, "-o", "binary.bin")
	err = libspectrum_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libspectrum_src_url := "https://downloads.sourceforge.net/project/fuse-emulator/libspectrum/1.5.0/libspectrum-1.5.0.src.tar.gz"
	libspectrum_cmd_src := exec.Command("curl", "-L", libspectrum_src_url, "-o", "source.tar.gz")
	err = libspectrum_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libspectrum_cmd_direct := exec.Command("./binary")
	err = libspectrum_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: audiofile")
	exec.Command("latte", "install", "audiofile").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: libgcrypt")
	exec.Command("latte", "install", "libgcrypt").Run()
}
