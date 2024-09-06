package main

// PcscLite - Middleware to access a smart card using SCard API
// Homepage: https://pcsclite.apdu.fr/

import (
	"fmt"
	
	"os/exec"
)

func installPcscLite() {
	// Método 1: Descargar y extraer .tar.gz
	pcsclite_tar_url := "https://pcsclite.apdu.fr/files/pcsc-lite-2.3.0.tar.xz"
	pcsclite_cmd_tar := exec.Command("curl", "-L", pcsclite_tar_url, "-o", "package.tar.gz")
	err := pcsclite_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pcsclite_zip_url := "https://pcsclite.apdu.fr/files/pcsc-lite-2.3.0.tar.xz"
	pcsclite_cmd_zip := exec.Command("curl", "-L", pcsclite_zip_url, "-o", "package.zip")
	err = pcsclite_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pcsclite_bin_url := "https://pcsclite.apdu.fr/files/pcsc-lite-2.3.0.tar.xz"
	pcsclite_cmd_bin := exec.Command("curl", "-L", pcsclite_bin_url, "-o", "binary.bin")
	err = pcsclite_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pcsclite_src_url := "https://pcsclite.apdu.fr/files/pcsc-lite-2.3.0.tar.xz"
	pcsclite_cmd_src := exec.Command("curl", "-L", pcsclite_src_url, "-o", "source.tar.gz")
	err = pcsclite_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pcsclite_cmd_direct := exec.Command("./binary")
	err = pcsclite_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libusb")
	exec.Command("latte", "install", "libusb").Run()
	fmt.Println("Instalando dependencia: systemd")
	exec.Command("latte", "install", "systemd").Run()
}
