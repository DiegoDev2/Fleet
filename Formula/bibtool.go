package main

// BibTool - Manipulates BibTeX databases
// Homepage: https://www.gerd-neugebauer.de/software/TeX/BibTool/en/

import (
	"fmt"
	
	"os/exec"
)

func installBibTool() {
	// Método 1: Descargar y extraer .tar.gz
	bibtool_tar_url := "https://github.com/ge-ne/bibtool/releases/download/BibTool_2_68/BibTool-2.68.tar.gz"
	bibtool_cmd_tar := exec.Command("curl", "-L", bibtool_tar_url, "-o", "package.tar.gz")
	err := bibtool_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bibtool_zip_url := "https://github.com/ge-ne/bibtool/releases/download/BibTool_2_68/BibTool-2.68.zip"
	bibtool_cmd_zip := exec.Command("curl", "-L", bibtool_zip_url, "-o", "package.zip")
	err = bibtool_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bibtool_bin_url := "https://github.com/ge-ne/bibtool/releases/download/BibTool_2_68/BibTool-2.68.bin"
	bibtool_cmd_bin := exec.Command("curl", "-L", bibtool_bin_url, "-o", "binary.bin")
	err = bibtool_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bibtool_src_url := "https://github.com/ge-ne/bibtool/releases/download/BibTool_2_68/BibTool-2.68.src.tar.gz"
	bibtool_cmd_src := exec.Command("curl", "-L", bibtool_src_url, "-o", "source.tar.gz")
	err = bibtool_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bibtool_cmd_direct := exec.Command("./binary")
	err = bibtool_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
