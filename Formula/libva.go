package main

// Libva - Hardware accelerated video processing library
// Homepage: https://github.com/intel/libva

import (
	"fmt"
	
	"os/exec"
)

func installLibva() {
	// Método 1: Descargar y extraer .tar.gz
	libva_tar_url := "https://github.com/intel/libva/releases/download/2.22.0/libva-2.22.0.tar.bz2"
	libva_cmd_tar := exec.Command("curl", "-L", libva_tar_url, "-o", "package.tar.gz")
	err := libva_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libva_zip_url := "https://github.com/intel/libva/releases/download/2.22.0/libva-2.22.0.tar.bz2"
	libva_cmd_zip := exec.Command("curl", "-L", libva_zip_url, "-o", "package.zip")
	err = libva_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libva_bin_url := "https://github.com/intel/libva/releases/download/2.22.0/libva-2.22.0.tar.bz2"
	libva_cmd_bin := exec.Command("curl", "-L", libva_bin_url, "-o", "binary.bin")
	err = libva_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libva_src_url := "https://github.com/intel/libva/releases/download/2.22.0/libva-2.22.0.tar.bz2"
	libva_cmd_src := exec.Command("curl", "-L", libva_src_url, "-o", "source.tar.gz")
	err = libva_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libva_cmd_direct := exec.Command("./binary")
	err = libva_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libdrm")
	exec.Command("latte", "install", "libdrm").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxext")
	exec.Command("latte", "install", "libxext").Run()
	fmt.Println("Instalando dependencia: libxfixes")
	exec.Command("latte", "install", "libxfixes").Run()
	fmt.Println("Instalando dependencia: wayland")
	exec.Command("latte", "install", "wayland").Run()
}
