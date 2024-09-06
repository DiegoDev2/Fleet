package main

// GnuTar - GNU version of the tar archiving utility
// Homepage: https://www.gnu.org/software/tar/

import (
	"fmt"
	
	"os/exec"
)

func installGnuTar() {
	// Método 1: Descargar y extraer .tar.gz
	gnutar_tar_url := "https://ftp.gnu.org/gnu/tar/tar-1.35.tar.gz"
	gnutar_cmd_tar := exec.Command("curl", "-L", gnutar_tar_url, "-o", "package.tar.gz")
	err := gnutar_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gnutar_zip_url := "https://ftp.gnu.org/gnu/tar/tar-1.35.zip"
	gnutar_cmd_zip := exec.Command("curl", "-L", gnutar_zip_url, "-o", "package.zip")
	err = gnutar_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gnutar_bin_url := "https://ftp.gnu.org/gnu/tar/tar-1.35.bin"
	gnutar_cmd_bin := exec.Command("curl", "-L", gnutar_bin_url, "-o", "binary.bin")
	err = gnutar_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gnutar_src_url := "https://ftp.gnu.org/gnu/tar/tar-1.35.src.tar.gz"
	gnutar_cmd_src := exec.Command("curl", "-L", gnutar_src_url, "-o", "source.tar.gz")
	err = gnutar_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gnutar_cmd_direct := exec.Command("./binary")
	err = gnutar_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: acl")
	exec.Command("latte", "install", "acl").Run()
}
