package main

// Jed - Powerful editor for programmers
// Homepage: https://www.jedsoft.org/jed/

import (
	"fmt"
	
	"os/exec"
)

func installJed() {
	// Método 1: Descargar y extraer .tar.gz
	jed_tar_url := "https://www.jedsoft.org/releases/jed/jed-0.99-19.tar.gz"
	jed_cmd_tar := exec.Command("curl", "-L", jed_tar_url, "-o", "package.tar.gz")
	err := jed_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jed_zip_url := "https://www.jedsoft.org/releases/jed/jed-0.99-19.zip"
	jed_cmd_zip := exec.Command("curl", "-L", jed_zip_url, "-o", "package.zip")
	err = jed_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jed_bin_url := "https://www.jedsoft.org/releases/jed/jed-0.99-19.bin"
	jed_cmd_bin := exec.Command("curl", "-L", jed_bin_url, "-o", "binary.bin")
	err = jed_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jed_src_url := "https://www.jedsoft.org/releases/jed/jed-0.99-19.src.tar.gz"
	jed_cmd_src := exec.Command("curl", "-L", jed_src_url, "-o", "source.tar.gz")
	err = jed_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jed_cmd_direct := exec.Command("./binary")
	err = jed_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: s-lang")
	exec.Command("latte", "install", "s-lang").Run()
}
