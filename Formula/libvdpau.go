package main

// Libvdpau - Open source Video Decode and Presentation API library
// Homepage: https://www.freedesktop.org/wiki/Software/VDPAU/

import (
	"fmt"
	
	"os/exec"
)

func installLibvdpau() {
	// Método 1: Descargar y extraer .tar.gz
	libvdpau_tar_url := "https://gitlab.freedesktop.org/vdpau/libvdpau/-/archive/1.5/libvdpau-1.5.tar.bz2"
	libvdpau_cmd_tar := exec.Command("curl", "-L", libvdpau_tar_url, "-o", "package.tar.gz")
	err := libvdpau_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libvdpau_zip_url := "https://gitlab.freedesktop.org/vdpau/libvdpau/-/archive/1.5/libvdpau-1.5.tar.bz2"
	libvdpau_cmd_zip := exec.Command("curl", "-L", libvdpau_zip_url, "-o", "package.zip")
	err = libvdpau_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libvdpau_bin_url := "https://gitlab.freedesktop.org/vdpau/libvdpau/-/archive/1.5/libvdpau-1.5.tar.bz2"
	libvdpau_cmd_bin := exec.Command("curl", "-L", libvdpau_bin_url, "-o", "binary.bin")
	err = libvdpau_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libvdpau_src_url := "https://gitlab.freedesktop.org/vdpau/libvdpau/-/archive/1.5/libvdpau-1.5.tar.bz2"
	libvdpau_cmd_src := exec.Command("curl", "-L", libvdpau_src_url, "-o", "source.tar.gz")
	err = libvdpau_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libvdpau_cmd_direct := exec.Command("./binary")
	err = libvdpau_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxext")
	exec.Command("latte", "install", "libxext").Run()
	fmt.Println("Instalando dependencia: xorgproto")
	exec.Command("latte", "install", "xorgproto").Run()
}
