package main

// Tnftp - NetBSD's FTP client
// Homepage: https://cdn.netbsd.org/pub/NetBSD/misc/tnftp/

import (
	"fmt"
	
	"os/exec"
)

func installTnftp() {
	// Método 1: Descargar y extraer .tar.gz
	tnftp_tar_url := "https://cdn.netbsd.org/pub/NetBSD/misc/tnftp/tnftp-20230507.tar.gz"
	tnftp_cmd_tar := exec.Command("curl", "-L", tnftp_tar_url, "-o", "package.tar.gz")
	err := tnftp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tnftp_zip_url := "https://cdn.netbsd.org/pub/NetBSD/misc/tnftp/tnftp-20230507.zip"
	tnftp_cmd_zip := exec.Command("curl", "-L", tnftp_zip_url, "-o", "package.zip")
	err = tnftp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tnftp_bin_url := "https://cdn.netbsd.org/pub/NetBSD/misc/tnftp/tnftp-20230507.bin"
	tnftp_cmd_bin := exec.Command("curl", "-L", tnftp_bin_url, "-o", "binary.bin")
	err = tnftp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tnftp_src_url := "https://cdn.netbsd.org/pub/NetBSD/misc/tnftp/tnftp-20230507.src.tar.gz"
	tnftp_cmd_src := exec.Command("curl", "-L", tnftp_src_url, "-o", "source.tar.gz")
	err = tnftp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tnftp_cmd_direct := exec.Command("./binary")
	err = tnftp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
