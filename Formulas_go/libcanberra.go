package main

// Libcanberra - Implementation of XDG Sound Theme and Name Specifications
// Homepage: https://0pointer.de/lennart/projects/libcanberra/

import (
	"fmt"
	
	"os/exec"
)

func installLibcanberra() {
	// Método 1: Descargar y extraer .tar.gz
	libcanberra_tar_url := "https://0pointer.de/lennart/projects/libcanberra/libcanberra-0.30.tar.xz"
	libcanberra_cmd_tar := exec.Command("curl", "-L", libcanberra_tar_url, "-o", "package.tar.gz")
	err := libcanberra_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libcanberra_zip_url := "https://0pointer.de/lennart/projects/libcanberra/libcanberra-0.30.tar.xz"
	libcanberra_cmd_zip := exec.Command("curl", "-L", libcanberra_zip_url, "-o", "package.zip")
	err = libcanberra_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libcanberra_bin_url := "https://0pointer.de/lennart/projects/libcanberra/libcanberra-0.30.tar.xz"
	libcanberra_cmd_bin := exec.Command("curl", "-L", libcanberra_bin_url, "-o", "binary.bin")
	err = libcanberra_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libcanberra_src_url := "https://0pointer.de/lennart/projects/libcanberra/libcanberra-0.30.tar.xz"
	libcanberra_cmd_src := exec.Command("curl", "-L", libcanberra_src_url, "-o", "source.tar.gz")
	err = libcanberra_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libcanberra_cmd_direct := exec.Command("./binary")
	err = libcanberra_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: gtk-doc")
exec.Command("latte", "install", "gtk-doc")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: libvorbis")
exec.Command("latte", "install", "libvorbis")
}
