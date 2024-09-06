package main

// Libquicktime - Library for reading and writing quicktime files
// Homepage: https://libquicktime.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installLibquicktime() {
	// Método 1: Descargar y extraer .tar.gz
	libquicktime_tar_url := "https://downloads.sourceforge.net/project/libquicktime/libquicktime/1.2.4/libquicktime-1.2.4.tar.gz"
	libquicktime_cmd_tar := exec.Command("curl", "-L", libquicktime_tar_url, "-o", "package.tar.gz")
	err := libquicktime_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libquicktime_zip_url := "https://downloads.sourceforge.net/project/libquicktime/libquicktime/1.2.4/libquicktime-1.2.4.zip"
	libquicktime_cmd_zip := exec.Command("curl", "-L", libquicktime_zip_url, "-o", "package.zip")
	err = libquicktime_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libquicktime_bin_url := "https://downloads.sourceforge.net/project/libquicktime/libquicktime/1.2.4/libquicktime-1.2.4.bin"
	libquicktime_cmd_bin := exec.Command("curl", "-L", libquicktime_bin_url, "-o", "binary.bin")
	err = libquicktime_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libquicktime_src_url := "https://downloads.sourceforge.net/project/libquicktime/libquicktime/1.2.4/libquicktime-1.2.4.src.tar.gz"
	libquicktime_cmd_src := exec.Command("curl", "-L", libquicktime_src_url, "-o", "source.tar.gz")
	err = libquicktime_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libquicktime_cmd_direct := exec.Command("./binary")
	err = libquicktime_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
