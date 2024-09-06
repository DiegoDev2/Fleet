package main

// Bluez - Bluetooth protocol stack for Linux
// Homepage: https://github.com/bluez/bluez

import (
	"fmt"
	
	"os/exec"
)

func installBluez() {
	// Método 1: Descargar y extraer .tar.gz
	bluez_tar_url := "https://mirrors.edge.kernel.org/pub/linux/bluetooth/bluez-5.77.tar.xz"
	bluez_cmd_tar := exec.Command("curl", "-L", bluez_tar_url, "-o", "package.tar.gz")
	err := bluez_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bluez_zip_url := "https://mirrors.edge.kernel.org/pub/linux/bluetooth/bluez-5.77.tar.xz"
	bluez_cmd_zip := exec.Command("curl", "-L", bluez_zip_url, "-o", "package.zip")
	err = bluez_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bluez_bin_url := "https://mirrors.edge.kernel.org/pub/linux/bluetooth/bluez-5.77.tar.xz"
	bluez_cmd_bin := exec.Command("curl", "-L", bluez_bin_url, "-o", "binary.bin")
	err = bluez_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bluez_src_url := "https://mirrors.edge.kernel.org/pub/linux/bluetooth/bluez-5.77.tar.xz"
	bluez_cmd_src := exec.Command("curl", "-L", bluez_src_url, "-o", "source.tar.gz")
	err = bluez_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bluez_cmd_direct := exec.Command("./binary")
	err = bluez_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: dbus")
exec.Command("latte", "install", "dbus")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: libical")
exec.Command("latte", "install", "libical")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
	fmt.Println("Instalando dependencia: systemd")
exec.Command("latte", "install", "systemd")
}
