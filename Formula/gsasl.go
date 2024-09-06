package main

// Gsasl - SASL library command-line interface
// Homepage: https://www.gnu.org/software/gsasl/

import (
	"fmt"
	
	"os/exec"
)

func installGsasl() {
	// Método 1: Descargar y extraer .tar.gz
	gsasl_tar_url := "https://ftp.gnu.org/gnu/gsasl/gsasl-2.2.1.tar.gz"
	gsasl_cmd_tar := exec.Command("curl", "-L", gsasl_tar_url, "-o", "package.tar.gz")
	err := gsasl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gsasl_zip_url := "https://ftp.gnu.org/gnu/gsasl/gsasl-2.2.1.zip"
	gsasl_cmd_zip := exec.Command("curl", "-L", gsasl_zip_url, "-o", "package.zip")
	err = gsasl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gsasl_bin_url := "https://ftp.gnu.org/gnu/gsasl/gsasl-2.2.1.bin"
	gsasl_cmd_bin := exec.Command("curl", "-L", gsasl_bin_url, "-o", "binary.bin")
	err = gsasl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gsasl_src_url := "https://ftp.gnu.org/gnu/gsasl/gsasl-2.2.1.src.tar.gz"
	gsasl_cmd_src := exec.Command("curl", "-L", gsasl_src_url, "-o", "source.tar.gz")
	err = gsasl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gsasl_cmd_direct := exec.Command("./binary")
	err = gsasl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libgcrypt")
	exec.Command("latte", "install", "libgcrypt").Run()
}
