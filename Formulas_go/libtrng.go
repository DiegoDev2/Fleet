package main

// Libtrng - Tina's Random Number Generator Library
// Homepage: https://www.numbercrunch.de/trng/

import (
	"fmt"
	
	"os/exec"
)

func installLibtrng() {
	// Método 1: Descargar y extraer .tar.gz
	libtrng_tar_url := "https://github.com/rabauke/trng4/archive/refs/tags/v4.26.tar.gz"
	libtrng_cmd_tar := exec.Command("curl", "-L", libtrng_tar_url, "-o", "package.tar.gz")
	err := libtrng_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libtrng_zip_url := "https://github.com/rabauke/trng4/archive/refs/tags/v4.26.zip"
	libtrng_cmd_zip := exec.Command("curl", "-L", libtrng_zip_url, "-o", "package.zip")
	err = libtrng_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libtrng_bin_url := "https://github.com/rabauke/trng4/archive/refs/tags/v4.26.bin"
	libtrng_cmd_bin := exec.Command("curl", "-L", libtrng_bin_url, "-o", "binary.bin")
	err = libtrng_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libtrng_src_url := "https://github.com/rabauke/trng4/archive/refs/tags/v4.26.src.tar.gz"
	libtrng_cmd_src := exec.Command("curl", "-L", libtrng_src_url, "-o", "source.tar.gz")
	err = libtrng_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libtrng_cmd_direct := exec.Command("./binary")
	err = libtrng_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
