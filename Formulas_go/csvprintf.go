package main

// Csvprintf - Command-line utility for parsing CSV files
// Homepage: https://github.com/archiecobbs/csvprintf

import (
	"fmt"
	
	"os/exec"
)

func installCsvprintf() {
	// Método 1: Descargar y extraer .tar.gz
	csvprintf_tar_url := "https://github.com/archiecobbs/csvprintf/archive/refs/tags/1.3.2.tar.gz"
	csvprintf_cmd_tar := exec.Command("curl", "-L", csvprintf_tar_url, "-o", "package.tar.gz")
	err := csvprintf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	csvprintf_zip_url := "https://github.com/archiecobbs/csvprintf/archive/refs/tags/1.3.2.zip"
	csvprintf_cmd_zip := exec.Command("curl", "-L", csvprintf_zip_url, "-o", "package.zip")
	err = csvprintf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	csvprintf_bin_url := "https://github.com/archiecobbs/csvprintf/archive/refs/tags/1.3.2.bin"
	csvprintf_cmd_bin := exec.Command("curl", "-L", csvprintf_bin_url, "-o", "binary.bin")
	err = csvprintf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	csvprintf_src_url := "https://github.com/archiecobbs/csvprintf/archive/refs/tags/1.3.2.src.tar.gz"
	csvprintf_cmd_src := exec.Command("curl", "-L", csvprintf_src_url, "-o", "source.tar.gz")
	err = csvprintf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	csvprintf_cmd_direct := exec.Command("./binary")
	err = csvprintf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
}
