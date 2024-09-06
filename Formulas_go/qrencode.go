package main

// Qrencode - QR Code generation
// Homepage: https://fukuchi.org/works/qrencode/index.html.en

import (
	"fmt"
	
	"os/exec"
)

func installQrencode() {
	// Método 1: Descargar y extraer .tar.gz
	qrencode_tar_url := "https://fukuchi.org/works/qrencode/qrencode-4.1.1.tar.gz"
	qrencode_cmd_tar := exec.Command("curl", "-L", qrencode_tar_url, "-o", "package.tar.gz")
	err := qrencode_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qrencode_zip_url := "https://fukuchi.org/works/qrencode/qrencode-4.1.1.zip"
	qrencode_cmd_zip := exec.Command("curl", "-L", qrencode_zip_url, "-o", "package.zip")
	err = qrencode_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qrencode_bin_url := "https://fukuchi.org/works/qrencode/qrencode-4.1.1.bin"
	qrencode_cmd_bin := exec.Command("curl", "-L", qrencode_bin_url, "-o", "binary.bin")
	err = qrencode_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qrencode_src_url := "https://fukuchi.org/works/qrencode/qrencode-4.1.1.src.tar.gz"
	qrencode_cmd_src := exec.Command("curl", "-L", qrencode_src_url, "-o", "source.tar.gz")
	err = qrencode_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qrencode_cmd_direct := exec.Command("./binary")
	err = qrencode_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
}
