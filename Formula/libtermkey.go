package main

// Libtermkey - Library for processing keyboard entry from the terminal
// Homepage: https://www.leonerd.org.uk/code/libtermkey/

import (
	"fmt"
	
	"os/exec"
)

func installLibtermkey() {
	// Método 1: Descargar y extraer .tar.gz
	libtermkey_tar_url := "https://www.leonerd.org.uk/code/libtermkey/libtermkey-0.22.tar.gz"
	libtermkey_cmd_tar := exec.Command("curl", "-L", libtermkey_tar_url, "-o", "package.tar.gz")
	err := libtermkey_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libtermkey_zip_url := "https://www.leonerd.org.uk/code/libtermkey/libtermkey-0.22.zip"
	libtermkey_cmd_zip := exec.Command("curl", "-L", libtermkey_zip_url, "-o", "package.zip")
	err = libtermkey_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libtermkey_bin_url := "https://www.leonerd.org.uk/code/libtermkey/libtermkey-0.22.bin"
	libtermkey_cmd_bin := exec.Command("curl", "-L", libtermkey_bin_url, "-o", "binary.bin")
	err = libtermkey_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libtermkey_src_url := "https://www.leonerd.org.uk/code/libtermkey/libtermkey-0.22.src.tar.gz"
	libtermkey_cmd_src := exec.Command("curl", "-L", libtermkey_src_url, "-o", "source.tar.gz")
	err = libtermkey_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libtermkey_cmd_direct := exec.Command("./binary")
	err = libtermkey_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: unibilium")
	exec.Command("latte", "install", "unibilium").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
}
