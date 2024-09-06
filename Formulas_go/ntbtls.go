package main

// Ntbtls - Not Too Bad TLS Library
// Homepage: https://gnupg.org/

import (
	"fmt"
	
	"os/exec"
)

func installNtbtls() {
	// Método 1: Descargar y extraer .tar.gz
	ntbtls_tar_url := "https://gnupg.org/ftp/gcrypt/ntbtls/ntbtls-0.3.2.tar.bz2"
	ntbtls_cmd_tar := exec.Command("curl", "-L", ntbtls_tar_url, "-o", "package.tar.gz")
	err := ntbtls_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ntbtls_zip_url := "https://gnupg.org/ftp/gcrypt/ntbtls/ntbtls-0.3.2.tar.bz2"
	ntbtls_cmd_zip := exec.Command("curl", "-L", ntbtls_zip_url, "-o", "package.zip")
	err = ntbtls_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ntbtls_bin_url := "https://gnupg.org/ftp/gcrypt/ntbtls/ntbtls-0.3.2.tar.bz2"
	ntbtls_cmd_bin := exec.Command("curl", "-L", ntbtls_bin_url, "-o", "binary.bin")
	err = ntbtls_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ntbtls_src_url := "https://gnupg.org/ftp/gcrypt/ntbtls/ntbtls-0.3.2.tar.bz2"
	ntbtls_cmd_src := exec.Command("curl", "-L", ntbtls_src_url, "-o", "source.tar.gz")
	err = ntbtls_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ntbtls_cmd_direct := exec.Command("./binary")
	err = ntbtls_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libgcrypt")
exec.Command("latte", "install", "libgcrypt")
	fmt.Println("Instalando dependencia: libgpg-error")
exec.Command("latte", "install", "libgpg-error")
	fmt.Println("Instalando dependencia: libksba")
exec.Command("latte", "install", "libksba")
}
