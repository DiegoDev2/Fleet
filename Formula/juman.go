package main

// Juman - Japanese morphological analysis system
// Homepage: https://nlp.ist.i.kyoto-u.ac.jp/index.php?JUMAN

import (
	"fmt"
	
	"os/exec"
)

func installJuman() {
	// Método 1: Descargar y extraer .tar.gz
	juman_tar_url := "https://nlp.ist.i.kyoto-u.ac.jp/nl-resource/juman/juman-7.01.tar.bz2"
	juman_cmd_tar := exec.Command("curl", "-L", juman_tar_url, "-o", "package.tar.gz")
	err := juman_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	juman_zip_url := "https://nlp.ist.i.kyoto-u.ac.jp/nl-resource/juman/juman-7.01.tar.bz2"
	juman_cmd_zip := exec.Command("curl", "-L", juman_zip_url, "-o", "package.zip")
	err = juman_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	juman_bin_url := "https://nlp.ist.i.kyoto-u.ac.jp/nl-resource/juman/juman-7.01.tar.bz2"
	juman_cmd_bin := exec.Command("curl", "-L", juman_bin_url, "-o", "binary.bin")
	err = juman_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	juman_src_url := "https://nlp.ist.i.kyoto-u.ac.jp/nl-resource/juman/juman-7.01.tar.bz2"
	juman_cmd_src := exec.Command("curl", "-L", juman_src_url, "-o", "source.tar.gz")
	err = juman_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	juman_cmd_direct := exec.Command("./binary")
	err = juman_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libnsl")
	exec.Command("latte", "install", "libnsl").Run()
}
