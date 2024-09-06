package main

// Dateutils - Tools to manipulate dates with a focus on financial data
// Homepage: https://www.fresse.org/dateutils/

import (
	"fmt"
	
	"os/exec"
)

func installDateutils() {
	// Método 1: Descargar y extraer .tar.gz
	dateutils_tar_url := "https://github.com/hroptatyr/dateutils/releases/download/v0.4.11/dateutils-0.4.11.tar.xz"
	dateutils_cmd_tar := exec.Command("curl", "-L", dateutils_tar_url, "-o", "package.tar.gz")
	err := dateutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dateutils_zip_url := "https://github.com/hroptatyr/dateutils/releases/download/v0.4.11/dateutils-0.4.11.tar.xz"
	dateutils_cmd_zip := exec.Command("curl", "-L", dateutils_zip_url, "-o", "package.zip")
	err = dateutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dateutils_bin_url := "https://github.com/hroptatyr/dateutils/releases/download/v0.4.11/dateutils-0.4.11.tar.xz"
	dateutils_cmd_bin := exec.Command("curl", "-L", dateutils_bin_url, "-o", "binary.bin")
	err = dateutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dateutils_src_url := "https://github.com/hroptatyr/dateutils/releases/download/v0.4.11/dateutils-0.4.11.tar.xz"
	dateutils_cmd_src := exec.Command("curl", "-L", dateutils_src_url, "-o", "source.tar.gz")
	err = dateutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dateutils_cmd_direct := exec.Command("./binary")
	err = dateutils_cmd_direct.Run()
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
