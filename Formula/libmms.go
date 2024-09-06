package main

// Libmms - Library for parsing mms:// and mmsh:// network streams
// Homepage: https://sourceforge.net/projects/libmms/

import (
	"fmt"
	
	"os/exec"
)

func installLibmms() {
	// Método 1: Descargar y extraer .tar.gz
	libmms_tar_url := "https://downloads.sourceforge.net/project/libmms/libmms/0.6.4/libmms-0.6.4.tar.gz"
	libmms_cmd_tar := exec.Command("curl", "-L", libmms_tar_url, "-o", "package.tar.gz")
	err := libmms_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libmms_zip_url := "https://downloads.sourceforge.net/project/libmms/libmms/0.6.4/libmms-0.6.4.zip"
	libmms_cmd_zip := exec.Command("curl", "-L", libmms_zip_url, "-o", "package.zip")
	err = libmms_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libmms_bin_url := "https://downloads.sourceforge.net/project/libmms/libmms/0.6.4/libmms-0.6.4.bin"
	libmms_cmd_bin := exec.Command("curl", "-L", libmms_bin_url, "-o", "binary.bin")
	err = libmms_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libmms_src_url := "https://downloads.sourceforge.net/project/libmms/libmms/0.6.4/libmms-0.6.4.src.tar.gz"
	libmms_cmd_src := exec.Command("curl", "-L", libmms_src_url, "-o", "source.tar.gz")
	err = libmms_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libmms_cmd_direct := exec.Command("./binary")
	err = libmms_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
}
