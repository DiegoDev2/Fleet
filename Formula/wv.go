package main

// Wv - Programs for accessing Microsoft Word documents
// Homepage: https://wvware.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installWv() {
	// Método 1: Descargar y extraer .tar.gz
	wv_tar_url := "https://deb.debian.org/debian/pool/main/w/wv/wv_1.2.9.orig.tar.gz"
	wv_cmd_tar := exec.Command("curl", "-L", wv_tar_url, "-o", "package.tar.gz")
	err := wv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wv_zip_url := "https://deb.debian.org/debian/pool/main/w/wv/wv_1.2.9.orig.zip"
	wv_cmd_zip := exec.Command("curl", "-L", wv_zip_url, "-o", "package.zip")
	err = wv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wv_bin_url := "https://deb.debian.org/debian/pool/main/w/wv/wv_1.2.9.orig.bin"
	wv_cmd_bin := exec.Command("curl", "-L", wv_bin_url, "-o", "binary.bin")
	err = wv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wv_src_url := "https://deb.debian.org/debian/pool/main/w/wv/wv_1.2.9.orig.src.tar.gz"
	wv_cmd_src := exec.Command("curl", "-L", wv_src_url, "-o", "source.tar.gz")
	err = wv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wv_cmd_direct := exec.Command("./binary")
	err = wv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: libgsf")
	exec.Command("latte", "install", "libgsf").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libwmf")
	exec.Command("latte", "install", "libwmf").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
