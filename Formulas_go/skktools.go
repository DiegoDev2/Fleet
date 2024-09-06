package main

// Skktools - SKK dictionary maintenance tools
// Homepage: https://github.com/skk-dev/skktools

import (
	"fmt"
	
	"os/exec"
)

func installSkktools() {
	// Método 1: Descargar y extraer .tar.gz
	skktools_tar_url := "https://deb.debian.org/debian/pool/main/s/skktools/skktools_1.3.4.orig.tar.gz"
	skktools_cmd_tar := exec.Command("curl", "-L", skktools_tar_url, "-o", "package.tar.gz")
	err := skktools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	skktools_zip_url := "https://deb.debian.org/debian/pool/main/s/skktools/skktools_1.3.4.orig.zip"
	skktools_cmd_zip := exec.Command("curl", "-L", skktools_zip_url, "-o", "package.zip")
	err = skktools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	skktools_bin_url := "https://deb.debian.org/debian/pool/main/s/skktools/skktools_1.3.4.orig.bin"
	skktools_cmd_bin := exec.Command("curl", "-L", skktools_bin_url, "-o", "binary.bin")
	err = skktools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	skktools_src_url := "https://deb.debian.org/debian/pool/main/s/skktools/skktools_1.3.4.orig.src.tar.gz"
	skktools_cmd_src := exec.Command("curl", "-L", skktools_src_url, "-o", "source.tar.gz")
	err = skktools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	skktools_cmd_direct := exec.Command("./binary")
	err = skktools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gdbm")
exec.Command("latte", "install", "gdbm")
}
