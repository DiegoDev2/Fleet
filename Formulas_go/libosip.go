package main

// Libosip - Implementation of the eXosip2 stack
// Homepage: https://www.gnu.org/software/osip/

import (
	"fmt"
	
	"os/exec"
)

func installLibosip() {
	// Método 1: Descargar y extraer .tar.gz
	libosip_tar_url := "https://ftp.gnu.org/gnu/osip/libosip2-5.3.1.tar.gz"
	libosip_cmd_tar := exec.Command("curl", "-L", libosip_tar_url, "-o", "package.tar.gz")
	err := libosip_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libosip_zip_url := "https://ftp.gnu.org/gnu/osip/libosip2-5.3.1.zip"
	libosip_cmd_zip := exec.Command("curl", "-L", libosip_zip_url, "-o", "package.zip")
	err = libosip_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libosip_bin_url := "https://ftp.gnu.org/gnu/osip/libosip2-5.3.1.bin"
	libosip_cmd_bin := exec.Command("curl", "-L", libosip_bin_url, "-o", "binary.bin")
	err = libosip_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libosip_src_url := "https://ftp.gnu.org/gnu/osip/libosip2-5.3.1.src.tar.gz"
	libosip_cmd_src := exec.Command("curl", "-L", libosip_src_url, "-o", "source.tar.gz")
	err = libosip_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libosip_cmd_direct := exec.Command("./binary")
	err = libosip_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
