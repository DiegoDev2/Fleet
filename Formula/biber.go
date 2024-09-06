package main

// Biber - Backend processor for BibLaTeX
// Homepage: https://sourceforge.net/projects/biblatex-biber/

import (
	"fmt"
	
	"os/exec"
)

func installBiber() {
	// Método 1: Descargar y extraer .tar.gz
	biber_tar_url := "https://github.com/plk/biber/archive/refs/tags/v2.20.tar.gz"
	biber_cmd_tar := exec.Command("curl", "-L", biber_tar_url, "-o", "package.tar.gz")
	err := biber_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	biber_zip_url := "https://github.com/plk/biber/archive/refs/tags/v2.20.zip"
	biber_cmd_zip := exec.Command("curl", "-L", biber_zip_url, "-o", "package.zip")
	err = biber_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	biber_bin_url := "https://github.com/plk/biber/archive/refs/tags/v2.20.bin"
	biber_cmd_bin := exec.Command("curl", "-L", biber_bin_url, "-o", "binary.bin")
	err = biber_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	biber_src_url := "https://github.com/plk/biber/archive/refs/tags/v2.20.src.tar.gz"
	biber_cmd_src := exec.Command("curl", "-L", biber_src_url, "-o", "source.tar.gz")
	err = biber_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	biber_cmd_direct := exec.Command("./binary")
	err = biber_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: texlive")
	exec.Command("latte", "install", "texlive").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: perl")
	exec.Command("latte", "install", "perl").Run()
}
