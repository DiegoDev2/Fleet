package main

// Libicns - Library for manipulation of the macOS .icns resource format
// Homepage: https://icns.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installLibicns() {
	// Método 1: Descargar y extraer .tar.gz
	libicns_tar_url := "https://downloads.sourceforge.net/project/icns/libicns-0.8.1.tar.gz"
	libicns_cmd_tar := exec.Command("curl", "-L", libicns_tar_url, "-o", "package.tar.gz")
	err := libicns_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libicns_zip_url := "https://downloads.sourceforge.net/project/icns/libicns-0.8.1.zip"
	libicns_cmd_zip := exec.Command("curl", "-L", libicns_zip_url, "-o", "package.zip")
	err = libicns_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libicns_bin_url := "https://downloads.sourceforge.net/project/icns/libicns-0.8.1.bin"
	libicns_cmd_bin := exec.Command("curl", "-L", libicns_bin_url, "-o", "binary.bin")
	err = libicns_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libicns_src_url := "https://downloads.sourceforge.net/project/icns/libicns-0.8.1.src.tar.gz"
	libicns_cmd_src := exec.Command("curl", "-L", libicns_src_url, "-o", "source.tar.gz")
	err = libicns_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libicns_cmd_direct := exec.Command("./binary")
	err = libicns_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: jasper")
	exec.Command("latte", "install", "jasper").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
}
