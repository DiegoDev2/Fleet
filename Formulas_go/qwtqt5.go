package main

// QwtQt5 - Qt Widgets for Technical Applications
// Homepage: https://qwt.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installQwtQt5() {
	// Método 1: Descargar y extraer .tar.gz
	qwtqt5_tar_url := "https://downloads.sourceforge.net/project/qwt/qwt/6.3.0/qwt-6.3.0.tar.bz2"
	qwtqt5_cmd_tar := exec.Command("curl", "-L", qwtqt5_tar_url, "-o", "package.tar.gz")
	err := qwtqt5_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qwtqt5_zip_url := "https://downloads.sourceforge.net/project/qwt/qwt/6.3.0/qwt-6.3.0.tar.bz2"
	qwtqt5_cmd_zip := exec.Command("curl", "-L", qwtqt5_zip_url, "-o", "package.zip")
	err = qwtqt5_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qwtqt5_bin_url := "https://downloads.sourceforge.net/project/qwt/qwt/6.3.0/qwt-6.3.0.tar.bz2"
	qwtqt5_cmd_bin := exec.Command("curl", "-L", qwtqt5_bin_url, "-o", "binary.bin")
	err = qwtqt5_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qwtqt5_src_url := "https://downloads.sourceforge.net/project/qwt/qwt/6.3.0/qwt-6.3.0.tar.bz2"
	qwtqt5_cmd_src := exec.Command("curl", "-L", qwtqt5_src_url, "-o", "source.tar.gz")
	err = qwtqt5_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qwtqt5_cmd_direct := exec.Command("./binary")
	err = qwtqt5_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: qt@5")
exec.Command("latte", "install", "qt@5")
}
