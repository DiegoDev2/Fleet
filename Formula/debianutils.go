package main

// Debianutils - Miscellaneous utilities specific to Debian
// Homepage: https://packages.debian.org/sid/debianutils

import (
	"fmt"
	
	"os/exec"
)

func installDebianutils() {
	// Método 1: Descargar y extraer .tar.gz
	debianutils_tar_url := "https://deb.debian.org/debian/pool/main/d/debianutils/debianutils_5.17.tar.xz"
	debianutils_cmd_tar := exec.Command("curl", "-L", debianutils_tar_url, "-o", "package.tar.gz")
	err := debianutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	debianutils_zip_url := "https://deb.debian.org/debian/pool/main/d/debianutils/debianutils_5.17.tar.xz"
	debianutils_cmd_zip := exec.Command("curl", "-L", debianutils_zip_url, "-o", "package.zip")
	err = debianutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	debianutils_bin_url := "https://deb.debian.org/debian/pool/main/d/debianutils/debianutils_5.17.tar.xz"
	debianutils_cmd_bin := exec.Command("curl", "-L", debianutils_bin_url, "-o", "binary.bin")
	err = debianutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	debianutils_src_url := "https://deb.debian.org/debian/pool/main/d/debianutils/debianutils_5.17.tar.xz"
	debianutils_cmd_src := exec.Command("curl", "-L", debianutils_src_url, "-o", "source.tar.gz")
	err = debianutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	debianutils_cmd_direct := exec.Command("./binary")
	err = debianutils_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: po4a")
	exec.Command("latte", "install", "po4a").Run()
}
