package main

// TtyShare - Terminal sharing over the Internet
// Homepage: https://tty-share.com/

import (
	"fmt"
	
	"os/exec"
)

func installTtyShare() {
	// Método 1: Descargar y extraer .tar.gz
	ttyshare_tar_url := "https://github.com/elisescu/tty-share/archive/refs/tags/v2.4.0.tar.gz"
	ttyshare_cmd_tar := exec.Command("curl", "-L", ttyshare_tar_url, "-o", "package.tar.gz")
	err := ttyshare_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ttyshare_zip_url := "https://github.com/elisescu/tty-share/archive/refs/tags/v2.4.0.zip"
	ttyshare_cmd_zip := exec.Command("curl", "-L", ttyshare_zip_url, "-o", "package.zip")
	err = ttyshare_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ttyshare_bin_url := "https://github.com/elisescu/tty-share/archive/refs/tags/v2.4.0.bin"
	ttyshare_cmd_bin := exec.Command("curl", "-L", ttyshare_bin_url, "-o", "binary.bin")
	err = ttyshare_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ttyshare_src_url := "https://github.com/elisescu/tty-share/archive/refs/tags/v2.4.0.src.tar.gz"
	ttyshare_cmd_src := exec.Command("curl", "-L", ttyshare_src_url, "-o", "source.tar.gz")
	err = ttyshare_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ttyshare_cmd_direct := exec.Command("./binary")
	err = ttyshare_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
