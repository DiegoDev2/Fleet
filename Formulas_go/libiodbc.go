package main

// Libiodbc - Database connectivity layer based on ODBC. (alternative to unixodbc)
// Homepage: https://www.iodbc.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibiodbc() {
	// Método 1: Descargar y extraer .tar.gz
	libiodbc_tar_url := "https://github.com/openlink/iODBC/archive/refs/tags/v3.52.16.tar.gz"
	libiodbc_cmd_tar := exec.Command("curl", "-L", libiodbc_tar_url, "-o", "package.tar.gz")
	err := libiodbc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libiodbc_zip_url := "https://github.com/openlink/iODBC/archive/refs/tags/v3.52.16.zip"
	libiodbc_cmd_zip := exec.Command("curl", "-L", libiodbc_zip_url, "-o", "package.zip")
	err = libiodbc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libiodbc_bin_url := "https://github.com/openlink/iODBC/archive/refs/tags/v3.52.16.bin"
	libiodbc_cmd_bin := exec.Command("curl", "-L", libiodbc_bin_url, "-o", "binary.bin")
	err = libiodbc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libiodbc_src_url := "https://github.com/openlink/iODBC/archive/refs/tags/v3.52.16.src.tar.gz"
	libiodbc_cmd_src := exec.Command("curl", "-L", libiodbc_src_url, "-o", "source.tar.gz")
	err = libiodbc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libiodbc_cmd_direct := exec.Command("./binary")
	err = libiodbc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
