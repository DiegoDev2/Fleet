package main

// Nano - Free (GNU) replacement for the Pico text editor
// Homepage: https://www.nano-editor.org/

import (
	"fmt"
	
	"os/exec"
)

func installNano() {
	// Método 1: Descargar y extraer .tar.gz
	nano_tar_url := "https://www.nano-editor.org/dist/v8/nano-8.2.tar.xz"
	nano_cmd_tar := exec.Command("curl", "-L", nano_tar_url, "-o", "package.tar.gz")
	err := nano_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nano_zip_url := "https://www.nano-editor.org/dist/v8/nano-8.2.tar.xz"
	nano_cmd_zip := exec.Command("curl", "-L", nano_zip_url, "-o", "package.zip")
	err = nano_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nano_bin_url := "https://www.nano-editor.org/dist/v8/nano-8.2.tar.xz"
	nano_cmd_bin := exec.Command("curl", "-L", nano_bin_url, "-o", "binary.bin")
	err = nano_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nano_src_url := "https://www.nano-editor.org/dist/v8/nano-8.2.tar.xz"
	nano_cmd_src := exec.Command("curl", "-L", nano_src_url, "-o", "source.tar.gz")
	err = nano_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nano_cmd_direct := exec.Command("./binary")
	err = nano_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: ncurses")
exec.Command("latte", "install", "ncurses")
	fmt.Println("Instalando dependencia: libmagic")
exec.Command("latte", "install", "libmagic")
}
