package main

// Mpop - POP3 client
// Homepage: https://marlam.de/mpop/

import (
	"fmt"
	
	"os/exec"
)

func installMpop() {
	// Método 1: Descargar y extraer .tar.gz
	mpop_tar_url := "https://marlam.de/mpop/releases/mpop-1.4.20.tar.xz"
	mpop_cmd_tar := exec.Command("curl", "-L", mpop_tar_url, "-o", "package.tar.gz")
	err := mpop_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mpop_zip_url := "https://marlam.de/mpop/releases/mpop-1.4.20.tar.xz"
	mpop_cmd_zip := exec.Command("curl", "-L", mpop_zip_url, "-o", "package.zip")
	err = mpop_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mpop_bin_url := "https://marlam.de/mpop/releases/mpop-1.4.20.tar.xz"
	mpop_cmd_bin := exec.Command("curl", "-L", mpop_bin_url, "-o", "binary.bin")
	err = mpop_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mpop_src_url := "https://marlam.de/mpop/releases/mpop-1.4.20.tar.xz"
	mpop_cmd_src := exec.Command("curl", "-L", mpop_src_url, "-o", "source.tar.gz")
	err = mpop_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mpop_cmd_direct := exec.Command("./binary")
	err = mpop_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: libidn2")
exec.Command("latte", "install", "libidn2")
}
