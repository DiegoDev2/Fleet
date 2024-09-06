package main

// Libgig - Library for Gigasampler and DLS (Downloadable Sounds) Level 1/2 files
// Homepage: https://www.linuxsampler.org/libgig/

import (
	"fmt"
	
	"os/exec"
)

func installLibgig() {
	// Método 1: Descargar y extraer .tar.gz
	libgig_tar_url := "https://download.linuxsampler.org/packages/libgig-4.4.1.tar.bz2"
	libgig_cmd_tar := exec.Command("curl", "-L", libgig_tar_url, "-o", "package.tar.gz")
	err := libgig_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libgig_zip_url := "https://download.linuxsampler.org/packages/libgig-4.4.1.tar.bz2"
	libgig_cmd_zip := exec.Command("curl", "-L", libgig_zip_url, "-o", "package.zip")
	err = libgig_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libgig_bin_url := "https://download.linuxsampler.org/packages/libgig-4.4.1.tar.bz2"
	libgig_cmd_bin := exec.Command("curl", "-L", libgig_bin_url, "-o", "binary.bin")
	err = libgig_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libgig_src_url := "https://download.linuxsampler.org/packages/libgig-4.4.1.tar.bz2"
	libgig_cmd_src := exec.Command("curl", "-L", libgig_src_url, "-o", "source.tar.gz")
	err = libgig_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libgig_cmd_direct := exec.Command("./binary")
	err = libgig_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libsndfile")
exec.Command("latte", "install", "libsndfile")
	fmt.Println("Instalando dependencia: e2fsprogs")
exec.Command("latte", "install", "e2fsprogs")
}
