package main

// Zeromq - High-performance, asynchronous messaging library
// Homepage: https://zeromq.org/

import (
	"fmt"
	
	"os/exec"
)

func installZeromq() {
	// Método 1: Descargar y extraer .tar.gz
	zeromq_tar_url := "https://github.com/zeromq/libzmq/releases/download/v4.3.5/zeromq-4.3.5.tar.gz"
	zeromq_cmd_tar := exec.Command("curl", "-L", zeromq_tar_url, "-o", "package.tar.gz")
	err := zeromq_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zeromq_zip_url := "https://github.com/zeromq/libzmq/releases/download/v4.3.5/zeromq-4.3.5.zip"
	zeromq_cmd_zip := exec.Command("curl", "-L", zeromq_zip_url, "-o", "package.zip")
	err = zeromq_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zeromq_bin_url := "https://github.com/zeromq/libzmq/releases/download/v4.3.5/zeromq-4.3.5.bin"
	zeromq_cmd_bin := exec.Command("curl", "-L", zeromq_bin_url, "-o", "binary.bin")
	err = zeromq_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zeromq_src_url := "https://github.com/zeromq/libzmq/releases/download/v4.3.5/zeromq-4.3.5.src.tar.gz"
	zeromq_cmd_src := exec.Command("curl", "-L", zeromq_src_url, "-o", "source.tar.gz")
	err = zeromq_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zeromq_cmd_direct := exec.Command("./binary")
	err = zeromq_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: asciidoc")
exec.Command("latte", "install", "asciidoc")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: xmlto")
exec.Command("latte", "install", "xmlto")
	fmt.Println("Instalando dependencia: libsodium")
exec.Command("latte", "install", "libsodium")
}
