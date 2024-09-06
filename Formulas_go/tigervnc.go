package main

// TigerVnc - High-performance, platform-neutral implementation of VNC
// Homepage: https://tigervnc.org/

import (
	"fmt"
	
	"os/exec"
)

func installTigerVnc() {
	// Método 1: Descargar y extraer .tar.gz
	tigervnc_tar_url := "https://github.com/TigerVNC/tigervnc/archive/refs/tags/v1.14.0.tar.gz"
	tigervnc_cmd_tar := exec.Command("curl", "-L", tigervnc_tar_url, "-o", "package.tar.gz")
	err := tigervnc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tigervnc_zip_url := "https://github.com/TigerVNC/tigervnc/archive/refs/tags/v1.14.0.zip"
	tigervnc_cmd_zip := exec.Command("curl", "-L", tigervnc_zip_url, "-o", "package.zip")
	err = tigervnc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tigervnc_bin_url := "https://github.com/TigerVNC/tigervnc/archive/refs/tags/v1.14.0.bin"
	tigervnc_cmd_bin := exec.Command("curl", "-L", tigervnc_bin_url, "-o", "binary.bin")
	err = tigervnc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tigervnc_src_url := "https://github.com/TigerVNC/tigervnc/archive/refs/tags/v1.14.0.src.tar.gz"
	tigervnc_cmd_src := exec.Command("curl", "-L", tigervnc_src_url, "-o", "source.tar.gz")
	err = tigervnc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tigervnc_cmd_direct := exec.Command("./binary")
	err = tigervnc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: fltk")
exec.Command("latte", "install", "fltk")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: nettle")
exec.Command("latte", "install", "nettle")
	fmt.Println("Instalando dependencia: pixman")
exec.Command("latte", "install", "pixman")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxcursor")
exec.Command("latte", "install", "libxcursor")
	fmt.Println("Instalando dependencia: libxdamage")
exec.Command("latte", "install", "libxdamage")
	fmt.Println("Instalando dependencia: libxext")
exec.Command("latte", "install", "libxext")
	fmt.Println("Instalando dependencia: libxfixes")
exec.Command("latte", "install", "libxfixes")
	fmt.Println("Instalando dependencia: libxft")
exec.Command("latte", "install", "libxft")
	fmt.Println("Instalando dependencia: libxi")
exec.Command("latte", "install", "libxi")
	fmt.Println("Instalando dependencia: libxinerama")
exec.Command("latte", "install", "libxinerama")
	fmt.Println("Instalando dependencia: libxrandr")
exec.Command("latte", "install", "libxrandr")
	fmt.Println("Instalando dependencia: libxrender")
exec.Command("latte", "install", "libxrender")
	fmt.Println("Instalando dependencia: libxtst")
exec.Command("latte", "install", "libxtst")
	fmt.Println("Instalando dependencia: linux-pam")
exec.Command("latte", "install", "linux-pam")
}
