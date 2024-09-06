package main

// GnupgAT14 - GNU Pretty Good Privacy (PGP) package
// Homepage: https://www.gnupg.org/

import (
	"fmt"
	
	"os/exec"
)

func installGnupgAT14() {
	// Método 1: Descargar y extraer .tar.gz
	gnupgat14_tar_url := "https://gnupg.org/ftp/gcrypt/gnupg/gnupg-1.4.23.tar.bz2"
	gnupgat14_cmd_tar := exec.Command("curl", "-L", gnupgat14_tar_url, "-o", "package.tar.gz")
	err := gnupgat14_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gnupgat14_zip_url := "https://gnupg.org/ftp/gcrypt/gnupg/gnupg-1.4.23.tar.bz2"
	gnupgat14_cmd_zip := exec.Command("curl", "-L", gnupgat14_zip_url, "-o", "package.zip")
	err = gnupgat14_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gnupgat14_bin_url := "https://gnupg.org/ftp/gcrypt/gnupg/gnupg-1.4.23.tar.bz2"
	gnupgat14_cmd_bin := exec.Command("curl", "-L", gnupgat14_bin_url, "-o", "binary.bin")
	err = gnupgat14_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gnupgat14_src_url := "https://gnupg.org/ftp/gcrypt/gnupg/gnupg-1.4.23.tar.bz2"
	gnupgat14_cmd_src := exec.Command("curl", "-L", gnupgat14_src_url, "-o", "source.tar.gz")
	err = gnupgat14_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gnupgat14_cmd_direct := exec.Command("./binary")
	err = gnupgat14_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
