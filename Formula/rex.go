package main

// Rex - Command-line tool which executes commands on remote servers
// Homepage: https://www.rexify.org

import (
	"fmt"
	
	"os/exec"
)

func installRex() {
	// Método 1: Descargar y extraer .tar.gz
	rex_tar_url := "https://cpan.metacpan.org/authors/id/F/FE/FERKI/Rex-1.14.3.tar.gz"
	rex_cmd_tar := exec.Command("curl", "-L", rex_tar_url, "-o", "package.tar.gz")
	err := rex_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rex_zip_url := "https://cpan.metacpan.org/authors/id/F/FE/FERKI/Rex-1.14.3.zip"
	rex_cmd_zip := exec.Command("curl", "-L", rex_zip_url, "-o", "package.zip")
	err = rex_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rex_bin_url := "https://cpan.metacpan.org/authors/id/F/FE/FERKI/Rex-1.14.3.bin"
	rex_cmd_bin := exec.Command("curl", "-L", rex_bin_url, "-o", "binary.bin")
	err = rex_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rex_src_url := "https://cpan.metacpan.org/authors/id/F/FE/FERKI/Rex-1.14.3.src.tar.gz"
	rex_cmd_src := exec.Command("curl", "-L", rex_src_url, "-o", "source.tar.gz")
	err = rex_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rex_cmd_direct := exec.Command("./binary")
	err = rex_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
