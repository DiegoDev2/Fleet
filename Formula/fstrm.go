package main

// Fstrm - Frame Streams implementation in C
// Homepage: https://github.com/farsightsec/fstrm

import (
	"fmt"
	
	"os/exec"
)

func installFstrm() {
	// Método 1: Descargar y extraer .tar.gz
	fstrm_tar_url := "https://dl.farsightsecurity.com/dist/fstrm/fstrm-0.6.1.tar.gz"
	fstrm_cmd_tar := exec.Command("curl", "-L", fstrm_tar_url, "-o", "package.tar.gz")
	err := fstrm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fstrm_zip_url := "https://dl.farsightsecurity.com/dist/fstrm/fstrm-0.6.1.zip"
	fstrm_cmd_zip := exec.Command("curl", "-L", fstrm_zip_url, "-o", "package.zip")
	err = fstrm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fstrm_bin_url := "https://dl.farsightsecurity.com/dist/fstrm/fstrm-0.6.1.bin"
	fstrm_cmd_bin := exec.Command("curl", "-L", fstrm_bin_url, "-o", "binary.bin")
	err = fstrm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fstrm_src_url := "https://dl.farsightsecurity.com/dist/fstrm/fstrm-0.6.1.src.tar.gz"
	fstrm_cmd_src := exec.Command("curl", "-L", fstrm_src_url, "-o", "source.tar.gz")
	err = fstrm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fstrm_cmd_direct := exec.Command("./binary")
	err = fstrm_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
}
