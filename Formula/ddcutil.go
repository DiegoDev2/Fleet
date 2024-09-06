package main

// Ddcutil - Control monitor settings using DDC/CI and USB
// Homepage: https://www.ddcutil.com

import (
	"fmt"
	
	"os/exec"
)

func installDdcutil() {
	// Método 1: Descargar y extraer .tar.gz
	ddcutil_tar_url := "https://www.ddcutil.com/tarballs/ddcutil-2.1.4.tar.gz"
	ddcutil_cmd_tar := exec.Command("curl", "-L", ddcutil_tar_url, "-o", "package.tar.gz")
	err := ddcutil_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ddcutil_zip_url := "https://www.ddcutil.com/tarballs/ddcutil-2.1.4.zip"
	ddcutil_cmd_zip := exec.Command("curl", "-L", ddcutil_zip_url, "-o", "package.zip")
	err = ddcutil_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ddcutil_bin_url := "https://www.ddcutil.com/tarballs/ddcutil-2.1.4.bin"
	ddcutil_cmd_bin := exec.Command("curl", "-L", ddcutil_bin_url, "-o", "binary.bin")
	err = ddcutil_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ddcutil_src_url := "https://www.ddcutil.com/tarballs/ddcutil-2.1.4.src.tar.gz"
	ddcutil_cmd_src := exec.Command("curl", "-L", ddcutil_src_url, "-o", "source.tar.gz")
	err = ddcutil_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ddcutil_cmd_direct := exec.Command("./binary")
	err = ddcutil_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: i2c-tools")
	exec.Command("latte", "install", "i2c-tools").Run()
	fmt.Println("Instalando dependencia: jansson")
	exec.Command("latte", "install", "jansson").Run()
	fmt.Println("Instalando dependencia: kmod")
	exec.Command("latte", "install", "kmod").Run()
	fmt.Println("Instalando dependencia: libdrm")
	exec.Command("latte", "install", "libdrm").Run()
	fmt.Println("Instalando dependencia: libusb")
	exec.Command("latte", "install", "libusb").Run()
	fmt.Println("Instalando dependencia: libxrandr")
	exec.Command("latte", "install", "libxrandr").Run()
	fmt.Println("Instalando dependencia: systemd")
	exec.Command("latte", "install", "systemd").Run()
}
