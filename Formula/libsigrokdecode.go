package main

// Libsigrokdecode - Drivers for logic analyzers and other supported devices
// Homepage: https://sigrok.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibsigrokdecode() {
	// Método 1: Descargar y extraer .tar.gz
	libsigrokdecode_tar_url := "https://sigrok.org/download/source/libsigrokdecode/libsigrokdecode-0.5.3.tar.gz"
	libsigrokdecode_cmd_tar := exec.Command("curl", "-L", libsigrokdecode_tar_url, "-o", "package.tar.gz")
	err := libsigrokdecode_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libsigrokdecode_zip_url := "https://sigrok.org/download/source/libsigrokdecode/libsigrokdecode-0.5.3.zip"
	libsigrokdecode_cmd_zip := exec.Command("curl", "-L", libsigrokdecode_zip_url, "-o", "package.zip")
	err = libsigrokdecode_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libsigrokdecode_bin_url := "https://sigrok.org/download/source/libsigrokdecode/libsigrokdecode-0.5.3.bin"
	libsigrokdecode_cmd_bin := exec.Command("curl", "-L", libsigrokdecode_bin_url, "-o", "binary.bin")
	err = libsigrokdecode_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libsigrokdecode_src_url := "https://sigrok.org/download/source/libsigrokdecode/libsigrokdecode-0.5.3.src.tar.gz"
	libsigrokdecode_cmd_src := exec.Command("curl", "-L", libsigrokdecode_src_url, "-o", "source.tar.gz")
	err = libsigrokdecode_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libsigrokdecode_cmd_direct := exec.Command("./binary")
	err = libsigrokdecode_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: doxygen")
	exec.Command("latte", "install", "doxygen").Run()
	fmt.Println("Instalando dependencia: graphviz")
	exec.Command("latte", "install", "graphviz").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
