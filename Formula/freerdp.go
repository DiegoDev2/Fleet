package main

// Freerdp - X11 implementation of the Remote Desktop Protocol (RDP)
// Homepage: https://www.freerdp.com/

import (
	"fmt"
	
	"os/exec"
)

func installFreerdp() {
	// Método 1: Descargar y extraer .tar.gz
	freerdp_tar_url := "https://github.com/FreeRDP/FreeRDP/archive/refs/tags/3.8.0.tar.gz"
	freerdp_cmd_tar := exec.Command("curl", "-L", freerdp_tar_url, "-o", "package.tar.gz")
	err := freerdp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	freerdp_zip_url := "https://github.com/FreeRDP/FreeRDP/archive/refs/tags/3.8.0.zip"
	freerdp_cmd_zip := exec.Command("curl", "-L", freerdp_zip_url, "-o", "package.zip")
	err = freerdp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	freerdp_bin_url := "https://github.com/FreeRDP/FreeRDP/archive/refs/tags/3.8.0.bin"
	freerdp_cmd_bin := exec.Command("curl", "-L", freerdp_bin_url, "-o", "binary.bin")
	err = freerdp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	freerdp_src_url := "https://github.com/FreeRDP/FreeRDP/archive/refs/tags/3.8.0.src.tar.gz"
	freerdp_cmd_src := exec.Command("curl", "-L", freerdp_src_url, "-o", "source.tar.gz")
	err = freerdp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	freerdp_cmd_direct := exec.Command("./binary")
	err = freerdp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: cjson")
	exec.Command("latte", "install", "cjson").Run()
	fmt.Println("Instalando dependencia: ffmpeg")
	exec.Command("latte", "install", "ffmpeg").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libusb")
	exec.Command("latte", "install", "libusb").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxcursor")
	exec.Command("latte", "install", "libxcursor").Run()
	fmt.Println("Instalando dependencia: libxext")
	exec.Command("latte", "install", "libxext").Run()
	fmt.Println("Instalando dependencia: libxfixes")
	exec.Command("latte", "install", "libxfixes").Run()
	fmt.Println("Instalando dependencia: libxi")
	exec.Command("latte", "install", "libxi").Run()
	fmt.Println("Instalando dependencia: libxinerama")
	exec.Command("latte", "install", "libxinerama").Run()
	fmt.Println("Instalando dependencia: libxrandr")
	exec.Command("latte", "install", "libxrandr").Run()
	fmt.Println("Instalando dependencia: libxrender")
	exec.Command("latte", "install", "libxrender").Run()
	fmt.Println("Instalando dependencia: libxv")
	exec.Command("latte", "install", "libxv").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pkcs11-helper")
	exec.Command("latte", "install", "pkcs11-helper").Run()
	fmt.Println("Instalando dependencia: alsa-lib")
	exec.Command("latte", "install", "alsa-lib").Run()
	fmt.Println("Instalando dependencia: ffmpeg")
	exec.Command("latte", "install", "ffmpeg").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: icu4c")
	exec.Command("latte", "install", "icu4c").Run()
	fmt.Println("Instalando dependencia: krb5")
	exec.Command("latte", "install", "krb5").Run()
	fmt.Println("Instalando dependencia: libfuse")
	exec.Command("latte", "install", "libfuse").Run()
	fmt.Println("Instalando dependencia: systemd")
	exec.Command("latte", "install", "systemd").Run()
	fmt.Println("Instalando dependencia: wayland")
	exec.Command("latte", "install", "wayland").Run()
}
