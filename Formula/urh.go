package main

// Urh - Universal Radio Hacker
// Homepage: https://github.com/jopohl/urh

import (
	"fmt"
	
	"os/exec"
)

func installUrh() {
	// Método 1: Descargar y extraer .tar.gz
	urh_tar_url := "https://files.pythonhosted.org/packages/d8/dc/a6dcf5686e980530b23bc16936cd9c879c50da133f319f729da6d20bd95b/urh-2.9.6.tar.gz"
	urh_cmd_tar := exec.Command("curl", "-L", urh_tar_url, "-o", "package.tar.gz")
	err := urh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	urh_zip_url := "https://files.pythonhosted.org/packages/d8/dc/a6dcf5686e980530b23bc16936cd9c879c50da133f319f729da6d20bd95b/urh-2.9.6.zip"
	urh_cmd_zip := exec.Command("curl", "-L", urh_zip_url, "-o", "package.zip")
	err = urh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	urh_bin_url := "https://files.pythonhosted.org/packages/d8/dc/a6dcf5686e980530b23bc16936cd9c879c50da133f319f729da6d20bd95b/urh-2.9.6.bin"
	urh_cmd_bin := exec.Command("curl", "-L", urh_bin_url, "-o", "binary.bin")
	err = urh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	urh_src_url := "https://files.pythonhosted.org/packages/d8/dc/a6dcf5686e980530b23bc16936cd9c879c50da133f319f729da6d20bd95b/urh-2.9.6.src.tar.gz"
	urh_cmd_src := exec.Command("curl", "-L", urh_src_url, "-o", "source.tar.gz")
	err = urh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	urh_cmd_direct := exec.Command("./binary")
	err = urh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: hackrf")
	exec.Command("latte", "install", "hackrf").Run()
	fmt.Println("Instalando dependencia: numpy")
	exec.Command("latte", "install", "numpy").Run()
	fmt.Println("Instalando dependencia: pyqt@5")
	exec.Command("latte", "install", "pyqt@5").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
