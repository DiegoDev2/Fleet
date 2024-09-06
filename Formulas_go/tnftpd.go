package main

// Tnftpd - NetBSD's FTP server
// Homepage: https://cdn.netbsd.org/pub/NetBSD/misc/tnftp/

import (
	"fmt"
	
	"os/exec"
)

func installTnftpd() {
	// Método 1: Descargar y extraer .tar.gz
	tnftpd_tar_url := "https://cdn.netbsd.org/pub/NetBSD/misc/tnftp/tnftpd-20231001.tar.gz"
	tnftpd_cmd_tar := exec.Command("curl", "-L", tnftpd_tar_url, "-o", "package.tar.gz")
	err := tnftpd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tnftpd_zip_url := "https://cdn.netbsd.org/pub/NetBSD/misc/tnftp/tnftpd-20231001.zip"
	tnftpd_cmd_zip := exec.Command("curl", "-L", tnftpd_zip_url, "-o", "package.zip")
	err = tnftpd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tnftpd_bin_url := "https://cdn.netbsd.org/pub/NetBSD/misc/tnftp/tnftpd-20231001.bin"
	tnftpd_cmd_bin := exec.Command("curl", "-L", tnftpd_bin_url, "-o", "binary.bin")
	err = tnftpd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tnftpd_src_url := "https://cdn.netbsd.org/pub/NetBSD/misc/tnftp/tnftpd-20231001.src.tar.gz"
	tnftpd_cmd_src := exec.Command("curl", "-L", tnftpd_src_url, "-o", "source.tar.gz")
	err = tnftpd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tnftpd_cmd_direct := exec.Command("./binary")
	err = tnftpd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
