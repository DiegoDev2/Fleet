package main

// Libksba - X.509 and CMS library
// Homepage: https://www.gnupg.org/related_software/libksba/

import (
	"fmt"
	
	"os/exec"
)

func installLibksba() {
	// Método 1: Descargar y extraer .tar.gz
	libksba_tar_url := "https://gnupg.org/ftp/gcrypt/libksba/libksba-1.6.7.tar.bz2"
	libksba_cmd_tar := exec.Command("curl", "-L", libksba_tar_url, "-o", "package.tar.gz")
	err := libksba_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libksba_zip_url := "https://gnupg.org/ftp/gcrypt/libksba/libksba-1.6.7.tar.bz2"
	libksba_cmd_zip := exec.Command("curl", "-L", libksba_zip_url, "-o", "package.zip")
	err = libksba_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libksba_bin_url := "https://gnupg.org/ftp/gcrypt/libksba/libksba-1.6.7.tar.bz2"
	libksba_cmd_bin := exec.Command("curl", "-L", libksba_bin_url, "-o", "binary.bin")
	err = libksba_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libksba_src_url := "https://gnupg.org/ftp/gcrypt/libksba/libksba-1.6.7.tar.bz2"
	libksba_cmd_src := exec.Command("curl", "-L", libksba_src_url, "-o", "source.tar.gz")
	err = libksba_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libksba_cmd_direct := exec.Command("./binary")
	err = libksba_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libgpg-error")
exec.Command("latte", "install", "libgpg-error")
}
