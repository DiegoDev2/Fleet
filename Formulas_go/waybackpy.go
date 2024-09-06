package main

// Waybackpy - Wayback Machine API interface & command-line tool
// Homepage: https://pypi.org/project/waybackpy/

import (
	"fmt"
	
	"os/exec"
)

func installWaybackpy() {
	// Método 1: Descargar y extraer .tar.gz
	waybackpy_tar_url := "https://files.pythonhosted.org/packages/34/ab/90085feb81e7fad7d00c736f98e74ec315159ebef2180a77c85a06b2f0aa/waybackpy-3.0.6.tar.gz"
	waybackpy_cmd_tar := exec.Command("curl", "-L", waybackpy_tar_url, "-o", "package.tar.gz")
	err := waybackpy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	waybackpy_zip_url := "https://files.pythonhosted.org/packages/34/ab/90085feb81e7fad7d00c736f98e74ec315159ebef2180a77c85a06b2f0aa/waybackpy-3.0.6.zip"
	waybackpy_cmd_zip := exec.Command("curl", "-L", waybackpy_zip_url, "-o", "package.zip")
	err = waybackpy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	waybackpy_bin_url := "https://files.pythonhosted.org/packages/34/ab/90085feb81e7fad7d00c736f98e74ec315159ebef2180a77c85a06b2f0aa/waybackpy-3.0.6.bin"
	waybackpy_cmd_bin := exec.Command("curl", "-L", waybackpy_bin_url, "-o", "binary.bin")
	err = waybackpy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	waybackpy_src_url := "https://files.pythonhosted.org/packages/34/ab/90085feb81e7fad7d00c736f98e74ec315159ebef2180a77c85a06b2f0aa/waybackpy-3.0.6.src.tar.gz"
	waybackpy_cmd_src := exec.Command("curl", "-L", waybackpy_src_url, "-o", "source.tar.gz")
	err = waybackpy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	waybackpy_cmd_direct := exec.Command("./binary")
	err = waybackpy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
