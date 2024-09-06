package main

// Parrot - Open source virtual machine (for Perl6, et al.)
// Homepage: http://www.parrot.org/

import (
	"fmt"
	
	"os/exec"
)

func installParrot() {
	// Método 1: Descargar y extraer .tar.gz
	parrot_tar_url := "http://ftp.parrot.org/releases/supported/8.1.0/parrot-8.1.0.tar.bz2"
	parrot_cmd_tar := exec.Command("curl", "-L", parrot_tar_url, "-o", "package.tar.gz")
	err := parrot_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	parrot_zip_url := "http://ftp.parrot.org/releases/supported/8.1.0/parrot-8.1.0.tar.bz2"
	parrot_cmd_zip := exec.Command("curl", "-L", parrot_zip_url, "-o", "package.zip")
	err = parrot_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	parrot_bin_url := "http://ftp.parrot.org/releases/supported/8.1.0/parrot-8.1.0.tar.bz2"
	parrot_cmd_bin := exec.Command("curl", "-L", parrot_bin_url, "-o", "binary.bin")
	err = parrot_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	parrot_src_url := "http://ftp.parrot.org/releases/supported/8.1.0/parrot-8.1.0.tar.bz2"
	parrot_cmd_src := exec.Command("curl", "-L", parrot_src_url, "-o", "source.tar.gz")
	err = parrot_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	parrot_cmd_direct := exec.Command("./binary")
	err = parrot_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
