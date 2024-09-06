package main

// Briss - Crop PDF files
// Homepage: https://briss.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installBriss() {
	// Método 1: Descargar y extraer .tar.gz
	briss_tar_url := "https://downloads.sourceforge.net/project/briss/release%200.9/briss-0.9.tar.gz"
	briss_cmd_tar := exec.Command("curl", "-L", briss_tar_url, "-o", "package.tar.gz")
	err := briss_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	briss_zip_url := "https://downloads.sourceforge.net/project/briss/release%200.9/briss-0.9.zip"
	briss_cmd_zip := exec.Command("curl", "-L", briss_zip_url, "-o", "package.zip")
	err = briss_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	briss_bin_url := "https://downloads.sourceforge.net/project/briss/release%200.9/briss-0.9.bin"
	briss_cmd_bin := exec.Command("curl", "-L", briss_bin_url, "-o", "binary.bin")
	err = briss_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	briss_src_url := "https://downloads.sourceforge.net/project/briss/release%200.9/briss-0.9.src.tar.gz"
	briss_cmd_src := exec.Command("curl", "-L", briss_src_url, "-o", "source.tar.gz")
	err = briss_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	briss_cmd_direct := exec.Command("./binary")
	err = briss_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
