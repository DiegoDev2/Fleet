package main

// Msieve - C library for factoring large integers
// Homepage: https://sourceforge.net/projects/msieve/

import (
	"fmt"
	
	"os/exec"
)

func installMsieve() {
	// Método 1: Descargar y extraer .tar.gz
	msieve_tar_url := "https://downloads.sourceforge.net/project/msieve/msieve/Msieve%20v1.53/msieve153_src.tar.gz"
	msieve_cmd_tar := exec.Command("curl", "-L", msieve_tar_url, "-o", "package.tar.gz")
	err := msieve_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	msieve_zip_url := "https://downloads.sourceforge.net/project/msieve/msieve/Msieve%20v1.53/msieve153_src.zip"
	msieve_cmd_zip := exec.Command("curl", "-L", msieve_zip_url, "-o", "package.zip")
	err = msieve_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	msieve_bin_url := "https://downloads.sourceforge.net/project/msieve/msieve/Msieve%20v1.53/msieve153_src.bin"
	msieve_cmd_bin := exec.Command("curl", "-L", msieve_bin_url, "-o", "binary.bin")
	err = msieve_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	msieve_src_url := "https://downloads.sourceforge.net/project/msieve/msieve/Msieve%20v1.53/msieve153_src.src.tar.gz"
	msieve_cmd_src := exec.Command("curl", "-L", msieve_src_url, "-o", "source.tar.gz")
	err = msieve_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	msieve_cmd_direct := exec.Command("./binary")
	err = msieve_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
}
