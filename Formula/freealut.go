package main

// Freealut - Implementation of OpenAL's ALUT standard
// Homepage: https://github.com/vancegroup/freealut

import (
	"fmt"
	
	"os/exec"
)

func installFreealut() {
	// Método 1: Descargar y extraer .tar.gz
	freealut_tar_url := "https://deb.debian.org/debian/pool/main/f/freealut/freealut_1.1.0.orig.tar.gz"
	freealut_cmd_tar := exec.Command("curl", "-L", freealut_tar_url, "-o", "package.tar.gz")
	err := freealut_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	freealut_zip_url := "https://deb.debian.org/debian/pool/main/f/freealut/freealut_1.1.0.orig.zip"
	freealut_cmd_zip := exec.Command("curl", "-L", freealut_zip_url, "-o", "package.zip")
	err = freealut_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	freealut_bin_url := "https://deb.debian.org/debian/pool/main/f/freealut/freealut_1.1.0.orig.bin"
	freealut_cmd_bin := exec.Command("curl", "-L", freealut_bin_url, "-o", "binary.bin")
	err = freealut_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	freealut_src_url := "https://deb.debian.org/debian/pool/main/f/freealut/freealut_1.1.0.orig.src.tar.gz"
	freealut_cmd_src := exec.Command("curl", "-L", freealut_src_url, "-o", "source.tar.gz")
	err = freealut_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	freealut_cmd_direct := exec.Command("./binary")
	err = freealut_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: openal-soft")
	exec.Command("latte", "install", "openal-soft").Run()
}
