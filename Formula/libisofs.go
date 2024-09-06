package main

// Libisofs - Library to create an ISO-9660 filesystem with various extensions
// Homepage: https://dev.lovelyhq.com/libburnia/libisofs

import (
	"fmt"
	
	"os/exec"
)

func installLibisofs() {
	// Método 1: Descargar y extraer .tar.gz
	libisofs_tar_url := "https://files.libburnia-project.org/releases/libisofs-1.5.6.pl01.tar.gz"
	libisofs_cmd_tar := exec.Command("curl", "-L", libisofs_tar_url, "-o", "package.tar.gz")
	err := libisofs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libisofs_zip_url := "https://files.libburnia-project.org/releases/libisofs-1.5.6.pl01.zip"
	libisofs_cmd_zip := exec.Command("curl", "-L", libisofs_zip_url, "-o", "package.zip")
	err = libisofs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libisofs_bin_url := "https://files.libburnia-project.org/releases/libisofs-1.5.6.pl01.bin"
	libisofs_cmd_bin := exec.Command("curl", "-L", libisofs_bin_url, "-o", "binary.bin")
	err = libisofs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libisofs_src_url := "https://files.libburnia-project.org/releases/libisofs-1.5.6.pl01.src.tar.gz"
	libisofs_cmd_src := exec.Command("curl", "-L", libisofs_src_url, "-o", "source.tar.gz")
	err = libisofs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libisofs_cmd_direct := exec.Command("./binary")
	err = libisofs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: libzip")
	exec.Command("latte", "install", "libzip").Run()
}
