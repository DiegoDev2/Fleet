package main

// Tmpmail - Temporary email right from your terminal written in POSIX sh
// Homepage: https://github.com/sdushantha/tmpmail

import (
	"fmt"
	
	"os/exec"
)

func installTmpmail() {
	// Método 1: Descargar y extraer .tar.gz
	tmpmail_tar_url := "https://github.com/sdushantha/tmpmail/archive/refs/tags/v1.2.3.tar.gz"
	tmpmail_cmd_tar := exec.Command("curl", "-L", tmpmail_tar_url, "-o", "package.tar.gz")
	err := tmpmail_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tmpmail_zip_url := "https://github.com/sdushantha/tmpmail/archive/refs/tags/v1.2.3.zip"
	tmpmail_cmd_zip := exec.Command("curl", "-L", tmpmail_zip_url, "-o", "package.zip")
	err = tmpmail_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tmpmail_bin_url := "https://github.com/sdushantha/tmpmail/archive/refs/tags/v1.2.3.bin"
	tmpmail_cmd_bin := exec.Command("curl", "-L", tmpmail_bin_url, "-o", "binary.bin")
	err = tmpmail_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tmpmail_src_url := "https://github.com/sdushantha/tmpmail/archive/refs/tags/v1.2.3.src.tar.gz"
	tmpmail_cmd_src := exec.Command("curl", "-L", tmpmail_src_url, "-o", "source.tar.gz")
	err = tmpmail_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tmpmail_cmd_direct := exec.Command("./binary")
	err = tmpmail_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: jq")
	exec.Command("latte", "install", "jq").Run()
	fmt.Println("Instalando dependencia: w3m")
	exec.Command("latte", "install", "w3m").Run()
	fmt.Println("Instalando dependencia: xclip")
	exec.Command("latte", "install", "xclip").Run()
}
