package main

// StartupNotification - Reference implementation of startup notification protocol
// Homepage: https://www.freedesktop.org/wiki/Software/startup-notification/

import (
	"fmt"
	
	"os/exec"
)

func installStartupNotification() {
	// Método 1: Descargar y extraer .tar.gz
	startupnotification_tar_url := "https://www.freedesktop.org/software/startup-notification/releases/startup-notification-0.12.tar.gz"
	startupnotification_cmd_tar := exec.Command("curl", "-L", startupnotification_tar_url, "-o", "package.tar.gz")
	err := startupnotification_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	startupnotification_zip_url := "https://www.freedesktop.org/software/startup-notification/releases/startup-notification-0.12.zip"
	startupnotification_cmd_zip := exec.Command("curl", "-L", startupnotification_zip_url, "-o", "package.zip")
	err = startupnotification_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	startupnotification_bin_url := "https://www.freedesktop.org/software/startup-notification/releases/startup-notification-0.12.bin"
	startupnotification_cmd_bin := exec.Command("curl", "-L", startupnotification_bin_url, "-o", "binary.bin")
	err = startupnotification_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	startupnotification_src_url := "https://www.freedesktop.org/software/startup-notification/releases/startup-notification-0.12.src.tar.gz"
	startupnotification_cmd_src := exec.Command("curl", "-L", startupnotification_src_url, "-o", "source.tar.gz")
	err = startupnotification_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	startupnotification_cmd_direct := exec.Command("./binary")
	err = startupnotification_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxcb")
exec.Command("latte", "install", "libxcb")
	fmt.Println("Instalando dependencia: xcb-util")
exec.Command("latte", "install", "xcb-util")
}
