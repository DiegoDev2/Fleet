package main

// Goolabs - Command-line tool for morphologically analyzing Japanese language
// Homepage: https://pypi.python.org/pypi/goolabs

import (
	"fmt"
	
	"os/exec"
)

func installGoolabs() {
	// Método 1: Descargar y extraer .tar.gz
	goolabs_tar_url := "https://files.pythonhosted.org/packages/ce/86/2d3b5bd85311ee3a7ae7a661b3619095431503cd0cae03048c646b700cad/goolabs-0.4.0.tar.gz"
	goolabs_cmd_tar := exec.Command("curl", "-L", goolabs_tar_url, "-o", "package.tar.gz")
	err := goolabs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	goolabs_zip_url := "https://files.pythonhosted.org/packages/ce/86/2d3b5bd85311ee3a7ae7a661b3619095431503cd0cae03048c646b700cad/goolabs-0.4.0.zip"
	goolabs_cmd_zip := exec.Command("curl", "-L", goolabs_zip_url, "-o", "package.zip")
	err = goolabs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	goolabs_bin_url := "https://files.pythonhosted.org/packages/ce/86/2d3b5bd85311ee3a7ae7a661b3619095431503cd0cae03048c646b700cad/goolabs-0.4.0.bin"
	goolabs_cmd_bin := exec.Command("curl", "-L", goolabs_bin_url, "-o", "binary.bin")
	err = goolabs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	goolabs_src_url := "https://files.pythonhosted.org/packages/ce/86/2d3b5bd85311ee3a7ae7a661b3619095431503cd0cae03048c646b700cad/goolabs-0.4.0.src.tar.gz"
	goolabs_cmd_src := exec.Command("curl", "-L", goolabs_src_url, "-o", "source.tar.gz")
	err = goolabs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	goolabs_cmd_direct := exec.Command("./binary")
	err = goolabs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
