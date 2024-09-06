package main

// Wgetpaste - Automate pasting to a number of pastebin services
// Homepage: https://wgetpaste.zlin.dk/

import (
	"fmt"
	
	"os/exec"
)

func installWgetpaste() {
	// Método 1: Descargar y extraer .tar.gz
	wgetpaste_tar_url := "https://github.com/zlin/wgetpaste/releases/download/2.34/wgetpaste-2.34.tar.xz"
	wgetpaste_cmd_tar := exec.Command("curl", "-L", wgetpaste_tar_url, "-o", "package.tar.gz")
	err := wgetpaste_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wgetpaste_zip_url := "https://github.com/zlin/wgetpaste/releases/download/2.34/wgetpaste-2.34.tar.xz"
	wgetpaste_cmd_zip := exec.Command("curl", "-L", wgetpaste_zip_url, "-o", "package.zip")
	err = wgetpaste_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wgetpaste_bin_url := "https://github.com/zlin/wgetpaste/releases/download/2.34/wgetpaste-2.34.tar.xz"
	wgetpaste_cmd_bin := exec.Command("curl", "-L", wgetpaste_bin_url, "-o", "binary.bin")
	err = wgetpaste_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wgetpaste_src_url := "https://github.com/zlin/wgetpaste/releases/download/2.34/wgetpaste-2.34.tar.xz"
	wgetpaste_cmd_src := exec.Command("curl", "-L", wgetpaste_src_url, "-o", "source.tar.gz")
	err = wgetpaste_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wgetpaste_cmd_direct := exec.Command("./binary")
	err = wgetpaste_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: wget")
exec.Command("latte", "install", "wget")
}
