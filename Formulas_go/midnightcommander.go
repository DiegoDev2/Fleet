package main

// MidnightCommander - Terminal-based visual file manager
// Homepage: https://www.midnight-commander.org/

import (
	"fmt"
	
	"os/exec"
)

func installMidnightCommander() {
	// Método 1: Descargar y extraer .tar.gz
	midnightcommander_tar_url := "https://www.midnight-commander.org/downloads/mc-4.8.32.tar.xz"
	midnightcommander_cmd_tar := exec.Command("curl", "-L", midnightcommander_tar_url, "-o", "package.tar.gz")
	err := midnightcommander_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	midnightcommander_zip_url := "https://www.midnight-commander.org/downloads/mc-4.8.32.tar.xz"
	midnightcommander_cmd_zip := exec.Command("curl", "-L", midnightcommander_zip_url, "-o", "package.zip")
	err = midnightcommander_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	midnightcommander_bin_url := "https://www.midnight-commander.org/downloads/mc-4.8.32.tar.xz"
	midnightcommander_cmd_bin := exec.Command("curl", "-L", midnightcommander_bin_url, "-o", "binary.bin")
	err = midnightcommander_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	midnightcommander_src_url := "https://www.midnight-commander.org/downloads/mc-4.8.32.tar.xz"
	midnightcommander_cmd_src := exec.Command("curl", "-L", midnightcommander_src_url, "-o", "source.tar.gz")
	err = midnightcommander_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	midnightcommander_cmd_direct := exec.Command("./binary")
	err = midnightcommander_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: libssh2")
exec.Command("latte", "install", "libssh2")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: s-lang")
exec.Command("latte", "install", "s-lang")
	fmt.Println("Instalando dependencia: diffutils")
exec.Command("latte", "install", "diffutils")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
