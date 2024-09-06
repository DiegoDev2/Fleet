package main

// Zyre - Local Area Clustering for Peer-to-Peer Applications
// Homepage: https://github.com/zeromq/zyre

import (
	"fmt"
	
	"os/exec"
)

func installZyre() {
	// Método 1: Descargar y extraer .tar.gz
	zyre_tar_url := "https://github.com/zeromq/zyre/releases/download/v2.0.1/zyre-2.0.1.tar.gz"
	zyre_cmd_tar := exec.Command("curl", "-L", zyre_tar_url, "-o", "package.tar.gz")
	err := zyre_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zyre_zip_url := "https://github.com/zeromq/zyre/releases/download/v2.0.1/zyre-2.0.1.zip"
	zyre_cmd_zip := exec.Command("curl", "-L", zyre_zip_url, "-o", "package.zip")
	err = zyre_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zyre_bin_url := "https://github.com/zeromq/zyre/releases/download/v2.0.1/zyre-2.0.1.bin"
	zyre_cmd_bin := exec.Command("curl", "-L", zyre_bin_url, "-o", "binary.bin")
	err = zyre_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zyre_src_url := "https://github.com/zeromq/zyre/releases/download/v2.0.1/zyre-2.0.1.src.tar.gz"
	zyre_cmd_src := exec.Command("curl", "-L", zyre_src_url, "-o", "source.tar.gz")
	err = zyre_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zyre_cmd_direct := exec.Command("./binary")
	err = zyre_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: czmq")
exec.Command("latte", "install", "czmq")
	fmt.Println("Instalando dependencia: zeromq")
exec.Command("latte", "install", "zeromq")
}
