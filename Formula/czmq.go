package main

// Czmq - High-level C binding for ZeroMQ
// Homepage: http://czmq.zeromq.org/

import (
	"fmt"
	
	"os/exec"
)

func installCzmq() {
	// Método 1: Descargar y extraer .tar.gz
	czmq_tar_url := "https://github.com/zeromq/czmq/releases/download/v4.2.1/czmq-4.2.1.tar.gz"
	czmq_cmd_tar := exec.Command("curl", "-L", czmq_tar_url, "-o", "package.tar.gz")
	err := czmq_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	czmq_zip_url := "https://github.com/zeromq/czmq/releases/download/v4.2.1/czmq-4.2.1.zip"
	czmq_cmd_zip := exec.Command("curl", "-L", czmq_zip_url, "-o", "package.zip")
	err = czmq_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	czmq_bin_url := "https://github.com/zeromq/czmq/releases/download/v4.2.1/czmq-4.2.1.bin"
	czmq_cmd_bin := exec.Command("curl", "-L", czmq_bin_url, "-o", "binary.bin")
	err = czmq_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	czmq_src_url := "https://github.com/zeromq/czmq/releases/download/v4.2.1/czmq-4.2.1.src.tar.gz"
	czmq_cmd_src := exec.Command("curl", "-L", czmq_src_url, "-o", "source.tar.gz")
	err = czmq_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	czmq_cmd_direct := exec.Command("./binary")
	err = czmq_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: asciidoc")
	exec.Command("latte", "install", "asciidoc").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: xmlto")
	exec.Command("latte", "install", "xmlto").Run()
	fmt.Println("Instalando dependencia: lz4")
	exec.Command("latte", "install", "lz4").Run()
	fmt.Println("Instalando dependencia: zeromq")
	exec.Command("latte", "install", "zeromq").Run()
}
