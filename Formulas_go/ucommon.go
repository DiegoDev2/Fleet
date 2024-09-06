package main

// Ucommon - GNU C++ runtime library for threads, sockets, and parsing
// Homepage: https://www.gnu.org/software/commoncpp/

import (
	"fmt"
	
	"os/exec"
)

func installUcommon() {
	// Método 1: Descargar y extraer .tar.gz
	ucommon_tar_url := "https://ftp.debian.org/debian/pool/main/u/ucommon/ucommon_7.0.1.orig.tar.gz"
	ucommon_cmd_tar := exec.Command("curl", "-L", ucommon_tar_url, "-o", "package.tar.gz")
	err := ucommon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ucommon_zip_url := "https://ftp.debian.org/debian/pool/main/u/ucommon/ucommon_7.0.1.orig.zip"
	ucommon_cmd_zip := exec.Command("curl", "-L", ucommon_zip_url, "-o", "package.zip")
	err = ucommon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ucommon_bin_url := "https://ftp.debian.org/debian/pool/main/u/ucommon/ucommon_7.0.1.orig.bin"
	ucommon_cmd_bin := exec.Command("curl", "-L", ucommon_bin_url, "-o", "binary.bin")
	err = ucommon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ucommon_src_url := "https://ftp.debian.org/debian/pool/main/u/ucommon/ucommon_7.0.1.orig.src.tar.gz"
	ucommon_cmd_src := exec.Command("curl", "-L", ucommon_src_url, "-o", "source.tar.gz")
	err = ucommon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ucommon_cmd_direct := exec.Command("./binary")
	err = ucommon_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
