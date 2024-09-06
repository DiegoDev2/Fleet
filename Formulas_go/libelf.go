package main

// Libelf - ELF object file access library
// Homepage: https://fossies.org/linux/misc/old/libelf-0.8.13.tar.gz/

import (
	"fmt"
	
	"os/exec"
)

func installLibelf() {
	// Método 1: Descargar y extraer .tar.gz
	libelf_tar_url := "https://www.mirrorservice.org/sites/ftp.netbsd.org/pub/pkgsrc/distfiles/libelf-0.8.13.tar.gz"
	libelf_cmd_tar := exec.Command("curl", "-L", libelf_tar_url, "-o", "package.tar.gz")
	err := libelf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libelf_zip_url := "https://www.mirrorservice.org/sites/ftp.netbsd.org/pub/pkgsrc/distfiles/libelf-0.8.13.zip"
	libelf_cmd_zip := exec.Command("curl", "-L", libelf_zip_url, "-o", "package.zip")
	err = libelf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libelf_bin_url := "https://www.mirrorservice.org/sites/ftp.netbsd.org/pub/pkgsrc/distfiles/libelf-0.8.13.bin"
	libelf_cmd_bin := exec.Command("curl", "-L", libelf_bin_url, "-o", "binary.bin")
	err = libelf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libelf_src_url := "https://www.mirrorservice.org/sites/ftp.netbsd.org/pub/pkgsrc/distfiles/libelf-0.8.13.src.tar.gz"
	libelf_cmd_src := exec.Command("curl", "-L", libelf_src_url, "-o", "source.tar.gz")
	err = libelf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libelf_cmd_direct := exec.Command("./binary")
	err = libelf_cmd_direct.Run()
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
}
