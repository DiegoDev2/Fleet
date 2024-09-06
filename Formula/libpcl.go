package main

// Libpcl - C library and API for coroutines
// Homepage: http://www.xmailserver.org/libpcl.html

import (
	"fmt"
	
	"os/exec"
)

func installLibpcl() {
	// Método 1: Descargar y extraer .tar.gz
	libpcl_tar_url := "http://www.xmailserver.org/pcl-1.12.tar.gz"
	libpcl_cmd_tar := exec.Command("curl", "-L", libpcl_tar_url, "-o", "package.tar.gz")
	err := libpcl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libpcl_zip_url := "http://www.xmailserver.org/pcl-1.12.zip"
	libpcl_cmd_zip := exec.Command("curl", "-L", libpcl_zip_url, "-o", "package.zip")
	err = libpcl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libpcl_bin_url := "http://www.xmailserver.org/pcl-1.12.bin"
	libpcl_cmd_bin := exec.Command("curl", "-L", libpcl_bin_url, "-o", "binary.bin")
	err = libpcl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libpcl_src_url := "http://www.xmailserver.org/pcl-1.12.src.tar.gz"
	libpcl_cmd_src := exec.Command("curl", "-L", libpcl_src_url, "-o", "source.tar.gz")
	err = libpcl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libpcl_cmd_direct := exec.Command("./binary")
	err = libpcl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
