package main

// BbftpClient - Secure file transfer software, optimized for large files
// Homepage: http://software.in2p3.fr/bbftp/

import (
	"fmt"
	
	"os/exec"
)

func installBbftpClient() {
	// Método 1: Descargar y extraer .tar.gz
	bbftpclient_tar_url := "http://software.in2p3.fr/bbftp/dist/bbftp-client-3.2.1.tar.gz"
	bbftpclient_cmd_tar := exec.Command("curl", "-L", bbftpclient_tar_url, "-o", "package.tar.gz")
	err := bbftpclient_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bbftpclient_zip_url := "http://software.in2p3.fr/bbftp/dist/bbftp-client-3.2.1.zip"
	bbftpclient_cmd_zip := exec.Command("curl", "-L", bbftpclient_zip_url, "-o", "package.zip")
	err = bbftpclient_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bbftpclient_bin_url := "http://software.in2p3.fr/bbftp/dist/bbftp-client-3.2.1.bin"
	bbftpclient_cmd_bin := exec.Command("curl", "-L", bbftpclient_bin_url, "-o", "binary.bin")
	err = bbftpclient_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bbftpclient_src_url := "http://software.in2p3.fr/bbftp/dist/bbftp-client-3.2.1.src.tar.gz"
	bbftpclient_cmd_src := exec.Command("curl", "-L", bbftpclient_src_url, "-o", "source.tar.gz")
	err = bbftpclient_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bbftpclient_cmd_direct := exec.Command("./binary")
	err = bbftpclient_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
