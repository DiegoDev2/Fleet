package main

// Libmaa - Low-level data structures including hash tables, sets, lists
// Homepage: https://dict.org/bin/Dict

import (
	"fmt"
	
	"os/exec"
)

func installLibmaa() {
	// Método 1: Descargar y extraer .tar.gz
	libmaa_tar_url := "https://downloads.sourceforge.net/project/dict/libmaa/libmaa-1.5.1/libmaa-1.5.1.tar.gz"
	libmaa_cmd_tar := exec.Command("curl", "-L", libmaa_tar_url, "-o", "package.tar.gz")
	err := libmaa_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libmaa_zip_url := "https://downloads.sourceforge.net/project/dict/libmaa/libmaa-1.5.1/libmaa-1.5.1.zip"
	libmaa_cmd_zip := exec.Command("curl", "-L", libmaa_zip_url, "-o", "package.zip")
	err = libmaa_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libmaa_bin_url := "https://downloads.sourceforge.net/project/dict/libmaa/libmaa-1.5.1/libmaa-1.5.1.bin"
	libmaa_cmd_bin := exec.Command("curl", "-L", libmaa_bin_url, "-o", "binary.bin")
	err = libmaa_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libmaa_src_url := "https://downloads.sourceforge.net/project/dict/libmaa/libmaa-1.5.1/libmaa-1.5.1.src.tar.gz"
	libmaa_cmd_src := exec.Command("curl", "-L", libmaa_src_url, "-o", "source.tar.gz")
	err = libmaa_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libmaa_cmd_direct := exec.Command("./binary")
	err = libmaa_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bmake")
	exec.Command("latte", "install", "bmake").Run()
	fmt.Println("Instalando dependencia: mk-configure")
	exec.Command("latte", "install", "mk-configure").Run()
}
