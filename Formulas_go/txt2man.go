package main

// Txt2man - Converts flat ASCII text to man page format
// Homepage: https://github.com/mvertes/txt2man/

import (
	"fmt"
	
	"os/exec"
)

func installTxt2man() {
	// Método 1: Descargar y extraer .tar.gz
	txt2man_tar_url := "https://github.com/mvertes/txt2man/archive/refs/tags/txt2man-1.7.1.tar.gz"
	txt2man_cmd_tar := exec.Command("curl", "-L", txt2man_tar_url, "-o", "package.tar.gz")
	err := txt2man_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	txt2man_zip_url := "https://github.com/mvertes/txt2man/archive/refs/tags/txt2man-1.7.1.zip"
	txt2man_cmd_zip := exec.Command("curl", "-L", txt2man_zip_url, "-o", "package.zip")
	err = txt2man_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	txt2man_bin_url := "https://github.com/mvertes/txt2man/archive/refs/tags/txt2man-1.7.1.bin"
	txt2man_cmd_bin := exec.Command("curl", "-L", txt2man_bin_url, "-o", "binary.bin")
	err = txt2man_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	txt2man_src_url := "https://github.com/mvertes/txt2man/archive/refs/tags/txt2man-1.7.1.src.tar.gz"
	txt2man_cmd_src := exec.Command("curl", "-L", txt2man_src_url, "-o", "source.tar.gz")
	err = txt2man_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	txt2man_cmd_direct := exec.Command("./binary")
	err = txt2man_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gawk")
exec.Command("latte", "install", "gawk")
}
