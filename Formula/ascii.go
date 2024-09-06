package main

// Ascii - List ASCII idiomatic names and octal/decimal code-point forms
// Homepage: http://www.catb.org/~esr/ascii/

import (
	"fmt"
	
	"os/exec"
)

func installAscii() {
	// Método 1: Descargar y extraer .tar.gz
	ascii_tar_url := "http://www.catb.org/~esr/ascii/ascii-3.30.tar.gz"
	ascii_cmd_tar := exec.Command("curl", "-L", ascii_tar_url, "-o", "package.tar.gz")
	err := ascii_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ascii_zip_url := "http://www.catb.org/~esr/ascii/ascii-3.30.zip"
	ascii_cmd_zip := exec.Command("curl", "-L", ascii_zip_url, "-o", "package.zip")
	err = ascii_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ascii_bin_url := "http://www.catb.org/~esr/ascii/ascii-3.30.bin"
	ascii_cmd_bin := exec.Command("curl", "-L", ascii_bin_url, "-o", "binary.bin")
	err = ascii_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ascii_src_url := "http://www.catb.org/~esr/ascii/ascii-3.30.src.tar.gz"
	ascii_cmd_src := exec.Command("curl", "-L", ascii_src_url, "-o", "source.tar.gz")
	err = ascii_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ascii_cmd_direct := exec.Command("./binary")
	err = ascii_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: xmlto")
	exec.Command("latte", "install", "xmlto").Run()
}
