package main

// Thrax - Tools for compiling grammars into finite state transducers
// Homepage: https://www.openfst.org/twiki/bin/view/GRM/Thrax

import (
	"fmt"
	
	"os/exec"
)

func installThrax() {
	// Método 1: Descargar y extraer .tar.gz
	thrax_tar_url := "https://www.openfst.org/twiki/pub/GRM/ThraxDownload/thrax-1.3.9.tar.gz"
	thrax_cmd_tar := exec.Command("curl", "-L", thrax_tar_url, "-o", "package.tar.gz")
	err := thrax_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	thrax_zip_url := "https://www.openfst.org/twiki/pub/GRM/ThraxDownload/thrax-1.3.9.zip"
	thrax_cmd_zip := exec.Command("curl", "-L", thrax_zip_url, "-o", "package.zip")
	err = thrax_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	thrax_bin_url := "https://www.openfst.org/twiki/pub/GRM/ThraxDownload/thrax-1.3.9.bin"
	thrax_cmd_bin := exec.Command("curl", "-L", thrax_bin_url, "-o", "binary.bin")
	err = thrax_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	thrax_src_url := "https://www.openfst.org/twiki/pub/GRM/ThraxDownload/thrax-1.3.9.src.tar.gz"
	thrax_cmd_src := exec.Command("curl", "-L", thrax_src_url, "-o", "source.tar.gz")
	err = thrax_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	thrax_cmd_direct := exec.Command("./binary")
	err = thrax_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: openfst")
	exec.Command("latte", "install", "openfst").Run()
}
