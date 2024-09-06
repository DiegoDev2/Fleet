package main

// Uftp - Secure, reliable, efficient multicast file transfer program
// Homepage: https://uftp-multicast.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installUftp() {
	// Método 1: Descargar y extraer .tar.gz
	uftp_tar_url := "https://downloads.sourceforge.net/project/uftp-multicast/source-tar/uftp-5.0.3.tar.gz"
	uftp_cmd_tar := exec.Command("curl", "-L", uftp_tar_url, "-o", "package.tar.gz")
	err := uftp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	uftp_zip_url := "https://downloads.sourceforge.net/project/uftp-multicast/source-tar/uftp-5.0.3.zip"
	uftp_cmd_zip := exec.Command("curl", "-L", uftp_zip_url, "-o", "package.zip")
	err = uftp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	uftp_bin_url := "https://downloads.sourceforge.net/project/uftp-multicast/source-tar/uftp-5.0.3.bin"
	uftp_cmd_bin := exec.Command("curl", "-L", uftp_bin_url, "-o", "binary.bin")
	err = uftp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	uftp_src_url := "https://downloads.sourceforge.net/project/uftp-multicast/source-tar/uftp-5.0.3.src.tar.gz"
	uftp_cmd_src := exec.Command("curl", "-L", uftp_src_url, "-o", "source.tar.gz")
	err = uftp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	uftp_cmd_direct := exec.Command("./binary")
	err = uftp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
