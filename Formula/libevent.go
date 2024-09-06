package main

// Libevent - Asynchronous event library
// Homepage: https://libevent.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibevent() {
	// Método 1: Descargar y extraer .tar.gz
	libevent_tar_url := "https://github.com/libevent/libevent/archive/refs/tags/release-2.1.12-stable.tar.gz"
	libevent_cmd_tar := exec.Command("curl", "-L", libevent_tar_url, "-o", "package.tar.gz")
	err := libevent_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libevent_zip_url := "https://github.com/libevent/libevent/archive/refs/tags/release-2.1.12-stable.zip"
	libevent_cmd_zip := exec.Command("curl", "-L", libevent_zip_url, "-o", "package.zip")
	err = libevent_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libevent_bin_url := "https://github.com/libevent/libevent/archive/refs/tags/release-2.1.12-stable.bin"
	libevent_cmd_bin := exec.Command("curl", "-L", libevent_bin_url, "-o", "binary.bin")
	err = libevent_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libevent_src_url := "https://github.com/libevent/libevent/archive/refs/tags/release-2.1.12-stable.src.tar.gz"
	libevent_cmd_src := exec.Command("curl", "-L", libevent_src_url, "-o", "source.tar.gz")
	err = libevent_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libevent_cmd_direct := exec.Command("./binary")
	err = libevent_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
