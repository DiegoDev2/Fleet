package main

// Libassuan - Assuan IPC Library
// Homepage: https://www.gnupg.org/related_software/libassuan/

import (
	"fmt"
	
	"os/exec"
)

func installLibassuan() {
	// Método 1: Descargar y extraer .tar.gz
	libassuan_tar_url := "https://gnupg.org/ftp/gcrypt/libassuan/libassuan-3.0.1.tar.bz2"
	libassuan_cmd_tar := exec.Command("curl", "-L", libassuan_tar_url, "-o", "package.tar.gz")
	err := libassuan_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libassuan_zip_url := "https://gnupg.org/ftp/gcrypt/libassuan/libassuan-3.0.1.tar.bz2"
	libassuan_cmd_zip := exec.Command("curl", "-L", libassuan_zip_url, "-o", "package.zip")
	err = libassuan_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libassuan_bin_url := "https://gnupg.org/ftp/gcrypt/libassuan/libassuan-3.0.1.tar.bz2"
	libassuan_cmd_bin := exec.Command("curl", "-L", libassuan_bin_url, "-o", "binary.bin")
	err = libassuan_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libassuan_src_url := "https://gnupg.org/ftp/gcrypt/libassuan/libassuan-3.0.1.tar.bz2"
	libassuan_cmd_src := exec.Command("curl", "-L", libassuan_src_url, "-o", "source.tar.gz")
	err = libassuan_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libassuan_cmd_direct := exec.Command("./binary")
	err = libassuan_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libgpg-error")
exec.Command("latte", "install", "libgpg-error")
}
