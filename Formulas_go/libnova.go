package main

// Libnova - Celestial mechanics, astrometry and astrodynamics library
// Homepage: https://libnova.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installLibnova() {
	// Método 1: Descargar y extraer .tar.gz
	libnova_tar_url := "https://downloads.sourceforge.net/project/libnova/libnova/v%200.15.0/libnova-0.15.0.tar.gz"
	libnova_cmd_tar := exec.Command("curl", "-L", libnova_tar_url, "-o", "package.tar.gz")
	err := libnova_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libnova_zip_url := "https://downloads.sourceforge.net/project/libnova/libnova/v%200.15.0/libnova-0.15.0.zip"
	libnova_cmd_zip := exec.Command("curl", "-L", libnova_zip_url, "-o", "package.zip")
	err = libnova_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libnova_bin_url := "https://downloads.sourceforge.net/project/libnova/libnova/v%200.15.0/libnova-0.15.0.bin"
	libnova_cmd_bin := exec.Command("curl", "-L", libnova_bin_url, "-o", "binary.bin")
	err = libnova_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libnova_src_url := "https://downloads.sourceforge.net/project/libnova/libnova/v%200.15.0/libnova-0.15.0.src.tar.gz"
	libnova_cmd_src := exec.Command("curl", "-L", libnova_src_url, "-o", "source.tar.gz")
	err = libnova_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libnova_cmd_direct := exec.Command("./binary")
	err = libnova_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
