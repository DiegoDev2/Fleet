package main

// Nvi - 44BSD re-implementation of vi
// Homepage: https://repo.or.cz/nvi.git

import (
	"fmt"
	
	"os/exec"
)

func installNvi() {
	// Método 1: Descargar y extraer .tar.gz
	nvi_tar_url := "https://deb.debian.org/debian/pool/main/n/nvi/nvi_1.81.6.orig.tar.gz"
	nvi_cmd_tar := exec.Command("curl", "-L", nvi_tar_url, "-o", "package.tar.gz")
	err := nvi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nvi_zip_url := "https://deb.debian.org/debian/pool/main/n/nvi/nvi_1.81.6.orig.zip"
	nvi_cmd_zip := exec.Command("curl", "-L", nvi_zip_url, "-o", "package.zip")
	err = nvi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nvi_bin_url := "https://deb.debian.org/debian/pool/main/n/nvi/nvi_1.81.6.orig.bin"
	nvi_cmd_bin := exec.Command("curl", "-L", nvi_bin_url, "-o", "binary.bin")
	err = nvi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nvi_src_url := "https://deb.debian.org/debian/pool/main/n/nvi/nvi_1.81.6.orig.src.tar.gz"
	nvi_cmd_src := exec.Command("curl", "-L", nvi_src_url, "-o", "source.tar.gz")
	err = nvi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nvi_cmd_direct := exec.Command("./binary")
	err = nvi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
	fmt.Println("Instalando dependencia: berkeley-db@5")
exec.Command("latte", "install", "berkeley-db@5")
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
