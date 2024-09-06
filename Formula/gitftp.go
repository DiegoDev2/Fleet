package main

// GitFtp - Git-powered FTP client
// Homepage: https://git-ftp.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installGitFtp() {
	// Método 1: Descargar y extraer .tar.gz
	gitftp_tar_url := "https://github.com/git-ftp/git-ftp/archive/refs/tags/1.6.0.tar.gz"
	gitftp_cmd_tar := exec.Command("curl", "-L", gitftp_tar_url, "-o", "package.tar.gz")
	err := gitftp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitftp_zip_url := "https://github.com/git-ftp/git-ftp/archive/refs/tags/1.6.0.zip"
	gitftp_cmd_zip := exec.Command("curl", "-L", gitftp_zip_url, "-o", "package.zip")
	err = gitftp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitftp_bin_url := "https://github.com/git-ftp/git-ftp/archive/refs/tags/1.6.0.bin"
	gitftp_cmd_bin := exec.Command("curl", "-L", gitftp_bin_url, "-o", "binary.bin")
	err = gitftp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitftp_src_url := "https://github.com/git-ftp/git-ftp/archive/refs/tags/1.6.0.src.tar.gz"
	gitftp_cmd_src := exec.Command("curl", "-L", gitftp_src_url, "-o", "source.tar.gz")
	err = gitftp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitftp_cmd_direct := exec.Command("./binary")
	err = gitftp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pandoc")
	exec.Command("latte", "install", "pandoc").Run()
	fmt.Println("Instalando dependencia: curl")
	exec.Command("latte", "install", "curl").Run()
	fmt.Println("Instalando dependencia: libssh2")
	exec.Command("latte", "install", "libssh2").Run()
}
