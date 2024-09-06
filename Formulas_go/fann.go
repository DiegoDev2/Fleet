package main

// Fann - Fast artificial neural network library
// Homepage: https://sourceforge.net/projects/fann/

import (
	"fmt"
	
	"os/exec"
)

func installFann() {
	// Método 1: Descargar y extraer .tar.gz
	fann_tar_url := "https://downloads.sourceforge.net/project/fann/fann/2.2.0/FANN-2.2.0-Source.tar.gz"
	fann_cmd_tar := exec.Command("curl", "-L", fann_tar_url, "-o", "package.tar.gz")
	err := fann_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fann_zip_url := "https://downloads.sourceforge.net/project/fann/fann/2.2.0/FANN-2.2.0-Source.zip"
	fann_cmd_zip := exec.Command("curl", "-L", fann_zip_url, "-o", "package.zip")
	err = fann_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fann_bin_url := "https://downloads.sourceforge.net/project/fann/fann/2.2.0/FANN-2.2.0-Source.bin"
	fann_cmd_bin := exec.Command("curl", "-L", fann_bin_url, "-o", "binary.bin")
	err = fann_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fann_src_url := "https://downloads.sourceforge.net/project/fann/fann/2.2.0/FANN-2.2.0-Source.src.tar.gz"
	fann_cmd_src := exec.Command("curl", "-L", fann_src_url, "-o", "source.tar.gz")
	err = fann_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fann_cmd_direct := exec.Command("./binary")
	err = fann_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
