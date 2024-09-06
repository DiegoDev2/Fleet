package main

// Libvncserver - VNC server and client libraries
// Homepage: https://libvnc.github.io

import (
	"fmt"
	
	"os/exec"
)

func installLibvncserver() {
	// Método 1: Descargar y extraer .tar.gz
	libvncserver_tar_url := "https://github.com/LibVNC/libvncserver/archive/refs/tags/LibVNCServer-0.9.14.tar.gz"
	libvncserver_cmd_tar := exec.Command("curl", "-L", libvncserver_tar_url, "-o", "package.tar.gz")
	err := libvncserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libvncserver_zip_url := "https://github.com/LibVNC/libvncserver/archive/refs/tags/LibVNCServer-0.9.14.zip"
	libvncserver_cmd_zip := exec.Command("curl", "-L", libvncserver_zip_url, "-o", "package.zip")
	err = libvncserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libvncserver_bin_url := "https://github.com/LibVNC/libvncserver/archive/refs/tags/LibVNCServer-0.9.14.bin"
	libvncserver_cmd_bin := exec.Command("curl", "-L", libvncserver_bin_url, "-o", "binary.bin")
	err = libvncserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libvncserver_src_url := "https://github.com/LibVNC/libvncserver/archive/refs/tags/LibVNCServer-0.9.14.src.tar.gz"
	libvncserver_cmd_src := exec.Command("curl", "-L", libvncserver_src_url, "-o", "source.tar.gz")
	err = libvncserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libvncserver_cmd_direct := exec.Command("./binary")
	err = libvncserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libgcrypt")
	exec.Command("latte", "install", "libgcrypt").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
