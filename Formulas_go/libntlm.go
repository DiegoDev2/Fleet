package main

// Libntlm - Implements Microsoft's NTLM authentication
// Homepage: https://gitlab.com/gsasl/libntlm/

import (
	"fmt"
	
	"os/exec"
)

func installLibntlm() {
	// Método 1: Descargar y extraer .tar.gz
	libntlm_tar_url := "https://download.savannah.nongnu.org/releases/libntlm/libntlm-1.8.tar.gz"
	libntlm_cmd_tar := exec.Command("curl", "-L", libntlm_tar_url, "-o", "package.tar.gz")
	err := libntlm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libntlm_zip_url := "https://download.savannah.nongnu.org/releases/libntlm/libntlm-1.8.zip"
	libntlm_cmd_zip := exec.Command("curl", "-L", libntlm_zip_url, "-o", "package.zip")
	err = libntlm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libntlm_bin_url := "https://download.savannah.nongnu.org/releases/libntlm/libntlm-1.8.bin"
	libntlm_cmd_bin := exec.Command("curl", "-L", libntlm_bin_url, "-o", "binary.bin")
	err = libntlm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libntlm_src_url := "https://download.savannah.nongnu.org/releases/libntlm/libntlm-1.8.src.tar.gz"
	libntlm_cmd_src := exec.Command("curl", "-L", libntlm_src_url, "-o", "source.tar.gz")
	err = libntlm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libntlm_cmd_direct := exec.Command("./binary")
	err = libntlm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
