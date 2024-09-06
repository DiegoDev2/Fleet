package main

// Uni2ascii - Bi-directional conversion between UTF-8 and various ASCII flavors
// Homepage: https://billposer.org/Software/uni2ascii.html

import (
	"fmt"
	
	"os/exec"
)

func installUni2ascii() {
	// Método 1: Descargar y extraer .tar.gz
	uni2ascii_tar_url := "http://billposer.org/Software/Downloads/uni2ascii-4.20.tar.bz2"
	uni2ascii_cmd_tar := exec.Command("curl", "-L", uni2ascii_tar_url, "-o", "package.tar.gz")
	err := uni2ascii_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	uni2ascii_zip_url := "http://billposer.org/Software/Downloads/uni2ascii-4.20.tar.bz2"
	uni2ascii_cmd_zip := exec.Command("curl", "-L", uni2ascii_zip_url, "-o", "package.zip")
	err = uni2ascii_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	uni2ascii_bin_url := "http://billposer.org/Software/Downloads/uni2ascii-4.20.tar.bz2"
	uni2ascii_cmd_bin := exec.Command("curl", "-L", uni2ascii_bin_url, "-o", "binary.bin")
	err = uni2ascii_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	uni2ascii_src_url := "http://billposer.org/Software/Downloads/uni2ascii-4.20.tar.bz2"
	uni2ascii_cmd_src := exec.Command("curl", "-L", uni2ascii_src_url, "-o", "source.tar.gz")
	err = uni2ascii_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	uni2ascii_cmd_direct := exec.Command("./binary")
	err = uni2ascii_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
