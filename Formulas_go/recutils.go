package main

// Recutils - Tools to work with human-editable, plain text data files
// Homepage: https://www.gnu.org/software/recutils/

import (
	"fmt"
	
	"os/exec"
)

func installRecutils() {
	// Método 1: Descargar y extraer .tar.gz
	recutils_tar_url := "https://ftp.gnu.org/gnu/recutils/recutils-1.9.tar.gz"
	recutils_cmd_tar := exec.Command("curl", "-L", recutils_tar_url, "-o", "package.tar.gz")
	err := recutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	recutils_zip_url := "https://ftp.gnu.org/gnu/recutils/recutils-1.9.zip"
	recutils_cmd_zip := exec.Command("curl", "-L", recutils_zip_url, "-o", "package.zip")
	err = recutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	recutils_bin_url := "https://ftp.gnu.org/gnu/recutils/recutils-1.9.bin"
	recutils_cmd_bin := exec.Command("curl", "-L", recutils_bin_url, "-o", "binary.bin")
	err = recutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	recutils_src_url := "https://ftp.gnu.org/gnu/recutils/recutils-1.9.src.tar.gz"
	recutils_cmd_src := exec.Command("curl", "-L", recutils_src_url, "-o", "source.tar.gz")
	err = recutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	recutils_cmd_direct := exec.Command("./binary")
	err = recutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libgcrypt")
exec.Command("latte", "install", "libgcrypt")
}
