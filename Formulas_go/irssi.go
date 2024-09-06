package main

// Irssi - Modular IRC client
// Homepage: https://irssi.org/

import (
	"fmt"
	
	"os/exec"
)

func installIrssi() {
	// Método 1: Descargar y extraer .tar.gz
	irssi_tar_url := "https://github.com/irssi/irssi/releases/download/1.4.5/irssi-1.4.5.tar.xz"
	irssi_cmd_tar := exec.Command("curl", "-L", irssi_tar_url, "-o", "package.tar.gz")
	err := irssi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	irssi_zip_url := "https://github.com/irssi/irssi/releases/download/1.4.5/irssi-1.4.5.tar.xz"
	irssi_cmd_zip := exec.Command("curl", "-L", irssi_zip_url, "-o", "package.zip")
	err = irssi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	irssi_bin_url := "https://github.com/irssi/irssi/releases/download/1.4.5/irssi-1.4.5.tar.xz"
	irssi_cmd_bin := exec.Command("curl", "-L", irssi_bin_url, "-o", "binary.bin")
	err = irssi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	irssi_src_url := "https://github.com/irssi/irssi/releases/download/1.4.5/irssi-1.4.5.tar.xz"
	irssi_cmd_src := exec.Command("curl", "-L", irssi_src_url, "-o", "source.tar.gz")
	err = irssi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	irssi_cmd_direct := exec.Command("./binary")
	err = irssi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
