package main

// Fred - Fully featured FRED Command-line Interface & Python API wrapper
// Homepage: https://fred.stlouisfed.org/docs/api/fred/

import (
	"fmt"
	
	"os/exec"
)

func installFred() {
	// Método 1: Descargar y extraer .tar.gz
	fred_tar_url := "https://files.pythonhosted.org/packages/ff/22/44051587a95461a3fb0cd57e5ba215f3c4d3086544294e5ac79ab0028c20/fred_py_api-1.2.0.tar.gz"
	fred_cmd_tar := exec.Command("curl", "-L", fred_tar_url, "-o", "package.tar.gz")
	err := fred_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fred_zip_url := "https://files.pythonhosted.org/packages/ff/22/44051587a95461a3fb0cd57e5ba215f3c4d3086544294e5ac79ab0028c20/fred_py_api-1.2.0.zip"
	fred_cmd_zip := exec.Command("curl", "-L", fred_zip_url, "-o", "package.zip")
	err = fred_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fred_bin_url := "https://files.pythonhosted.org/packages/ff/22/44051587a95461a3fb0cd57e5ba215f3c4d3086544294e5ac79ab0028c20/fred_py_api-1.2.0.bin"
	fred_cmd_bin := exec.Command("curl", "-L", fred_bin_url, "-o", "binary.bin")
	err = fred_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fred_src_url := "https://files.pythonhosted.org/packages/ff/22/44051587a95461a3fb0cd57e5ba215f3c4d3086544294e5ac79ab0028c20/fred_py_api-1.2.0.src.tar.gz"
	fred_cmd_src := exec.Command("curl", "-L", fred_src_url, "-o", "source.tar.gz")
	err = fred_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fred_cmd_direct := exec.Command("./binary")
	err = fred_cmd_direct.Run()
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
