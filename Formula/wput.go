package main

// Wput - Tiny, wget-like FTP client for uploading files
// Homepage: https://wput.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installWput() {
	// Método 1: Descargar y extraer .tar.gz
	wput_tar_url := "https://downloads.sourceforge.net/project/wput/wput/0.6.2/wput-0.6.2.tgz"
	wput_cmd_tar := exec.Command("curl", "-L", wput_tar_url, "-o", "package.tar.gz")
	err := wput_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wput_zip_url := "https://downloads.sourceforge.net/project/wput/wput/0.6.2/wput-0.6.2.tgz"
	wput_cmd_zip := exec.Command("curl", "-L", wput_zip_url, "-o", "package.zip")
	err = wput_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wput_bin_url := "https://downloads.sourceforge.net/project/wput/wput/0.6.2/wput-0.6.2.tgz"
	wput_cmd_bin := exec.Command("curl", "-L", wput_bin_url, "-o", "binary.bin")
	err = wput_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wput_src_url := "https://downloads.sourceforge.net/project/wput/wput/0.6.2/wput-0.6.2.tgz"
	wput_cmd_src := exec.Command("curl", "-L", wput_src_url, "-o", "source.tar.gz")
	err = wput_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wput_cmd_direct := exec.Command("./binary")
	err = wput_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
}
