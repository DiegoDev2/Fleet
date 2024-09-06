package main

// Qwt - Qt Widgets for Technical Applications
// Homepage: https://qwt.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installQwt() {
	// Método 1: Descargar y extraer .tar.gz
	qwt_tar_url := "https://downloads.sourceforge.net/project/qwt/qwt/6.3.0/qwt-6.3.0.tar.bz2"
	qwt_cmd_tar := exec.Command("curl", "-L", qwt_tar_url, "-o", "package.tar.gz")
	err := qwt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qwt_zip_url := "https://downloads.sourceforge.net/project/qwt/qwt/6.3.0/qwt-6.3.0.tar.bz2"
	qwt_cmd_zip := exec.Command("curl", "-L", qwt_zip_url, "-o", "package.zip")
	err = qwt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qwt_bin_url := "https://downloads.sourceforge.net/project/qwt/qwt/6.3.0/qwt-6.3.0.tar.bz2"
	qwt_cmd_bin := exec.Command("curl", "-L", qwt_bin_url, "-o", "binary.bin")
	err = qwt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qwt_src_url := "https://downloads.sourceforge.net/project/qwt/qwt/6.3.0/qwt-6.3.0.tar.bz2"
	qwt_cmd_src := exec.Command("curl", "-L", qwt_src_url, "-o", "source.tar.gz")
	err = qwt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qwt_cmd_direct := exec.Command("./binary")
	err = qwt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: qt")
exec.Command("latte", "install", "qt")
}
