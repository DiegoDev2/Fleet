package main

// Libwps - Library to import files in MS Works format
// Homepage: https://libwps.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installLibwps() {
	// Método 1: Descargar y extraer .tar.gz
	libwps_tar_url := "https://downloads.sourceforge.net/project/libwps/libwps/libwps-0.4.14/libwps-0.4.14.tar.xz"
	libwps_cmd_tar := exec.Command("curl", "-L", libwps_tar_url, "-o", "package.tar.gz")
	err := libwps_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libwps_zip_url := "https://downloads.sourceforge.net/project/libwps/libwps/libwps-0.4.14/libwps-0.4.14.tar.xz"
	libwps_cmd_zip := exec.Command("curl", "-L", libwps_zip_url, "-o", "package.zip")
	err = libwps_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libwps_bin_url := "https://downloads.sourceforge.net/project/libwps/libwps/libwps-0.4.14/libwps-0.4.14.tar.xz"
	libwps_cmd_bin := exec.Command("curl", "-L", libwps_bin_url, "-o", "binary.bin")
	err = libwps_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libwps_src_url := "https://downloads.sourceforge.net/project/libwps/libwps/libwps-0.4.14/libwps-0.4.14.tar.xz"
	libwps_cmd_src := exec.Command("curl", "-L", libwps_src_url, "-o", "source.tar.gz")
	err = libwps_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libwps_cmd_direct := exec.Command("./binary")
	err = libwps_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: librevenge")
	exec.Command("latte", "install", "librevenge").Run()
	fmt.Println("Instalando dependencia: libwpd")
	exec.Command("latte", "install", "libwpd").Run()
}
