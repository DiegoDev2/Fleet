package main

// Recode - Convert character set (charsets)
// Homepage: https://github.com/rrthomas/recode

import (
	"fmt"
	
	"os/exec"
)

func installRecode() {
	// Método 1: Descargar y extraer .tar.gz
	recode_tar_url := "https://github.com/rrthomas/recode/releases/download/v3.7.14/recode-3.7.14.tar.gz"
	recode_cmd_tar := exec.Command("curl", "-L", recode_tar_url, "-o", "package.tar.gz")
	err := recode_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	recode_zip_url := "https://github.com/rrthomas/recode/releases/download/v3.7.14/recode-3.7.14.zip"
	recode_cmd_zip := exec.Command("curl", "-L", recode_zip_url, "-o", "package.zip")
	err = recode_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	recode_bin_url := "https://github.com/rrthomas/recode/releases/download/v3.7.14/recode-3.7.14.bin"
	recode_cmd_bin := exec.Command("curl", "-L", recode_bin_url, "-o", "binary.bin")
	err = recode_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	recode_src_url := "https://github.com/rrthomas/recode/releases/download/v3.7.14/recode-3.7.14.src.tar.gz"
	recode_cmd_src := exec.Command("curl", "-L", recode_src_url, "-o", "source.tar.gz")
	err = recode_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	recode_cmd_direct := exec.Command("./binary")
	err = recode_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
