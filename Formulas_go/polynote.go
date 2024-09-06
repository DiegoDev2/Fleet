package main

// Polynote - Polyglot notebook with first-class Scala support
// Homepage: https://polynote.org/

import (
	"fmt"
	
	"os/exec"
)

func installPolynote() {
	// Método 1: Descargar y extraer .tar.gz
	polynote_tar_url := "https://github.com/polynote/polynote/releases/download/0.6.0/polynote-dist.tar.gz"
	polynote_cmd_tar := exec.Command("curl", "-L", polynote_tar_url, "-o", "package.tar.gz")
	err := polynote_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	polynote_zip_url := "https://github.com/polynote/polynote/releases/download/0.6.0/polynote-dist.zip"
	polynote_cmd_zip := exec.Command("curl", "-L", polynote_zip_url, "-o", "package.zip")
	err = polynote_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	polynote_bin_url := "https://github.com/polynote/polynote/releases/download/0.6.0/polynote-dist.bin"
	polynote_cmd_bin := exec.Command("curl", "-L", polynote_bin_url, "-o", "binary.bin")
	err = polynote_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	polynote_src_url := "https://github.com/polynote/polynote/releases/download/0.6.0/polynote-dist.src.tar.gz"
	polynote_cmd_src := exec.Command("curl", "-L", polynote_src_url, "-o", "source.tar.gz")
	err = polynote_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	polynote_cmd_direct := exec.Command("./binary")
	err = polynote_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: numpy")
exec.Command("latte", "install", "numpy")
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
