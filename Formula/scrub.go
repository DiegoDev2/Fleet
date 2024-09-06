package main

// Scrub - Writes patterns on magnetic media to thwart data recovery
// Homepage: https://github.com/chaos/scrub

import (
	"fmt"
	
	"os/exec"
)

func installScrub() {
	// Método 1: Descargar y extraer .tar.gz
	scrub_tar_url := "https://github.com/chaos/scrub/releases/download/2.6.1/scrub-2.6.1.tar.gz"
	scrub_cmd_tar := exec.Command("curl", "-L", scrub_tar_url, "-o", "package.tar.gz")
	err := scrub_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scrub_zip_url := "https://github.com/chaos/scrub/releases/download/2.6.1/scrub-2.6.1.zip"
	scrub_cmd_zip := exec.Command("curl", "-L", scrub_zip_url, "-o", "package.zip")
	err = scrub_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scrub_bin_url := "https://github.com/chaos/scrub/releases/download/2.6.1/scrub-2.6.1.bin"
	scrub_cmd_bin := exec.Command("curl", "-L", scrub_bin_url, "-o", "binary.bin")
	err = scrub_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scrub_src_url := "https://github.com/chaos/scrub/releases/download/2.6.1/scrub-2.6.1.src.tar.gz"
	scrub_cmd_src := exec.Command("curl", "-L", scrub_src_url, "-o", "source.tar.gz")
	err = scrub_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scrub_cmd_direct := exec.Command("./binary")
	err = scrub_cmd_direct.Run()
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
}
