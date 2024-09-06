package main

// Qrupdate - Fast updates of QR and Cholesky decompositions
// Homepage: https://sourceforge.net/projects/qrupdate/

import (
	"fmt"
	
	"os/exec"
)

func installQrupdate() {
	// Método 1: Descargar y extraer .tar.gz
	qrupdate_tar_url := "https://downloads.sourceforge.net/project/qrupdate/qrupdate/1.2/qrupdate-1.1.2.tar.gz"
	qrupdate_cmd_tar := exec.Command("curl", "-L", qrupdate_tar_url, "-o", "package.tar.gz")
	err := qrupdate_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qrupdate_zip_url := "https://downloads.sourceforge.net/project/qrupdate/qrupdate/1.2/qrupdate-1.1.2.zip"
	qrupdate_cmd_zip := exec.Command("curl", "-L", qrupdate_zip_url, "-o", "package.zip")
	err = qrupdate_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qrupdate_bin_url := "https://downloads.sourceforge.net/project/qrupdate/qrupdate/1.2/qrupdate-1.1.2.bin"
	qrupdate_cmd_bin := exec.Command("curl", "-L", qrupdate_bin_url, "-o", "binary.bin")
	err = qrupdate_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qrupdate_src_url := "https://downloads.sourceforge.net/project/qrupdate/qrupdate/1.2/qrupdate-1.1.2.src.tar.gz"
	qrupdate_cmd_src := exec.Command("curl", "-L", qrupdate_src_url, "-o", "source.tar.gz")
	err = qrupdate_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qrupdate_cmd_direct := exec.Command("./binary")
	err = qrupdate_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
}
