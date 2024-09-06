package main

// X11vnc - VNC server for real X displays
// Homepage: https://github.com/LibVNC/x11vnc

import (
	"fmt"
	
	"os/exec"
)

func installX11vnc() {
	// Método 1: Descargar y extraer .tar.gz
	x11vnc_tar_url := "https://github.com/LibVNC/x11vnc/archive/refs/tags/0.9.16.tar.gz"
	x11vnc_cmd_tar := exec.Command("curl", "-L", x11vnc_tar_url, "-o", "package.tar.gz")
	err := x11vnc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	x11vnc_zip_url := "https://github.com/LibVNC/x11vnc/archive/refs/tags/0.9.16.zip"
	x11vnc_cmd_zip := exec.Command("curl", "-L", x11vnc_zip_url, "-o", "package.zip")
	err = x11vnc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	x11vnc_bin_url := "https://github.com/LibVNC/x11vnc/archive/refs/tags/0.9.16.bin"
	x11vnc_cmd_bin := exec.Command("curl", "-L", x11vnc_bin_url, "-o", "binary.bin")
	err = x11vnc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	x11vnc_src_url := "https://github.com/LibVNC/x11vnc/archive/refs/tags/0.9.16.src.tar.gz"
	x11vnc_cmd_src := exec.Command("curl", "-L", x11vnc_src_url, "-o", "source.tar.gz")
	err = x11vnc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	x11vnc_cmd_direct := exec.Command("./binary")
	err = x11vnc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libvncserver")
	exec.Command("latte", "install", "libvncserver").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
