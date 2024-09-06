package main

// Yash - Yet another shell: a POSIX-compliant command-line shell
// Homepage: https://magicant.github.io/yash/

import (
	"fmt"
	
	"os/exec"
)

func installYash() {
	// Método 1: Descargar y extraer .tar.gz
	yash_tar_url := "https://github.com/magicant/yash/releases/download/2.57/yash-2.57.tar.xz"
	yash_cmd_tar := exec.Command("curl", "-L", yash_tar_url, "-o", "package.tar.gz")
	err := yash_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yash_zip_url := "https://github.com/magicant/yash/releases/download/2.57/yash-2.57.tar.xz"
	yash_cmd_zip := exec.Command("curl", "-L", yash_zip_url, "-o", "package.zip")
	err = yash_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yash_bin_url := "https://github.com/magicant/yash/releases/download/2.57/yash-2.57.tar.xz"
	yash_cmd_bin := exec.Command("curl", "-L", yash_bin_url, "-o", "binary.bin")
	err = yash_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yash_src_url := "https://github.com/magicant/yash/releases/download/2.57/yash-2.57.tar.xz"
	yash_cmd_src := exec.Command("curl", "-L", yash_src_url, "-o", "source.tar.gz")
	err = yash_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yash_cmd_direct := exec.Command("./binary")
	err = yash_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asciidoc")
exec.Command("latte", "install", "asciidoc")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
