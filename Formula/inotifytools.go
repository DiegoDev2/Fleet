package main

// InotifyTools - C library and command-line programs providing a simple interface to inotify
// Homepage: https://github.com/inotify-tools/inotify-tools

import (
	"fmt"
	
	"os/exec"
)

func installInotifyTools() {
	// Método 1: Descargar y extraer .tar.gz
	inotifytools_tar_url := "https://github.com/inotify-tools/inotify-tools/archive/refs/tags/4.23.9.0.tar.gz"
	inotifytools_cmd_tar := exec.Command("curl", "-L", inotifytools_tar_url, "-o", "package.tar.gz")
	err := inotifytools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	inotifytools_zip_url := "https://github.com/inotify-tools/inotify-tools/archive/refs/tags/4.23.9.0.zip"
	inotifytools_cmd_zip := exec.Command("curl", "-L", inotifytools_zip_url, "-o", "package.zip")
	err = inotifytools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	inotifytools_bin_url := "https://github.com/inotify-tools/inotify-tools/archive/refs/tags/4.23.9.0.bin"
	inotifytools_cmd_bin := exec.Command("curl", "-L", inotifytools_bin_url, "-o", "binary.bin")
	err = inotifytools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	inotifytools_src_url := "https://github.com/inotify-tools/inotify-tools/archive/refs/tags/4.23.9.0.src.tar.gz"
	inotifytools_cmd_src := exec.Command("curl", "-L", inotifytools_src_url, "-o", "source.tar.gz")
	err = inotifytools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	inotifytools_cmd_direct := exec.Command("./binary")
	err = inotifytools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
}
