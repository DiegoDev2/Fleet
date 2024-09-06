package main

// Fdupes - Identify or delete duplicate files
// Homepage: https://github.com/adrianlopezroche/fdupes

import (
	"fmt"
	
	"os/exec"
)

func installFdupes() {
	// Método 1: Descargar y extraer .tar.gz
	fdupes_tar_url := "https://github.com/adrianlopezroche/fdupes/releases/download/v2.3.2/fdupes-2.3.2.tar.gz"
	fdupes_cmd_tar := exec.Command("curl", "-L", fdupes_tar_url, "-o", "package.tar.gz")
	err := fdupes_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fdupes_zip_url := "https://github.com/adrianlopezroche/fdupes/releases/download/v2.3.2/fdupes-2.3.2.zip"
	fdupes_cmd_zip := exec.Command("curl", "-L", fdupes_zip_url, "-o", "package.zip")
	err = fdupes_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fdupes_bin_url := "https://github.com/adrianlopezroche/fdupes/releases/download/v2.3.2/fdupes-2.3.2.bin"
	fdupes_cmd_bin := exec.Command("curl", "-L", fdupes_bin_url, "-o", "binary.bin")
	err = fdupes_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fdupes_src_url := "https://github.com/adrianlopezroche/fdupes/releases/download/v2.3.2/fdupes-2.3.2.src.tar.gz"
	fdupes_cmd_src := exec.Command("curl", "-L", fdupes_src_url, "-o", "source.tar.gz")
	err = fdupes_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fdupes_cmd_direct := exec.Command("./binary")
	err = fdupes_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
}
