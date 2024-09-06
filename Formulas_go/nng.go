package main

// Nng - Nanomsg-next-generation -- light-weight brokerless messaging
// Homepage: https://nanomsg.github.io/nng/

import (
	"fmt"
	
	"os/exec"
)

func installNng() {
	// Método 1: Descargar y extraer .tar.gz
	nng_tar_url := "https://github.com/nanomsg/nng/archive/refs/tags/v1.8.0.tar.gz"
	nng_cmd_tar := exec.Command("curl", "-L", nng_tar_url, "-o", "package.tar.gz")
	err := nng_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nng_zip_url := "https://github.com/nanomsg/nng/archive/refs/tags/v1.8.0.zip"
	nng_cmd_zip := exec.Command("curl", "-L", nng_zip_url, "-o", "package.zip")
	err = nng_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nng_bin_url := "https://github.com/nanomsg/nng/archive/refs/tags/v1.8.0.bin"
	nng_cmd_bin := exec.Command("curl", "-L", nng_bin_url, "-o", "binary.bin")
	err = nng_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nng_src_url := "https://github.com/nanomsg/nng/archive/refs/tags/v1.8.0.src.tar.gz"
	nng_cmd_src := exec.Command("curl", "-L", nng_src_url, "-o", "source.tar.gz")
	err = nng_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nng_cmd_direct := exec.Command("./binary")
	err = nng_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asciidoctor")
exec.Command("latte", "install", "asciidoctor")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
}
