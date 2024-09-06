package main

// NumUtils - Programs for dealing with numbers from the command-line
// Homepage: https://suso.suso.org/xulu/Num-utils

import (
	"fmt"
	
	"os/exec"
)

func installNumUtils() {
	// Método 1: Descargar y extraer .tar.gz
	numutils_tar_url := "https://suso.suso.org/programs/num-utils/downloads/num-utils-0.5.tar.gz"
	numutils_cmd_tar := exec.Command("curl", "-L", numutils_tar_url, "-o", "package.tar.gz")
	err := numutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	numutils_zip_url := "https://suso.suso.org/programs/num-utils/downloads/num-utils-0.5.zip"
	numutils_cmd_zip := exec.Command("curl", "-L", numutils_zip_url, "-o", "package.zip")
	err = numutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	numutils_bin_url := "https://suso.suso.org/programs/num-utils/downloads/num-utils-0.5.bin"
	numutils_cmd_bin := exec.Command("curl", "-L", numutils_bin_url, "-o", "binary.bin")
	err = numutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	numutils_src_url := "https://suso.suso.org/programs/num-utils/downloads/num-utils-0.5.src.tar.gz"
	numutils_cmd_src := exec.Command("curl", "-L", numutils_src_url, "-o", "source.tar.gz")
	err = numutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	numutils_cmd_direct := exec.Command("./binary")
	err = numutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pod2man")
exec.Command("latte", "install", "pod2man")
}
