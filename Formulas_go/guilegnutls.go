package main

// GuileGnutls - Guile bindings for the GnuTLS library
// Homepage: https://gitlab.com/gnutls/guile

import (
	"fmt"
	
	"os/exec"
)

func installGuileGnutls() {
	// Método 1: Descargar y extraer .tar.gz
	guilegnutls_tar_url := "https://gitlab.com/gnutls/guile/uploads/9060bc55069cedb40ab46cea49b439c0/guile-gnutls-4.0.0.tar.gz"
	guilegnutls_cmd_tar := exec.Command("curl", "-L", guilegnutls_tar_url, "-o", "package.tar.gz")
	err := guilegnutls_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	guilegnutls_zip_url := "https://gitlab.com/gnutls/guile/uploads/9060bc55069cedb40ab46cea49b439c0/guile-gnutls-4.0.0.zip"
	guilegnutls_cmd_zip := exec.Command("curl", "-L", guilegnutls_zip_url, "-o", "package.zip")
	err = guilegnutls_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	guilegnutls_bin_url := "https://gitlab.com/gnutls/guile/uploads/9060bc55069cedb40ab46cea49b439c0/guile-gnutls-4.0.0.bin"
	guilegnutls_cmd_bin := exec.Command("curl", "-L", guilegnutls_bin_url, "-o", "binary.bin")
	err = guilegnutls_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	guilegnutls_src_url := "https://gitlab.com/gnutls/guile/uploads/9060bc55069cedb40ab46cea49b439c0/guile-gnutls-4.0.0.src.tar.gz"
	guilegnutls_cmd_src := exec.Command("curl", "-L", guilegnutls_src_url, "-o", "source.tar.gz")
	err = guilegnutls_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	guilegnutls_cmd_direct := exec.Command("./binary")
	err = guilegnutls_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: guile")
exec.Command("latte", "install", "guile")
	fmt.Println("Instalando dependencia: bdw-gc")
exec.Command("latte", "install", "bdw-gc")
}
