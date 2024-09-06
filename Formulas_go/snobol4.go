package main

// Snobol4 - String oriented and symbolic programming language
// Homepage: https://www.regressive.org/snobol4/

import (
	"fmt"
	
	"os/exec"
)

func installSnobol4() {
	// Método 1: Descargar y extraer .tar.gz
	snobol4_tar_url := "https://ftp.regressive.org/snobol/snobol4-2.3.2.tar.gz"
	snobol4_cmd_tar := exec.Command("curl", "-L", snobol4_tar_url, "-o", "package.tar.gz")
	err := snobol4_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	snobol4_zip_url := "https://ftp.regressive.org/snobol/snobol4-2.3.2.zip"
	snobol4_cmd_zip := exec.Command("curl", "-L", snobol4_zip_url, "-o", "package.zip")
	err = snobol4_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	snobol4_bin_url := "https://ftp.regressive.org/snobol/snobol4-2.3.2.bin"
	snobol4_cmd_bin := exec.Command("curl", "-L", snobol4_bin_url, "-o", "binary.bin")
	err = snobol4_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	snobol4_src_url := "https://ftp.regressive.org/snobol/snobol4-2.3.2.src.tar.gz"
	snobol4_cmd_src := exec.Command("curl", "-L", snobol4_src_url, "-o", "source.tar.gz")
	err = snobol4_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	snobol4_cmd_direct := exec.Command("./binary")
	err = snobol4_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
