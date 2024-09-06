package main

// Pushpin - Reverse proxy for realtime web services
// Homepage: https://pushpin.org/

import (
	"fmt"
	
	"os/exec"
)

func installPushpin() {
	// Método 1: Descargar y extraer .tar.gz
	pushpin_tar_url := "https://github.com/fastly/pushpin/releases/download/v1.40.1/pushpin-1.40.1.tar.bz2"
	pushpin_cmd_tar := exec.Command("curl", "-L", pushpin_tar_url, "-o", "package.tar.gz")
	err := pushpin_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pushpin_zip_url := "https://github.com/fastly/pushpin/releases/download/v1.40.1/pushpin-1.40.1.tar.bz2"
	pushpin_cmd_zip := exec.Command("curl", "-L", pushpin_zip_url, "-o", "package.zip")
	err = pushpin_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pushpin_bin_url := "https://github.com/fastly/pushpin/releases/download/v1.40.1/pushpin-1.40.1.tar.bz2"
	pushpin_cmd_bin := exec.Command("curl", "-L", pushpin_bin_url, "-o", "binary.bin")
	err = pushpin_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pushpin_src_url := "https://github.com/fastly/pushpin/releases/download/v1.40.1/pushpin-1.40.1.tar.bz2"
	pushpin_cmd_src := exec.Command("curl", "-L", pushpin_src_url, "-o", "source.tar.gz")
	err = pushpin_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pushpin_cmd_direct := exec.Command("./binary")
	err = pushpin_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: qt")
exec.Command("latte", "install", "qt")
	fmt.Println("Instalando dependencia: zeromq")
exec.Command("latte", "install", "zeromq")
	fmt.Println("Instalando dependencia: zurl")
exec.Command("latte", "install", "zurl")
	fmt.Println("Instalando dependencia: mongrel2")
exec.Command("latte", "install", "mongrel2")
}
