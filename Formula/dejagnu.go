package main

// DejaGnu - Framework for testing other programs
// Homepage: https://www.gnu.org/software/dejagnu/

import (
	"fmt"
	
	"os/exec"
)

func installDejaGnu() {
	// Método 1: Descargar y extraer .tar.gz
	dejagnu_tar_url := "https://ftp.gnu.org/gnu/dejagnu/dejagnu-1.6.3.tar.gz"
	dejagnu_cmd_tar := exec.Command("curl", "-L", dejagnu_tar_url, "-o", "package.tar.gz")
	err := dejagnu_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dejagnu_zip_url := "https://ftp.gnu.org/gnu/dejagnu/dejagnu-1.6.3.zip"
	dejagnu_cmd_zip := exec.Command("curl", "-L", dejagnu_zip_url, "-o", "package.zip")
	err = dejagnu_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dejagnu_bin_url := "https://ftp.gnu.org/gnu/dejagnu/dejagnu-1.6.3.bin"
	dejagnu_cmd_bin := exec.Command("curl", "-L", dejagnu_bin_url, "-o", "binary.bin")
	err = dejagnu_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dejagnu_src_url := "https://ftp.gnu.org/gnu/dejagnu/dejagnu-1.6.3.src.tar.gz"
	dejagnu_cmd_src := exec.Command("curl", "-L", dejagnu_src_url, "-o", "source.tar.gz")
	err = dejagnu_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dejagnu_cmd_direct := exec.Command("./binary")
	err = dejagnu_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
}
