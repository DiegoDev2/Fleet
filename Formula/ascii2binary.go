package main

// Ascii2binary - Converting Text to Binary and Back
// Homepage: https://billposer.org/Software/a2b.html

import (
	"fmt"
	
	"os/exec"
)

func installAscii2binary() {
	// Método 1: Descargar y extraer .tar.gz
	ascii2binary_tar_url := "https://www.billposer.org/Software/Downloads/ascii2binary-2.14.tar.gz"
	ascii2binary_cmd_tar := exec.Command("curl", "-L", ascii2binary_tar_url, "-o", "package.tar.gz")
	err := ascii2binary_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ascii2binary_zip_url := "https://www.billposer.org/Software/Downloads/ascii2binary-2.14.zip"
	ascii2binary_cmd_zip := exec.Command("curl", "-L", ascii2binary_zip_url, "-o", "package.zip")
	err = ascii2binary_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ascii2binary_bin_url := "https://www.billposer.org/Software/Downloads/ascii2binary-2.14.bin"
	ascii2binary_cmd_bin := exec.Command("curl", "-L", ascii2binary_bin_url, "-o", "binary.bin")
	err = ascii2binary_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ascii2binary_src_url := "https://www.billposer.org/Software/Downloads/ascii2binary-2.14.src.tar.gz"
	ascii2binary_cmd_src := exec.Command("curl", "-L", ascii2binary_src_url, "-o", "source.tar.gz")
	err = ascii2binary_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ascii2binary_cmd_direct := exec.Command("./binary")
	err = ascii2binary_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
