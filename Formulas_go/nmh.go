package main

// Nmh - New version of the MH mail handler
// Homepage: https://www.nongnu.org/nmh/

import (
	"fmt"
	
	"os/exec"
)

func installNmh() {
	// Método 1: Descargar y extraer .tar.gz
	nmh_tar_url := "https://download.savannah.gnu.org/releases/nmh/nmh-1.8.tar.gz"
	nmh_cmd_tar := exec.Command("curl", "-L", nmh_tar_url, "-o", "package.tar.gz")
	err := nmh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nmh_zip_url := "https://download.savannah.gnu.org/releases/nmh/nmh-1.8.zip"
	nmh_cmd_zip := exec.Command("curl", "-L", nmh_zip_url, "-o", "package.zip")
	err = nmh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nmh_bin_url := "https://download.savannah.gnu.org/releases/nmh/nmh-1.8.bin"
	nmh_cmd_bin := exec.Command("curl", "-L", nmh_bin_url, "-o", "binary.bin")
	err = nmh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nmh_src_url := "https://download.savannah.gnu.org/releases/nmh/nmh-1.8.src.tar.gz"
	nmh_cmd_src := exec.Command("curl", "-L", nmh_src_url, "-o", "source.tar.gz")
	err = nmh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nmh_cmd_direct := exec.Command("./binary")
	err = nmh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: w3m")
exec.Command("latte", "install", "w3m")
	fmt.Println("Instalando dependencia: gdbm")
exec.Command("latte", "install", "gdbm")
}
