package main

// GumboParser - C99 library for parsing HTML5
// Homepage: https://codeberg.org/gumbo-parser/gumbo-parser

import (
	"fmt"
	
	"os/exec"
)

func installGumboParser() {
	// Método 1: Descargar y extraer .tar.gz
	gumboparser_tar_url := "https://codeberg.org/gumbo-parser/gumbo-parser/archive/0.12.1.tar.gz"
	gumboparser_cmd_tar := exec.Command("curl", "-L", gumboparser_tar_url, "-o", "package.tar.gz")
	err := gumboparser_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gumboparser_zip_url := "https://codeberg.org/gumbo-parser/gumbo-parser/archive/0.12.1.zip"
	gumboparser_cmd_zip := exec.Command("curl", "-L", gumboparser_zip_url, "-o", "package.zip")
	err = gumboparser_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gumboparser_bin_url := "https://codeberg.org/gumbo-parser/gumbo-parser/archive/0.12.1.bin"
	gumboparser_cmd_bin := exec.Command("curl", "-L", gumboparser_bin_url, "-o", "binary.bin")
	err = gumboparser_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gumboparser_src_url := "https://codeberg.org/gumbo-parser/gumbo-parser/archive/0.12.1.src.tar.gz"
	gumboparser_cmd_src := exec.Command("curl", "-L", gumboparser_src_url, "-o", "source.tar.gz")
	err = gumboparser_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gumboparser_cmd_direct := exec.Command("./binary")
	err = gumboparser_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
