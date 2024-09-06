package main

// Unixodbc - ODBC 3 connectivity for UNIX
// Homepage: https://www.unixodbc.org/

import (
	"fmt"
	
	"os/exec"
)

func installUnixodbc() {
	// Método 1: Descargar y extraer .tar.gz
	unixodbc_tar_url := "https://www.unixodbc.org/unixODBC-2.3.12.tar.gz"
	unixodbc_cmd_tar := exec.Command("curl", "-L", unixodbc_tar_url, "-o", "package.tar.gz")
	err := unixodbc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	unixodbc_zip_url := "https://www.unixodbc.org/unixODBC-2.3.12.zip"
	unixodbc_cmd_zip := exec.Command("curl", "-L", unixodbc_zip_url, "-o", "package.zip")
	err = unixodbc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	unixodbc_bin_url := "https://www.unixodbc.org/unixODBC-2.3.12.bin"
	unixodbc_cmd_bin := exec.Command("curl", "-L", unixodbc_bin_url, "-o", "binary.bin")
	err = unixodbc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	unixodbc_src_url := "https://www.unixodbc.org/unixODBC-2.3.12.src.tar.gz"
	unixodbc_cmd_src := exec.Command("curl", "-L", unixodbc_src_url, "-o", "source.tar.gz")
	err = unixodbc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	unixodbc_cmd_direct := exec.Command("./binary")
	err = unixodbc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
