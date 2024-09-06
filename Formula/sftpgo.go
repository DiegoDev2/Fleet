package main

// Sftpgo - Fully featured SFTP server with optional HTTP/S, FTP/S and WebDAV support
// Homepage: https://github.com/drakkan/sftpgo

import (
	"fmt"
	
	"os/exec"
)

func installSftpgo() {
	// Método 1: Descargar y extraer .tar.gz
	sftpgo_tar_url := "https://github.com/drakkan/sftpgo/releases/download/v2.6.2/sftpgo_v2.6.2_src_with_deps.tar.xz"
	sftpgo_cmd_tar := exec.Command("curl", "-L", sftpgo_tar_url, "-o", "package.tar.gz")
	err := sftpgo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sftpgo_zip_url := "https://github.com/drakkan/sftpgo/releases/download/v2.6.2/sftpgo_v2.6.2_src_with_deps.tar.xz"
	sftpgo_cmd_zip := exec.Command("curl", "-L", sftpgo_zip_url, "-o", "package.zip")
	err = sftpgo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sftpgo_bin_url := "https://github.com/drakkan/sftpgo/releases/download/v2.6.2/sftpgo_v2.6.2_src_with_deps.tar.xz"
	sftpgo_cmd_bin := exec.Command("curl", "-L", sftpgo_bin_url, "-o", "binary.bin")
	err = sftpgo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sftpgo_src_url := "https://github.com/drakkan/sftpgo/releases/download/v2.6.2/sftpgo_v2.6.2_src_with_deps.tar.xz"
	sftpgo_cmd_src := exec.Command("curl", "-L", sftpgo_src_url, "-o", "source.tar.gz")
	err = sftpgo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sftpgo_cmd_direct := exec.Command("./binary")
	err = sftpgo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
