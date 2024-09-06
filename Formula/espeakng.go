package main

// EspeakNg - Speech synthesizer that supports more than hundred languages and accents
// Homepage: https://github.com/espeak-ng/espeak-ng

import (
	"fmt"
	
	"os/exec"
)

func installEspeakNg() {
	// Método 1: Descargar y extraer .tar.gz
	espeakng_tar_url := "https://github.com/espeak-ng/espeak-ng/archive/refs/tags/1.51.tar.gz"
	espeakng_cmd_tar := exec.Command("curl", "-L", espeakng_tar_url, "-o", "package.tar.gz")
	err := espeakng_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	espeakng_zip_url := "https://github.com/espeak-ng/espeak-ng/archive/refs/tags/1.51.zip"
	espeakng_cmd_zip := exec.Command("curl", "-L", espeakng_zip_url, "-o", "package.zip")
	err = espeakng_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	espeakng_bin_url := "https://github.com/espeak-ng/espeak-ng/archive/refs/tags/1.51.bin"
	espeakng_cmd_bin := exec.Command("curl", "-L", espeakng_bin_url, "-o", "binary.bin")
	err = espeakng_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	espeakng_src_url := "https://github.com/espeak-ng/espeak-ng/archive/refs/tags/1.51.src.tar.gz"
	espeakng_cmd_src := exec.Command("curl", "-L", espeakng_src_url, "-o", "source.tar.gz")
	err = espeakng_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	espeakng_cmd_direct := exec.Command("./binary")
	err = espeakng_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pcaudiolib")
	exec.Command("latte", "install", "pcaudiolib").Run()
}
