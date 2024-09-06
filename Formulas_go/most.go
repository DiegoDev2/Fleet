package main

// Most - Powerful paging program
// Homepage: https://www.jedsoft.org/most/

import (
	"fmt"
	
	"os/exec"
)

func installMost() {
	// Método 1: Descargar y extraer .tar.gz
	most_tar_url := "https://www.jedsoft.org/releases/most/most-5.2.0.tar.gz"
	most_cmd_tar := exec.Command("curl", "-L", most_tar_url, "-o", "package.tar.gz")
	err := most_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	most_zip_url := "https://www.jedsoft.org/releases/most/most-5.2.0.zip"
	most_cmd_zip := exec.Command("curl", "-L", most_zip_url, "-o", "package.zip")
	err = most_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	most_bin_url := "https://www.jedsoft.org/releases/most/most-5.2.0.bin"
	most_cmd_bin := exec.Command("curl", "-L", most_bin_url, "-o", "binary.bin")
	err = most_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	most_src_url := "https://www.jedsoft.org/releases/most/most-5.2.0.src.tar.gz"
	most_cmd_src := exec.Command("curl", "-L", most_src_url, "-o", "source.tar.gz")
	err = most_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	most_cmd_direct := exec.Command("./binary")
	err = most_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: s-lang")
exec.Command("latte", "install", "s-lang")
}
