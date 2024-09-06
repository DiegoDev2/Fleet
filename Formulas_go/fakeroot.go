package main

// Fakeroot - Provide a fake root environment
// Homepage: https://tracker.debian.org/pkg/fakeroot

import (
	"fmt"
	
	"os/exec"
)

func installFakeroot() {
	// Método 1: Descargar y extraer .tar.gz
	fakeroot_tar_url := "https://deb.debian.org/debian/pool/main/f/fakeroot/fakeroot_1.36.orig.tar.gz"
	fakeroot_cmd_tar := exec.Command("curl", "-L", fakeroot_tar_url, "-o", "package.tar.gz")
	err := fakeroot_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fakeroot_zip_url := "https://deb.debian.org/debian/pool/main/f/fakeroot/fakeroot_1.36.orig.zip"
	fakeroot_cmd_zip := exec.Command("curl", "-L", fakeroot_zip_url, "-o", "package.zip")
	err = fakeroot_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fakeroot_bin_url := "https://deb.debian.org/debian/pool/main/f/fakeroot/fakeroot_1.36.orig.bin"
	fakeroot_cmd_bin := exec.Command("curl", "-L", fakeroot_bin_url, "-o", "binary.bin")
	err = fakeroot_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fakeroot_src_url := "https://deb.debian.org/debian/pool/main/f/fakeroot/fakeroot_1.36.orig.src.tar.gz"
	fakeroot_cmd_src := exec.Command("curl", "-L", fakeroot_src_url, "-o", "source.tar.gz")
	err = fakeroot_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fakeroot_cmd_direct := exec.Command("./binary")
	err = fakeroot_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libcap")
exec.Command("latte", "install", "libcap")
}
