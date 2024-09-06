package main

// Libcap - User-space interfaces to POSIX 1003.1e capabilities
// Homepage: https://sites.google.com/site/fullycapable/

import (
	"fmt"
	
	"os/exec"
)

func installLibcap() {
	// Método 1: Descargar y extraer .tar.gz
	libcap_tar_url := "https://mirrors.edge.kernel.org/pub/linux/libs/security/linux-privs/libcap2/libcap-2.70.tar.xz"
	libcap_cmd_tar := exec.Command("curl", "-L", libcap_tar_url, "-o", "package.tar.gz")
	err := libcap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libcap_zip_url := "https://mirrors.edge.kernel.org/pub/linux/libs/security/linux-privs/libcap2/libcap-2.70.tar.xz"
	libcap_cmd_zip := exec.Command("curl", "-L", libcap_zip_url, "-o", "package.zip")
	err = libcap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libcap_bin_url := "https://mirrors.edge.kernel.org/pub/linux/libs/security/linux-privs/libcap2/libcap-2.70.tar.xz"
	libcap_cmd_bin := exec.Command("curl", "-L", libcap_bin_url, "-o", "binary.bin")
	err = libcap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libcap_src_url := "https://mirrors.edge.kernel.org/pub/linux/libs/security/linux-privs/libcap2/libcap-2.70.tar.xz"
	libcap_cmd_src := exec.Command("curl", "-L", libcap_src_url, "-o", "source.tar.gz")
	err = libcap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libcap_cmd_direct := exec.Command("./binary")
	err = libcap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
